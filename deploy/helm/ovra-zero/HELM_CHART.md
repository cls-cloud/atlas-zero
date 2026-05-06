# Ovra-Zero Helm Chart 说明文档

本文档说明 `deploy/helm/ovra-zero` 这套 Helm Chart 的结构、配置项、模板逻辑、安装升级方式和常见维护操作。

## 一、Chart 目录结构

```text
deploy/helm/ovra-zero
├── Chart.yaml
├── values.yaml
├── README.md
├── INSTALL_REDIS_ETCD.md
├── HELM_CHART.md
├── files
│   └── config
│       ├── auth.yaml.gotmpl
│       ├── demo.yaml.gotmpl
│       └── system.yaml.gotmpl
└── templates
    ├── _helpers.tpl
    ├── apps.yaml
    ├── configmap.yaml
    ├── etcd.yaml
    ├── ingress.yaml
    ├── mysql.yaml
    └── redis.yaml
```

### 1. Chart.yaml

Chart 元信息：

- `name: ovra-zero`
- `type: application`
- `version: 0.1.0`
- `appVersion: local`

### 2. values.yaml

集中管理所有可调配置：

- 应用镜像
- MySQL 连接方式
- Redis Cluster 配置
- etcd 集群配置
- auth/system/demo 服务端口
- Ingress 路由

### 3. files/config/*.gotmpl

这三份文件是应用配置模板：

- `auth.yaml.gotmpl`
- `system.yaml.gotmpl`
- `demo.yaml.gotmpl`

它们不是普通静态配置，而是由 Helm `tpl` 函数渲染。这样 Redis 和 etcd 的节点地址可以根据 `values.yaml` 中的副本数量自动生成。

### 4. templates

Kubernetes 资源模板：

| 文件 | 作用 |
| --- | --- |
| `_helpers.tpl` | 公共 label、etcd 地址、Redis Cluster 地址生成函数 |
| `apps.yaml` | auth/system/demo Deployment 与 Service |
| `configmap.yaml` | 渲染应用配置 ConfigMap |
| `etcd.yaml` | etcd Service、Headless Service、StatefulSet |
| `redis.yaml` | Redis Service、Headless Service、StatefulSet、Cluster 初始化 Job |
| `mysql.yaml` | MySQL Service、外部 EndpointSlice 或内置 Deployment |
| `ingress.yaml` | Traefik Ingress |

## 二、核心资源说明

### 1. 应用服务

`apps.yaml` 会根据 `values.services` 生成 3 组应用资源：

- `auth`
- `system`
- `demo`

每个服务会生成：

- `Deployment`
- `Service`

应用容器使用同一个镜像：

```yaml
image:
  repository: ovra-zero
  tag: local
  pullPolicy: IfNotPresent
```

容器启动命令由各服务单独配置：

```yaml
services:
  auth:
    command: /app/app-auth
  system:
    command: /app/app-system
  demo:
    command: /app/app-demo
```

应用配置挂载路径：

```yaml
config:
  mountPath: /app/etc/k3d
```

应用代码里 `ENV=k3d` 时，会读取：

```text
etc/k3d/auth.yaml
etc/k3d/system.yaml
etc/k3d/demo.yaml
```

### 2. ConfigMap

`configmap.yaml` 会把 `files/config/*.gotmpl` 渲染为：

```text
ovra-zero-config
```

最终 ConfigMap 中包含：

- `auth.yaml`
- `system.yaml`
- `demo.yaml`

模板里会动态渲染：

- etcd endpoint 列表
- Redis Cluster host 字符串
- MySQL host

`apps.yaml` 中有配置校验和 annotation：

```yaml
checksum/config: ...
```

当配置模板内容变化时，Deployment 会自动滚动重启。

### 3. etcd

`etcd.yaml` 创建：

- `Service/etcd`
- `Service/etcd-headless`
- `StatefulSet/etcd`

默认配置：

