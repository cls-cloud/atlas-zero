# Ovra-Zero Helm chart

This chart deploys Ovra-Zero to a standard Kubernetes cluster.

By default it creates:

- auth, system, and demo application Deployments
- a 3-member etcd StatefulSet
- a 3-node Redis Cluster StatefulSet
- an internal MySQL Deployment and Service
- an optional Ingress

## Install

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --wait \
  --timeout 5m
```

## Image

Set the application image repository and tag for your registry:

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --set image.repository=registry.example.com/ovra-zero \
  --set image.tag=1.0.0
```

## External MySQL

The chart uses the internal MySQL Deployment by default. To point the `mysql`
Service at an existing database endpoint, disable the internal deployment and
provide an endpoint IP:

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --set mysql.internal.enabled=false \
  --set mysql.external.enabled=true \
  --set mysql.external.ip=10.0.0.10
```

## Ingress

Set your ingress class and host as needed:

```sh
helm upgrade --install ovra-zero deploy/helm/ovra-zero \
  --namespace ovra-zero \
  --create-namespace \
  --set ingress.className=nginx \
  --set ingress.host=ovra-zero.example.com
```

## Verify

```sh
kubectl -n ovra-zero get pods,svc,ingress
kubectl -n ovra-zero rollout status deployment/auth
kubectl -n ovra-zero rollout status deployment/system
kubectl -n ovra-zero rollout status deployment/demo
```

For Redis and etcd installation notes, see:

```text
deploy/helm/ovra-zero/INSTALL_REDIS_ETCD.md
```
