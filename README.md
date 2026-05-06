[![GitHub](https://img.shields.io/github/stars/cls-cloud/ovra-zero.svg?style=social&label=Stars)](https://github.com/cls-cloud/ovra-zero)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/cls-cloud/ovra-zero/blob/main/LICENSE)

<h1 align="center" style="margin: 30px 0 30px; font-weight: bold; font-size: 30px">Ovra-Zero</h1>
<h4 align="center">基于 Go-Zero 实现的若依服务端脚手架，支持多租户、RBAC、微服务与本地 Helm 部署</h4>

## 项目简介

Ovra-Zero 是一个基于 [Go-Zero](https://go-zero.dev/) 重写的若依风格后端脚手架，功能模型参考 [RuoYi-Vue-Plus](https://gitee.com/dromara/RuoYi-Vue-Plus.git)，提供企业后台常见的权限、租户、系统管理、审计日志与资源管理能力。

项目适合用于：

- 快速搭建企业级后台管理系统后端。
- 学习 Go-Zero 微服务项目组织方式。
- 基于若依生态前端快速对接后端接口。
- 在本地 k3d/Kubernetes 中验证微服务部署拓扑。

前端项目当前主要适配 [RuoYi-Plus-Vben5](https://gitee.com/dapppp/ruoyi-plus-vben5.git)，也可对接其他兼容 RuoYi-Plus 接口风格的前端。

单体版项目：

- [ovra](https://github.com/ovra-cloud/ovra.git)

## 在线体验

账号密码：

```text
admin/admin123
```

演示与文档：

- 演示地址：[https://vben5.ovra.dev/](https://vben5.ovra.dev/)
- 文档地址：[https://ovra.dev/](https://ovra.dev/)

## 技术栈

- Go 1.24+
- Go-Zero
- gRPC
- GORM / MySQL
- Redis / Redis Cluster
- etcd
- Traefik
- Docker
- Kubernetes / k3d
- Helm

## 服务模块

| 模块 | 路径 | 说明 |
| --- | --- | --- |
| auth | `app/auth` | 登录、登出、验证码、租户列表、认证中间件 |
| system | `app/system` | 用户、角色、菜单、部门、字典、参数、日志、租户、资源等系统功能 |
| demo | `app/demo` | 示例业务模块 |
| toolkit | `toolkit` | 通用工具、认证、租户、GORM 插件、中间件等 |
| desc | `desc` | goctl API/RPC 描述文件 |
| etc | `etc` | 本地开发环境配置 |
| deploy | `deploy` | Docker、Kubernetes、Helm 部署文件 |

## 内置功能

1. 用户管理：系统用户维护、分配角色、岗位、部门等。
2. 部门管理：组织机构树维护，支持数据权限。
3. 岗位管理：维护用户岗位信息。
4. 菜单管理：菜单、按钮、权限标识维护。
5. 角色管理：角色授权、数据范围控制。
6. 字典管理：维护常用固定枚举数据。
7. 参数管理：维护系统动态参数。
8. 操作日志：记录系统正常操作与异常信息。
9. 登录日志：记录登录行为与异常登录。
10. 文件管理：文件上传、下载、配置管理。
11. 租户管理：多租户隔离、租户启停管理。
12. 租户套餐：模块授权、容量限制、套餐维护。
13. 代码生成：配合 GoLand 插件生成后端代码。

## 项目结构

```text
.
├── app                 # 微服务应用
│   ├── auth            # 认证服务
│   ├── demo            # 示例服务
│   └── system          # 系统服务，包含 API 与 RPC
├── bin                 # SQL、Traefik 配置、离线数据文件
├── deploy              # Docker、Kubernetes、Helm 部署文件
├── desc                # goctl API/RPC 描述
├── etc                 # 开发环境配置
├── gen                 # 代码生成配置
├── toolkit             # 公共工具库
├── Makefile
└── go.mod
```

## 环境要求

本地开发：

- Go 1.24+
- MySQL 8.0+
- Redis 7+
- etcd 3.5+
- Traefik，可选，用作本地网关

本地 k3d/Helm 部署：

- Docker
- k3d
- kubectl
- Helm

## 快速启动：本地开发

### 1. 克隆项目

```sh
git clone https://github.com/cls-cloud/ovra-zero.git
cd ovra-zero
```

### 2. 初始化工具与依赖

```sh
make init
```

### 3. 修改配置

开发配置位于：

```text
etc/dev/auth.yaml
etc/dev/system.yaml
etc/dev/demo.yaml
```

需要按你的本机环境修改：

- MySQL 地址、端口、账号、密码、数据库名
- Redis 地址和密码
- etcd 地址

默认开发配置使用：

```text
MySQL: 127.0.0.1:3306
Redis: 127.0.0.1:6379
etcd:  127.0.0.1:2379
```

初始化 SQL：

```text
bin/sql/ovra_zero.sql
```

### 4. 构建并启动后端

```sh
make build-all
make back-all
```

也可以单独运行某个服务：

```sh
make run-auth
make run-system
make run-demo
```

### 5. 启动 Traefik 网关

Traefik 配置位于：

```text
bin/traefik/traefik.yaml
bin/traefik/dynamic.yaml
```

启动：

```sh
make traefik-run
```

默认网关端口：

```text
http://127.0.0.1:28080
```

## 快速启动：k3d + Helm

项目已提供本地 Helm Chart：

```text
deploy/helm/ovra-zero
```

默认部署内容：

- auth/system/demo 应用服务
- etcd 3 节点 StatefulSet
- Redis 3 节点 Redis Cluster
- Traefik Ingress
- 外部 MySQL EndpointSlice

### 1. 构建 Linux 容器二进制

Apple Silicon / ARM64：

```sh
mkdir -p .deploy/k3d/bin

GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags='-s -w' -tags no_k8s -o .deploy/k3d/bin/app-auth app/auth/auth.go
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags='-s -w' -tags no_k8s -o .deploy/k3d/bin/app-system app/system/system.go
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags='-s -w' -tags no_k8s -o .deploy/k3d/bin/app-demo app/demo/demo.go
```

x86_64 环境将 `GOARCH=arm64` 改为：

```text
GOARCH=amd64
```

### 2. 构建并导入镜像

```sh
docker build --platform linux/arm64 -t ovra-zero:local -f deploy/k3d/Dockerfile .
k3d image import ovra-zero:local -c ovra
```

x86_64 环境将 `--platform linux/arm64` 改为：

```text
--platform linux/amd64
```

### 3. 准备 MySQL

默认 Helm values 使用同一个 k3d Docker 网络里的外部 MySQL 容器：

```sh
docker run -d --name ovra-zero-mysql --network k3d-ovra \
  -e MYSQL_ROOT_PASSWORD='Pl@1221view' \
  -e MYSQL_DATABASE=ovra_zero \
  -v "$PWD/bin/sql/ovra_zero.sql:/docker-entrypoint-initdb.d/ovra_zero.sql:ro" \
  mysql:8.4
```

查看 MySQL 容器 IP：

```sh
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ovra-zero-mysql
```

如果 IP 不是 `values.yaml` 中的 `mysql.external.ip`，安装时覆盖：

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --set mysql.external.ip=<MYSQL_CONTAINER_IP>
```

### 4. 安装 Helm Chart

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --wait \
  --timeout 5m
```

### 5. 验证服务

```sh
kubectl -n ovra-zero get pods,svc,statefulset,ingress
```

验证入口：

```sh
curl -sS http://ovra-zero.localhost:18080/auth/code
```

验证 etcd：

```sh
kubectl -n ovra-zero exec etcd-0 -- etcdctl \
  --endpoints=http://etcd-0.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-1.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-2.etcd-headless.ovra-zero.svc.cluster.local:2379 \
  member list
```

验证 Redis Cluster：

```sh
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster nodes
```

## Helm 与部署文档

详细文档位于：

- [Helm Chart 说明](deploy/helm/ovra-zero/HELM_CHART.md)
- [Redis 与 etcd 安装部署文档](deploy/helm/ovra-zero/INSTALL_REDIS_ETCD.md)
- [Helm Chart README](deploy/helm/ovra-zero/README.md)

其中 `INSTALL_REDIS_ETCD.md` 同时包含 Linux 裸机部署方式。

## 常用 Make 命令

```sh
make help
```

常用命令：

| 命令 | 说明 |
| --- | --- |
| `make init` | 初始化 goctl 与 Go 依赖 |
| `make build-all` | 构建所有服务 |
| `make build-auth` | 构建 auth 服务 |
| `make build-system` | 构建 system 服务 |
| `make build-demo` | 构建 demo 服务 |
| `make run-auth` | 本地运行 auth 服务 |
| `make run-system` | 本地运行 system 服务 |
| `make run-demo` | 本地运行 demo 服务 |
| `make back-all` | 后台启动所有服务 |
| `make gen-all` | 批量生成 API/RPC/DB 代码 |
| `make traefik-run` | 启动本地 Traefik |

## 接口与代码生成

API/RPC 描述文件位于：

```text
desc/auth
desc/system
desc/demo
```

代码生成配置位于：

```text
gen/auth
gen/system
gen/demo
```

生成全部模块：

```sh
make gen-all
```

单模块生成：

```sh
make api-system
make grpc-system
make db-system
```

## 配置说明

开发环境配置：

```text
etc/dev
```

本地 k3d/Helm 配置模板：

```text
deploy/helm/ovra-zero/files/config
```

关键配置：

- `RestConf`：HTTP 服务配置。
- `RpcConf`：gRPC 服务配置。
- `SystemRpc` / `MonitorRpc`：RPC 客户端服务发现配置。
- `JwtAuth`：JWT 密钥与过期时间。
- `ApiDecrypt`：接口加解密配置。
- `Tenant`：多租户忽略表配置。
- `Data.Database`：MySQL 配置。
- `Data.Redis`：Redis 或 Redis Cluster 配置。

Redis Cluster 示例：

```yaml
Data:
  Redis:
    Pass: Pl@1221view
    Host: redis-0.redis-headless.ovra-zero.svc.cluster.local:6379,redis-1.redis-headless.ovra-zero.svc.cluster.local:6379,redis-2.redis-headless.ovra-zero.svc.cluster.local:6379
    Type: cluster
    Tls: false
```

## 当前状态与计划

- [x] 权限系统
- [x] 多租户
- [x] RBAC 权限控制
- [x] 菜单管理
- [x] 系统日志
- [x] 资源管理
- [x] k3d + Helm 本地部署
- [x] etcd 3 节点集群部署
- [x] Redis 3 节点 Cluster 部署
- [ ] 报表大屏可视化

## 许可证

本项目使用 MIT License，详见 [LICENSE](LICENSE)。

## 联系方式 / 技术交流

- Telegram：[@ovra12](https://t.me/ovra12)
- QQ：2579260178，备注 `ovra-zero`
- 邮箱：ut1221@icloud.com，标题格式 `[ovra-zero] : 简要说明问题`