```yaml
etcd:
  name: etcd
  replicas: 3
  port: 2379
  peerPort: 2380
  headlessServiceName: etcd-headless
```

Headless Service 用于生成稳定 DNS：

```text
etcd-0.etcd-headless.ovra-zero.svc.cluster.local
etcd-1.etcd-headless.ovra-zero.svc.cluster.local
etcd-2.etcd-headless.ovra-zero.svc.cluster.local
```

`_helpers.tpl` 中的 `ovra-zero.etcdInitialCluster` 会生成 etcd 初始集群参数：

```text
etcd-0=http://etcd-0.etcd-headless...:2380,
etcd-1=http://etcd-1.etcd-headless...:2380,
etcd-2=http://etcd-2.etcd-headless...:2380
```

应用配置中的 etcd client hosts 由 `ovra-zero.etcdClientHostsYaml` 生成。

### 4. Redis Cluster

`redis.yaml` 创建：

- `Service/redis`
- `Service/redis-headless`
- `StatefulSet/redis`
- `Job/redis-cluster-init`

默认配置：

```yaml
redis:
  password: Pl@1221view
  port: 6379
  busPort: 16379
  cluster:
    enabled: true
    replicas: 3
    replicasPerMaster: 0
```

Redis Cluster 使用两个端口：

- `6379`：客户端访问端口
- `16379`：cluster bus 端口，一般为 `6379 + 10000`

Headless Service 用于稳定节点地址：

```text
redis-0.redis-headless.ovra-zero.svc.cluster.local
redis-1.redis-headless.ovra-zero.svc.cluster.local
redis-2.redis-headless.ovra-zero.svc.cluster.local
```

`redis-cluster-init` 是 Helm hook Job：

```yaml
"helm.sh/hook": post-install,post-upgrade
```

它会在安装或升级后执行：

```sh
redis-cli --cluster create ... --cluster-replicas 0 --cluster-yes
```

如果 Redis Cluster 已经是 `cluster_state:ok`，Job 会直接退出，不会重复初始化。

### 5. MySQL

`mysql.yaml` 支持两种模式。

#### 外部 MySQL

默认启用：

```yaml
mysql:
  external:
    enabled: true
    ip: 172.18.0.2
  internal:
    enabled: false
```

Chart 会创建：

- `Service/mysql`
- `EndpointSlice/mysql-external`

这种模式适合本地 k3d：MySQL 容器跑在 `k3d-ovra` Docker 网络里，Kubernetes 通过 EndpointSlice 访问。

#### 内置 MySQL

如果希望由 Kubernetes 内部创建 MySQL，可以改为：

```yaml
mysql:
  external:
    enabled: false
  internal:
    enabled: true
```

可选初始化 SQL ConfigMap：

```yaml
mysql:
  internal:
    initSqlConfigMap: ovra-zero-sql
```

注意：当前内置 MySQL 使用 `emptyDir`，只适合开发测试。生产环境应改成 PVC。

### 6. Ingress

`ingress.yaml` 默认使用 Traefik：

```yaml
ingress:
  enabled: true
  className: traefik
  host: ovra-zero.localhost
```

默认路由：

| Path | Service |
| --- | --- |
| `/auth` | `auth:8091` |
| `/system` | `system:8092` |
| `/monitor` | `system:8092` |
| `/resource` | `system:8092` |
| `/demo` | `demo:8099` |

## 三、安装前准备

### 1. 构建应用镜像

```sh
mkdir -p .deploy/k3d/bin

GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build \
  -ldflags='-s -w' \
  -tags no_k8s \
  -o .deploy/k3d/bin/app-auth \
  app/auth/auth.go

GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build \
  -ldflags='-s -w' \
  -tags no_k8s \
  -o .deploy/k3d/bin/app-system \
  app/system/system.go

GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build \
  -ldflags='-s -w' \
  -tags no_k8s \
  -o .deploy/k3d/bin/app-demo \
  app/demo/demo.go

docker build --platform linux/arm64 \
  -t ovra-zero:local \
  -f deploy/k3d/Dockerfile .
```

