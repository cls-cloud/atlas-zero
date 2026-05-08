# Redis 与 etcd 安装部署文档

本文档说明 Ovra-Zero 使用的 etcd 集群与 Redis Cluster 的安装方式，覆盖两种场景：

- Kubernetes/Helm 部署
- Linux 裸机 systemd 部署

当前默认拓扑：

- etcd：3 节点集群
- Redis：3 节点 Redis Cluster，3 个 master，0 个 replica

> 说明：3 节点 Redis Cluster 可以完成槽位分片，但没有 Redis 层面的副本容灾。如果生产环境需要 Redis 自动故障转移，建议使用 6 节点：3 master + 3 replica。

## 一、端口与拓扑

### 1. Kubernetes 拓扑

| 组件 | 资源类型 | 副本数 | 端口 | Service |
| --- | --- | ---: | --- | --- |
| etcd | StatefulSet | 3 | `2379` client，`2380` peer | `etcd`，`etcd-headless` |
| Redis | StatefulSet | 3 | `6379` client，`16379` cluster bus | `redis`，`redis-headless` |

Helm 会自动渲染应用配置：

- etcd 地址渲染为 `etcd-0`、`etcd-1`、`etcd-2` 三个节点。
- Redis 配置为 `Type: cluster`。
- go-zero 的 `RedisConf` 要求 Redis Cluster 地址写成一个逗号分隔字符串，所以 Helm 会把 3 个 Redis 节点渲染到同一个 `Host` 字段。

### 2. Linux 裸机示例拓扑

下面示例使用 3 台机器，请替换成你的真实服务器 IP：

| 节点 | IP | etcd 名称 | Redis 角色 |
| --- | --- | --- | --- |
| node1 | `10.0.0.11` | `etcd-1` | master |
| node2 | `10.0.0.12` | `etcd-2` | master |
| node3 | `10.0.0.13` | `etcd-3` | master |

服务器之间需要开放端口：

- etcd：`2379/tcp`、`2380/tcp`
- Redis：`6379/tcp`、`16379/tcp`

如果使用防火墙，示例：

```sh
firewall-cmd --permanent --add-port=2379/tcp
firewall-cmd --permanent --add-port=2380/tcp
firewall-cmd --permanent --add-port=6379/tcp
firewall-cmd --permanent --add-port=16379/tcp
firewall-cmd --reload
```

## 二、Helm 部署

### 1. 配置项

配置文件位置：

```text
deploy/helm/ovra-zero/values.yaml
```

关键配置：

```yaml
redis:
  password: Pl@1221view
  port: 6379
  busPort: 16379
  cluster:
    enabled: true
    replicas: 3
    replicasPerMaster: 0

etcd:
  port: 2379
  peerPort: 2380
  replicas: 3
  initialClusterToken: ovra-zero-etcd
```

### 2. 安装或升级

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --wait \
  --timeout 5m
```

### 3. 验证 etcd

```sh
kubectl -n ovra-zero get statefulset etcd
```

期望看到：

```text
etcd   3/3
```

查看成员列表：

```sh
kubectl -n ovra-zero exec etcd-0 -- etcdctl \
  --endpoints=http://etcd-0.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-1.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-2.etcd-headless.ovra-zero.svc.cluster.local:2379 \
  member list
```

期望结果：

- 有 3 个 member。
- 状态均为 `started`。

### 4. 验证 Redis Cluster

```sh
kubectl -n ovra-zero get statefulset redis
```

期望看到：

```text
redis   3/3
```

查看 Redis Cluster 状态：

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
```

期望关键字段：

```text
cluster_state:ok
cluster_known_nodes:3
cluster_slots_assigned:16384
cluster_slots_ok:16384
```

查看节点与槽位：

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster nodes
```

3 节点模式下，期望看到 3 个 master，并且 16384 个 slot 被分配到 3 个节点。

### 5. 修改 Redis 节点数后的重建方式

Redis Cluster 会在节点本地保存 cluster 元数据。开发环境使用 `emptyDir` 时，修改节点数后建议直接重建 Redis StatefulSet：

```sh
kubectl -n ovra-zero delete statefulset redis --ignore-not-found
kubectl -n ovra-zero delete pod -l app.kubernetes.io/component=redis --ignore-not-found
kubectl -n ovra-zero delete job redis-cluster-init --ignore-not-found

