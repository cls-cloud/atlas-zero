# Ovra-Zero Helm chart

This chart deploys the local k3d version of Ovra-Zero.

By default it creates:

- a 3-member etcd StatefulSet
- a 3-node Redis Cluster StatefulSet with 3 masters
- auth, system, and demo application Deployments

## Build and import the app image

```sh
mkdir -p .deploy/k3d/bin
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags='-s -w' -tags no_k8s -o .deploy/k3d/bin/app-auth app/auth/auth.go
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags='-s -w' -tags no_k8s -o .deploy/k3d/bin/app-system app/system/system.go
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags='-s -w' -tags no_k8s -o .deploy/k3d/bin/app-demo app/demo/demo.go
docker build --platform linux/arm64 -t ovra-zero:local -f deploy/k3d/Dockerfile .
k3d image import ovra-zero:local -c ovra
```

## MySQL for local k3d

The default values expect a MySQL container on the `k3d-ovra` Docker network:

```sh
docker run -d --name ovra-zero-mysql --network k3d-ovra \
  -e MYSQL_ROOT_PASSWORD='Pl@1221view' \
  -e MYSQL_DATABASE=ovra_zero \
  -v "$PWD/bin/sql/ovra_zero.sql:/docker-entrypoint-initdb.d/ovra_zero.sql:ro" \
  mysql:8.4
```

Set `mysql.external.ip` in `values.yaml` if Docker gives that container a different IP.

## Install

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --wait \
  --timeout 5m
```

Then open:

```text
http://ovra-zero.localhost:18080/auth/code
```

## Verify clusters

```sh
kubectl -n ovra-zero exec etcd-0 -- etcdctl \
  --endpoints=http://etcd-0.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-1.etcd-headless.ovra-zero.svc.cluster.local:2379,http://etcd-2.etcd-headless.ovra-zero.svc.cluster.local:2379 \
  member list

kubectl -n ovra-zero exec redis-0 -- redis-cli -a 'Pl@1221view' cluster info
```

For detailed Kubernetes and bare-metal Linux installation steps, see:

```text
deploy/helm/ovra-zero/INSTALL_REDIS_ETCD.md
```

For detailed Helm chart structure, values, upgrade, and troubleshooting notes, see:

```text
deploy/helm/ovra-zero/HELM_CHART.md
```