如果你的 k3d 节点是 x86_64，把 `GOARCH=arm64` 和 `--platform linux/arm64` 改成：

```text
GOARCH=amd64
--platform linux/amd64
```

### 2. 导入镜像到 k3d

```sh
k3d image import ovra-zero:local -c ovra
```

### 3. 准备本地 MySQL

默认 values 使用外部 MySQL EndpointSlice。启动本地 MySQL 容器：

```sh
docker run -d --name ovra-zero-mysql --network k3d-ovra \
  -e MYSQL_ROOT_PASSWORD='Pl@1221view' \
  -e MYSQL_DATABASE=ovra_zero \
  -v "$PWD/bin/sql/ovra_zero.sql:/docker-entrypoint-initdb.d/ovra_zero.sql:ro" \
  mysql:8.4
```

确认容器 IP：

```sh
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ovra-zero-mysql
```

如果 IP 不是 `172.18.0.2`，安装时覆盖：

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --set mysql.external.ip=<MYSQL_CONTAINER_IP>
```

## 四、安装与升级

### 1. 安装

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --wait \
  --timeout 5m
```

### 2. 查看 release

```sh
helm status ovra-zero -n ovra-zero
helm history ovra-zero -n ovra-zero
```

### 3. 升级

修改 `values.yaml` 或模板后执行：

```sh
helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --wait \
  --timeout 5m
```

### 4. 回滚

查看历史：

```sh
helm history ovra-zero -n ovra-zero
```

回滚到指定 revision：

```sh
helm rollback ovra-zero <REVISION> -n ovra-zero
```

### 5. 卸载

```sh
helm uninstall ovra-zero -n ovra-zero
```

如果保留了外部 MySQL 容器，Helm 卸载不会删除它。

## 五、渲染与校验

### 1. Helm lint

```sh
helm lint deploy/helm/ovra-zero
```

### 2. 渲染 YAML

```sh
helm template ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  > /tmp/ovra-zero-rendered.yaml
```

### 3. Kubernetes dry-run

```sh
kubectl apply --dry-run=client -f /tmp/ovra-zero-rendered.yaml
```

### 4. 查看渲染后的应用配置

```sh
helm template ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --show-only templates/configmap.yaml
```

## 六、常见配置调整

### 1. 修改应用镜像

```yaml
image:
  repository: your-registry.example.com/ovra-zero
  tag: v1.0.0
  pullPolicy: IfNotPresent
```

升级：

```sh
helm upgrade ovra-zero deploy/helm/ovra-zero -n ovra-zero --wait --timeout 5m
```

### 2. 修改 Ingress 域名

```yaml
ingress:
  host: api.example.com
```

或命令行覆盖：

```sh
helm upgrade ovra-zero deploy/helm/ovra-zero \
  -n ovra-zero \
  --set ingress.host=api.example.com
```

### 3. 关闭 demo 服务

```yaml
services:
  demo:
    enabled: false
```

同时建议删除 Ingress 里的 `/demo` 路由。

### 4. 使用内置 MySQL

```yaml
mysql:
  external:
    enabled: false
  internal:
    enabled: true
    initSqlConfigMap: ovra-zero-sql
```

先创建 SQL ConfigMap：

```sh
kubectl -n ovra-zero create configmap ovra-zero-sql \
  --from-file=ovra_zero.sql=bin/sql/ovra_zero.sql
```

### 5. 改 Redis 节点数量

当前默认是 3 节点：

```yaml
redis:
  cluster:
    replicas: 3
    replicasPerMaster: 0
```

改成 6 节点，3 master + 3 replica：

```yaml
redis:
  cluster:
    replicas: 6
    replicasPerMaster: 1
```