helm upgrade ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --wait \
  --timeout 5m
```

生产环境不要直接删除有持久化数据的 Redis 节点，应先迁移 slot、下线节点，再调整拓扑。

## 三、Linux 裸机部署 etcd 集群

### 1. 安装 etcd

在三台机器都执行。以下示例以 etcd `v3.5.13` 为例。

ARM64 机器：

```sh
ETCD_VERSION=v3.5.13
ARCH=arm64
```

x86_64 机器：

```sh
ETCD_VERSION=v3.5.13
ARCH=amd64
```

下载安装：

```sh
curl -L "https://github.com/etcd-io/etcd/releases/download/${ETCD_VERSION}/etcd-${ETCD_VERSION}-linux-${ARCH}.tar.gz" \
  -o /tmp/etcd.tar.gz

tar -xzf /tmp/etcd.tar.gz -C /tmp
install -m 0755 /tmp/etcd-${ETCD_VERSION}-linux-${ARCH}/etcd /usr/local/bin/etcd
install -m 0755 /tmp/etcd-${ETCD_VERSION}-linux-${ARCH}/etcdctl /usr/local/bin/etcdctl
```

创建用户和目录：

```sh
useradd --system --home /var/lib/etcd --shell /usr/sbin/nologin etcd || true
mkdir -p /var/lib/etcd
chown -R etcd:etcd /var/lib/etcd
chmod 700 /var/lib/etcd
```

### 2. 配置 node1

创建：

```text
/etc/systemd/system/etcd.service
```

node1 示例：

```ini
[Unit]
Description=etcd key-value store
Documentation=https://etcd.io/docs/
After=network-online.target
Wants=network-online.target

[Service]
User=etcd
Type=notify
ExecStart=/usr/local/bin/etcd \
  --name=etcd-1 \
  --data-dir=/var/lib/etcd \
  --listen-client-urls=http://0.0.0.0:2379 \
  --advertise-client-urls=http://10.0.0.11:2379 \
  --listen-peer-urls=http://0.0.0.0:2380 \
  --initial-advertise-peer-urls=http://10.0.0.11:2380 \
  --initial-cluster=etcd-1=http://10.0.0.11:2380,etcd-2=http://10.0.0.12:2380,etcd-3=http://10.0.0.13:2380 \
  --initial-cluster-token=ovra-zero-etcd \
  --initial-cluster-state=new
Restart=always
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

### 3. 配置 node2

node2 的 systemd 文件与 node1 基本一致，只改下面几项：

```text
--name=etcd-2
--advertise-client-urls=http://10.0.0.12:2379
--initial-advertise-peer-urls=http://10.0.0.12:2380
```

`--initial-cluster` 保持三台机器一致：

```text
etcd-1=http://10.0.0.11:2380,etcd-2=http://10.0.0.12:2380,etcd-3=http://10.0.0.13:2380
```

### 4. 配置 node3

node3 只改下面几项：

```text
--name=etcd-3
--advertise-client-urls=http://10.0.0.13:2379
--initial-advertise-peer-urls=http://10.0.0.13:2380
```

`--initial-cluster` 同样保持三台机器一致。

### 5. 启动 etcd

三台机器都执行：

```sh
systemctl daemon-reload
systemctl enable --now etcd
systemctl status etcd --no-pager
```

### 6. 验证 etcd

任意一台机器执行：

```sh
export ETCDCTL_API=3

etcdctl \
  --endpoints=http://10.0.0.11:2379,http://10.0.0.12:2379,http://10.0.0.13:2379 \
  endpoint health

etcdctl \
  --endpoints=http://10.0.0.11:2379,http://10.0.0.12:2379,http://10.0.0.13:2379 \
  member list
```

### 7. Ovra-Zero 裸机 etcd 配置

配置文件中的 etcd hosts 写三台机器：

```yaml
Etcd:
  Hosts:
    - 10.0.0.11:2379
    - 10.0.0.12:2379
    - 10.0.0.13:2379
  Key: system.rpc
```

`auth` 服务里的 `SystemRpc.Etcd.Hosts` 也要写这三个地址。

