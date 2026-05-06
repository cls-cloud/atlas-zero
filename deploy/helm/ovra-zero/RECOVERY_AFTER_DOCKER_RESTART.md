# Docker 重启后本地 k3d 服务恢复手册

本文档用于处理本地 Docker Desktop 或 Docker daemon 重启后，`ovra-zero` 在 k3d 中出现服务不可用的问题。

适用场景：

- `curl http://ovra-zero.localhost:18080/auth/code` 返回 `no available server`
- 应用 Pod `CrashLoopBackOff`
- etcd 或 Redis Pod 状态异常
- MySQL EndpointSlice 指向了错误 IP
- k3d 节点变成 `NotReady`

## 一、快速体检

### 1. 检查 Docker 容器

```sh
docker ps --format '{{.Names}} {{.Status}} {{.Networks}}'
```

重点确认这些容器存在且是 `Up`：

```text
k3d-ovra-server-0
k3d-ovra-agent-0
k3d-ovra-agent-1
k3d-ovra-serverlb
ovra-zero-mysql
```

### 2. 检查 k3d 集群

```sh
k3d cluster list
kubectl get nodes -o wide
```

期望：

```text
k3d-ovra-server-0   Ready
k3d-ovra-agent-0    Ready
k3d-ovra-agent-1    Ready
```

如果某个节点是 `NotReady`，按“二、恢复 k3d 节点”处理。

### 3. 检查业务资源

```sh
kubectl -n ovra-zero get pods,svc,statefulset,ingress,endpointslice
helm status ovra-zero -n ovra-zero
```

正常状态应类似：

```text
auth      1/1 Running
demo      1/1 Running
system    1/1 Running
etcd-0    1/1 Running
etcd-1    1/1 Running
etcd-2    1/1 Running
redis-0   1/1 Running
redis-1   1/1 Running
redis-2   1/1 Running

statefulset.apps/etcd    3/3
statefulset.apps/redis   3/3
```

## 二、恢复 k3d 节点

如果 `kubectl get nodes` 中有节点 `NotReady`，先重启对应 Docker 容器。

例如 `k3d-ovra-agent-1` 异常：

```sh
docker restart k3d-ovra-agent-1
```

等待 10-30 秒后检查：

```sh
kubectl get nodes -o wide
```

如果还是 `NotReady`，查看节点原因：

```sh
kubectl describe node k3d-ovra-agent-1
```

常见表现：

```text
Kubelet stopped posting node status
node.kubernetes.io/unreachable
```

一般重启对应 k3d 节点容器即可恢复。

## 三、修复 MySQL EndpointSlice

本地部署默认使用 Docker 网络里的外部 MySQL 容器。Docker 重启后，`ovra-zero-mysql` 的 IP 可能变化。

### 1. 查看 MySQL 新 IP

```sh
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ovra-zero-mysql
```

示例输出：

```text
172.18.0.6
```

### 2. 查看当前 Kubernetes EndpointSlice

```sh
kubectl -n ovra-zero get endpointslice mysql-external -o yaml
```

看这里：

```yaml
endpoints:
  - addresses:
      - 172.18.0.2
```

如果这个 IP 和 MySQL 新 IP 不一致，就需要更新 Helm。

### 3. 用 Helm 更新 MySQL IP

把 `<MYSQL_CONTAINER_IP>` 替换成上一步查到的 IP：

```sh
helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --set mysql.external.ip=<MYSQL_CONTAINER_IP> \
  --wait \
  --timeout 5m
```

示例：

```sh
helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --set mysql.external.ip=172.18.0.6 \
  --wait \
  --timeout 5m
```

### 4. 验证 EndpointSlice

```sh
kubectl -n ovra-zero get endpointslice mysql-external
```

期望：

```text
mysql-external   IPv4   3306   172.18.0.6
```

## 四、恢复 etcd

Docker 重启后，本地 `emptyDir` 模式下的 etcd 可能出现 member 状态不一致。

常见错误：

```text
member xxx has already been bootstrapped
context deadline exceeded
```

或：

```text
statefulset.apps/etcd   1/3
```

### 1. 查看 etcd 状态

```sh
kubectl -n ovra-zero get pods -l app.kubernetes.io/component=etcd
kubectl -n ovra-zero get statefulset etcd
```

### 2. 重建 etcd Pod

本地开发环境使用 `emptyDir`，可以直接重建 3 个 Pod：

```sh
kubectl -n ovra-zero delete pod etcd-0 etcd-1 etcd-2 --ignore-not-found
```

等待恢复：

```sh
kubectl -n ovra-zero rollout status statefulset/etcd --timeout=180s
```

如果 `rollout status` 不适用，也可以直接看：

```sh
kubectl -n ovra-zero get pods -l app.kubernetes.io/component=etcd
```

期望：

```text
etcd-0   1/1 Running
etcd-1   1/1 Running
etcd-2   1/1 Running
```

### 3. 验证 etcd member

```sh
kubectl -n ovra-zero exec etcd-0 -- etcdctl \
  --endpoints=http://etcd-0.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-1.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-2.etcd-headless.ovra-zero.svc.cluster.local:2379 \
  member list
```

期望看到 3 个 member，状态均为：

```text
started
```