注意：Redis Cluster 拓扑变更涉及 slot 和 node 元数据。开发环境可以重建 StatefulSet；生产环境应做 slot 迁移和节点下线。

开发环境重建：

```sh
kubectl -n ovra-zero delete statefulset redis --ignore-not-found
kubectl -n ovra-zero delete pod -l app.kubernetes.io/component=redis --ignore-not-found
kubectl -n ovra-zero delete job redis-cluster-init --ignore-not-found
helm upgrade ovra-zero deploy/helm/ovra-zero -n ovra-zero --wait --timeout 5m
```

### 6. 改 etcd 节点数量

当前默认是 3 节点：

```yaml
etcd:
  replicas: 3
```

etcd 成员数量不建议随意调整。生产环境扩缩容应使用 `etcdctl member add/remove`，并同步调整 Helm values。

开发环境如需重建：

```sh
kubectl -n ovra-zero delete statefulset etcd --ignore-not-found
kubectl -n ovra-zero delete pod -l app.kubernetes.io/component=etcd --ignore-not-found
helm upgrade ovra-zero deploy/helm/ovra-zero -n ovra-zero --wait --timeout 5m
```

## 七、验证命令

### 1. 查看所有资源

```sh
kubectl -n ovra-zero get pods,svc,statefulset,ingress,endpointslice
```

### 2. 验证 etcd

```sh
kubectl -n ovra-zero exec etcd-0 -- etcdctl \
  --endpoints=http://etcd-0.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-1.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-2.etcd-headless.ovra-zero.svc.cluster.local:2379 \
  member list
```

### 3. 验证 Redis

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster nodes
```

### 4. 验证应用入口

```sh
curl -sS http://ovra-zero.localhost:18080/auth/code
```

期望返回：

```json
{"code":200,"msg":"操作成功",...}
```

## 八、排障

### 1. Helm upgrade 卡住

查看状态：

```sh
helm status ovra-zero -n ovra-zero
kubectl -n ovra-zero get pods
```

查看事件：

```sh
kubectl -n ovra-zero describe pod <POD_NAME>
```

### 2. 应用连接不上 etcd

检查 etcd 是否 Ready：

```sh
kubectl -n ovra-zero get statefulset etcd
kubectl -n ovra-zero logs etcd-0 --tail=100
```

检查应用配置：

```sh
kubectl -n ovra-zero get configmap ovra-zero-config -o yaml
```

### 3. 应用连接不上 Redis

检查 Redis Cluster：

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
```

如果 `cluster_state` 不是 `ok`，查看初始化 Job：

```sh
kubectl -n ovra-zero get job redis-cluster-init
kubectl -n ovra-zero logs job/redis-cluster-init
```

### 4. MySQL 连接失败

如果使用外部 MySQL EndpointSlice：

```sh
kubectl -n ovra-zero get endpointslice mysql-external
docker ps --filter name=ovra-zero-mysql
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ovra-zero-mysql
```

确认 `mysql.external.ip` 与容器 IP 一致。

### 5. Pod 使用旧配置

查看 Deployment annotation：

```sh
kubectl -n ovra-zero get deployment auth -o yaml | grep checksum/config
```

手动重启：

```sh
kubectl -n ovra-zero rollout restart deployment/auth deployment/system deployment/demo
```

## 九、生产环境注意事项

当前 chart 是本地 k3d 和开发环境友好的版本。生产使用前建议调整：

- etcd 使用 PVC。
- Redis 使用 PVC。
- MySQL 使用外部高可用数据库或 StatefulSet + PVC。
- etcd 开启 TLS 和认证。
- Redis 密码改为 Secret，不要明文放在 values。
- 应用镜像推送到正式镜像仓库。
- Ingress 使用正式域名和 TLS。
- Redis 生产建议 6 节点：3 master + 3 replica。
- 设置 resources requests/limits。
- 增加 livenessProbe 和更严格的 readinessProbe。