## 四、Linux 裸机部署 Redis Cluster

### 1. 安装 Redis

三台机器都执行。

Debian/Ubuntu：

```sh
apt-get update
apt-get install -y redis-server
```

RHEL/CentOS/Rocky：

```sh
dnf install -y redis
```

如果系统仓库中的 Redis 版本过旧，建议使用源码编译或公司内部软件仓库安装较新的 Redis。

### 2. 准备目录

三台机器都执行：

```sh
mkdir -p /var/lib/redis /var/log/redis
chown -R redis:redis /var/lib/redis /var/log/redis
```

如果系统用户不是 `redis`，请按实际用户调整。

### 3. 配置 node1

编辑：

```text
/etc/redis/redis.conf
```

node1 示例：

```conf
bind 10.0.0.11
port 6379
protected-mode yes

requirepass Pl@1221view
masterauth Pl@1221view

cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 10.0.0.11
cluster-announce-port 6379
cluster-announce-bus-port 16379

appendonly yes
dir /var/lib/redis
logfile /var/log/redis/redis-server.log
supervised systemd
```

### 4. 配置 node2

node2 修改：

```conf
bind 10.0.0.12
cluster-announce-ip 10.0.0.12
```

其他配置保持一致。

### 5. 配置 node3

node3 修改：

```conf
bind 10.0.0.13
cluster-announce-ip 10.0.0.13
```

其他配置保持一致。

### 6. 启动 Redis

不同 Linux 发行版的 Redis service 名称可能不同，二选一执行即可：

```sh
systemctl enable --now redis-server
systemctl status redis-server --no-pager
```

或：

```sh
systemctl enable --now redis
systemctl status redis --no-pager
```

### 7. 创建 Redis Cluster

只需要在任意一台机器执行一次：

```sh
redis-cli -a 'Pl@1221view' --cluster create \
  10.0.0.11:6379 \
  10.0.0.12:6379 \
  10.0.0.13:6379 \
  --cluster-replicas 0 \
  --cluster-yes
```

### 8. 验证 Redis Cluster

```sh
redis-cli -h 10.0.0.11 -p 6379 -a 'Pl@1221view' cluster info
redis-cli -h 10.0.0.11 -p 6379 -a 'Pl@1221view' cluster nodes
```

期望关键字段：

```text
cluster_state:ok
cluster_known_nodes:3
cluster_slots_assigned:16384
```

### 9. Ovra-Zero 裸机 Redis 配置

go-zero 的 Redis Cluster 地址要写成逗号分隔字符串：

```yaml
Data:
  Redis:
    Pass: Pl@1221view
    Host: 10.0.0.11:6379,10.0.0.12:6379,10.0.0.13:6379
    Type: cluster
    Tls: false
```

## 五、常用运维命令

### 1. 查看 etcd 日志

```sh
journalctl -u etcd -f
```

### 2. 查看 Redis 日志

```sh
journalctl -u redis-server -f
```

或：

```sh
journalctl -u redis -f
```

### 3. 检查 Redis slot

```sh
redis-cli -h 10.0.0.11 -p 6379 -a 'Pl@1221view' --cluster check 10.0.0.11:6379
```

### 4. 备份 etcd

```sh
export ETCDCTL_API=3

etcdctl \
  --endpoints=http://10.0.0.11:2379 \
  snapshot save /backup/etcd-$(date +%F-%H%M%S).db
```

### 5. 停止服务

```sh
systemctl stop etcd
systemctl stop redis-server || systemctl stop redis
```

## 六、生产环境建议

- Kubernetes 生产环境应使用 PVC，不要使用 `emptyDir` 保存 etcd 或 Redis 数据。
- etcd 生产环境建议启用 TLS、认证、定期快照备份。
- Redis 生产环境建议至少 6 节点：3 master + 3 replica。
- Redis 的 `requirepass` 和 `masterauth` 必须在所有节点保持一致。
- 所有节点必须保持时间同步，建议启用 chrony 或 ntpd。
- etcd 3 节点集群最多容忍 1 个节点故障。
- Redis 3 master 0 replica 模式不具备 Redis 层面的故障转移能力，适合开发、测试或资源受限环境。