## 五、恢复 Redis Cluster

Docker 重启后，本地 `emptyDir` 模式下 Redis Cluster 可能出现 slots 丢失或节点元数据不一致。

常见错误：

```text
CLUSTERDOWN Hash slot not served
cluster_state:fail
cluster_slots_assigned:0
```

或初始化 Job 报：

```text
Node ... is not empty. Either the node already knows other nodes or contains some key in database 0.
```

### 1. 查看 Redis 状态

```sh
kubectl -n ovra-zero get pods -l app.kubernetes.io/component=redis
kubectl -n ovra-zero get statefulset redis
```

查看 cluster info：

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
```

正常应该包含：

```text
cluster_state:ok
cluster_slots_assigned:16384
cluster_known_nodes:3
```

### 2. 重建 Redis Cluster

本地开发环境可以直接重建 Redis StatefulSet 和初始化 Job：

```sh
kubectl -n ovra-zero delete statefulset redis --ignore-not-found
kubectl -n ovra-zero delete pod -l app.kubernetes.io/component=redis --ignore-not-found
kubectl -n ovra-zero delete job redis-cluster-init --ignore-not-found
```

然后重新应用 Helm：

```sh
MYSQL_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ovra-zero-mysql)

helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --set mysql.external.ip=${MYSQL_IP} \
  --wait \
  --timeout 5m
```

### 3. 验证 Redis

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster nodes
```

期望：

```text
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_known_nodes:3
```

## 六、重启应用服务

当 MySQL、etcd、Redis 都恢复后，重启应用服务，让它们重新连接依赖并注册 RPC：

```sh
kubectl -n ovra-zero rollout restart deployment/auth deployment/system deployment/demo
```

等待完成：

```sh
kubectl -n ovra-zero rollout status deployment/auth --timeout=180s
kubectl -n ovra-zero rollout status deployment/system --timeout=180s
kubectl -n ovra-zero rollout status deployment/demo --timeout=180s
```

## 七、最终验证

### 1. 查看资源

```sh
kubectl -n ovra-zero get pods,svc,statefulset,ingress,endpointslice
helm status ovra-zero -n ovra-zero
```

期望：

```text
STATUS: deployed
auth      1/1 Running
demo      1/1 Running
system    1/1 Running
etcd      3/3
redis     3/3
```

### 2. 验证入口

```sh
curl -sS http://ovra-zero.localhost:18080/auth/code
```

期望返回：

```json
{"code":200,"msg":"操作成功",...}
```

## 八、一键恢复命令

如果只是本地开发环境，并且你能接受重建本地 etcd/Redis 数据，可以直接按下面顺序执行：

```sh
# 1. 恢复可能 NotReady 的 k3d agent
docker restart k3d-ovra-agent-1 || true

# 2. 获取 Docker 重启后的 MySQL 新 IP
MYSQL_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ovra-zero-mysql)
echo "MySQL IP: ${MYSQL_IP}"

# 3. 更新 Helm 中的 MySQL EndpointSlice
helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --set mysql.external.ip=${MYSQL_IP} \
  --wait \
  --timeout 5m || true

# 4. 重建本地 etcd
kubectl -n ovra-zero delete pod etcd-0 etcd-1 etcd-2 --ignore-not-found

# 5. 重建本地 Redis Cluster
kubectl -n ovra-zero delete statefulset redis --ignore-not-found
kubectl -n ovra-zero delete pod -l app.kubernetes.io/component=redis --ignore-not-found
kubectl -n ovra-zero delete job redis-cluster-init --ignore-not-found

# 6. 重新 Helm upgrade，触发 Redis 初始化 Job
helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --set mysql.external.ip=${MYSQL_IP} \
  --wait \
  --timeout 5m

# 7. 重启应用
kubectl -n ovra-zero rollout restart deployment/auth deployment/system deployment/demo
kubectl -n ovra-zero rollout status deployment/auth --timeout=180s
kubectl -n ovra-zero rollout status deployment/system --timeout=180s
kubectl -n ovra-zero rollout status deployment/demo --timeout=180s

# 8. 验证
kubectl -n ovra-zero get pods,svc,statefulset,ingress,endpointslice
curl -sS http://ovra-zero.localhost:18080/auth/code
```

## 九、为什么 Docker 重启后会这样

当前本地部署为了轻量和快速，etcd 与 Redis 使用的是 `emptyDir`：

- Pod 重建时数据会丢失。
- Docker 重启时不同 Pod 恢复顺序不一致，可能出现部分旧状态、部分新状态。
- Redis Cluster 的节点元数据和 slot 分配保存在节点本地。
- etcd 的 member bootstrap 信息也保存在节点本地。

另外，本地 MySQL 是 Docker 网络里的外部容器：

- Docker 重启后容器 IP 可能变化。
- Kubernetes 的 `mysql-external` EndpointSlice 不会自动感知 Docker 容器 IP 变化。
- 所以需要用 Helm 重新设置 `mysql.external.ip`。

生产环境建议：

- etcd 使用 PVC。
- Redis 使用 PVC 或外部托管 Redis。
- MySQL 使用固定 Service、固定域名或外部数据库地址。
- 不依赖 Docker bridge 网络里的动态容器 IP。
