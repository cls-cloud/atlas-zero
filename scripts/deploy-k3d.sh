#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

MODE="${1:-deploy}"
CLUSTER="${CLUSTER:-ovra}"
NAMESPACE="${NAMESPACE:-ovra-zero}"
RELEASE="${RELEASE:-ovra-zero}"
IMAGE_REPOSITORY="${IMAGE_REPOSITORY:-ovra-zero}"
IMAGE_TAG="${IMAGE_TAG:-local}"
PLATFORM="${PLATFORM:-linux/arm64}"
GOOS_TARGET="${GOOS_TARGET:-linux}"
GOARCH_TARGET="${GOARCH_TARGET:-arm64}"
GOCACHE_DIR="${GOCACHE_DIR:-/private/tmp/ovra-zero-go-build}"
HELM_TIMEOUT="${HELM_TIMEOUT:-5m}"
MYSQL_CONTAINER="${MYSQL_CONTAINER:-ovra-zero-mysql}"
MYSQL_EXTERNAL_IP="${MYSQL_EXTERNAL_IP:-}"
K3D_NODE_READY_TIMEOUT="${K3D_NODE_READY_TIMEOUT:-120}"
K3D_IMPORT_RETRIES="${K3D_IMPORT_RETRIES:-3}"
K3D_AUTO_RESTART_NOTREADY_NODES="${K3D_AUTO_RESTART_NOTREADY_NODES:-true}"

BIN_DIR=".deploy/k3d/bin"

list_k3d_nodes() {
  docker ps --filter "name=k3d-${CLUSTER}-" --format '{{.Names}}' \
    | grep -E "^k3d-${CLUSTER}-(server|agent)-[0-9]+$" || true
}

restart_notready_k3d_nodes() {
  if [[ "$K3D_AUTO_RESTART_NOTREADY_NODES" != "true" ]]; then
    return 0
  fi

  local nodes
  nodes="$(kubectl get nodes --no-headers 2>/dev/null \
    | awk -v prefix="k3d-${CLUSTER}-" '$1 ~ "^" prefix "(server|agent)-[0-9]+$" && $2 != "Ready" {print $1}' || true)"

  if [[ -z "$nodes" ]]; then
    return 0
  fi

  echo "==> Restarting NotReady k3d nodes"
  while IFS= read -r node; do
    if [[ -n "$node" ]]; then
      echo "    restarting ${node}"
      docker restart "$node" >/dev/null
    fi
  done <<< "$nodes"
}

wait_for_k3d_nodes() {
  local deadline
  deadline=$((SECONDS + K3D_NODE_READY_TIMEOUT))

  echo "==> Waiting for k3d Kubernetes nodes"
  while (( SECONDS < deadline )); do
    local not_ready
    not_ready="$(kubectl get nodes --no-headers 2>/dev/null \
      | awk -v prefix="k3d-${CLUSTER}-" '$1 ~ "^" prefix "(server|agent)-[0-9]+$" && $2 != "Ready" {print $1}' || true)"

    if [[ -z "$not_ready" ]]; then
      return 0
    fi

    sleep 3
  done

  echo "timed out waiting for k3d Kubernetes nodes" >&2
  kubectl get nodes || true
  exit 1
}

wait_for_k3d_containerd() {
  local deadline
  deadline=$((SECONDS + K3D_NODE_READY_TIMEOUT))

  echo "==> Waiting for k3d node containerd sockets"
  while (( SECONDS < deadline )); do
    local pending=0
    local nodes
    nodes="$(list_k3d_nodes)"

    if [[ -z "$nodes" ]]; then
      echo "no k3d nodes found for cluster ${CLUSTER}" >&2
      exit 1
    fi

    while IFS= read -r node; do
      if [[ -z "$node" ]]; then
        continue
      fi
      if ! docker exec "$node" sh -c 'test -S /run/k3s/containerd/containerd.sock && ctr version >/dev/null 2>&1'; then
        pending=1
        break
      fi
    done <<< "$nodes"

    if [[ "$pending" -eq 0 ]]; then
      return 0
    fi

    sleep 3
  done

  echo "timed out waiting for k3d node containerd sockets" >&2
  docker ps --filter "name=k3d-${CLUSTER}-" --format 'table {{.Names}}\t{{.Status}}'
  exit 1
}

import_image_to_k3d() {
  local attempt=1

  restart_notready_k3d_nodes
  wait_for_k3d_nodes
  wait_for_k3d_containerd
  while (( attempt <= K3D_IMPORT_RETRIES )); do
    if k3d image import "${IMAGE_REPOSITORY}:${IMAGE_TAG}" -c "$CLUSTER"; then
      return 0
    fi

    echo "k3d image import failed, retrying (${attempt}/${K3D_IMPORT_RETRIES})" >&2
    sleep 5
    wait_for_k3d_containerd
    attempt=$((attempt + 1))
  done

  echo "failed to import image into k3d cluster ${CLUSTER}" >&2
  exit 1
}

usage() {
  cat <<EOF
Usage:
  $0 [deploy|redeploy]

Modes:
  deploy    Build image and run helm upgrade --install. This is the default.
  redeploy  Build image, uninstall the existing Helm release, then install it again.

Environment:
  CLUSTER=${CLUSTER}
  NAMESPACE=${NAMESPACE}
  RELEASE=${RELEASE}
  IMAGE_REPOSITORY=${IMAGE_REPOSITORY}
  IMAGE_TAG=${IMAGE_TAG}
  MYSQL_CONTAINER=${MYSQL_CONTAINER}
  K3D_NODE_READY_TIMEOUT=${K3D_NODE_READY_TIMEOUT}
  K3D_IMPORT_RETRIES=${K3D_IMPORT_RETRIES}
  K3D_AUTO_RESTART_NOTREADY_NODES=${K3D_AUTO_RESTART_NOTREADY_NODES}
EOF
}

case "$MODE" in
  deploy | redeploy)
    ;;
  -h | --help | help)
    usage
    exit 0
    ;;
  *)
    usage >&2
    exit 1
    ;;
esac

echo "==> Building linux binaries (${GOOS_TARGET}/${GOARCH_TARGET})"
mkdir -p "$BIN_DIR"
GOCACHE="$GOCACHE_DIR" GOOS="$GOOS_TARGET" GOARCH="$GOARCH_TARGET" CGO_ENABLED=0 \
  go build -ldflags="-s -w" -tags no_k8s -o "$BIN_DIR/app-auth" app/auth/auth.go
GOCACHE="$GOCACHE_DIR" GOOS="$GOOS_TARGET" GOARCH="$GOARCH_TARGET" CGO_ENABLED=0 \
  go build -ldflags="-s -w" -tags no_k8s -o "$BIN_DIR/app-system" app/system/system.go
GOCACHE="$GOCACHE_DIR" GOOS="$GOOS_TARGET" GOARCH="$GOARCH_TARGET" CGO_ENABLED=0 \
  go build -ldflags="-s -w" -tags no_k8s -o "$BIN_DIR/app-demo" app/demo/demo.go

echo "==> Building image ${IMAGE_REPOSITORY}:${IMAGE_TAG} (${PLATFORM})"
docker build --platform "$PLATFORM" -t "${IMAGE_REPOSITORY}:${IMAGE_TAG}" -f deploy/k3d/Dockerfile .

echo "==> Importing image into k3d cluster ${CLUSTER}"
import_image_to_k3d

if [[ -z "$MYSQL_EXTERNAL_IP" ]]; then
  MYSQL_EXTERNAL_IP="$(docker inspect "$MYSQL_CONTAINER" --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}')"
fi

if [[ -z "$MYSQL_EXTERNAL_IP" ]]; then
  echo "failed to detect MySQL container IP; set MYSQL_EXTERNAL_IP or MYSQL_CONTAINER" >&2
  exit 1
fi

echo "==> Using MySQL endpoint ${MYSQL_EXTERNAL_IP}:3306"

if [[ "$MODE" == "redeploy" ]]; then
  echo "==> Reinstalling Helm release ${RELEASE} in namespace ${NAMESPACE}"
  if helm -n "$NAMESPACE" status "$RELEASE" >/dev/null 2>&1; then
    helm -n "$NAMESPACE" uninstall "$RELEASE" --wait --timeout "$HELM_TIMEOUT"
    kubectl -n "$NAMESPACE" wait --for=delete pod \
      -l "app.kubernetes.io/instance=${RELEASE}" \
      --timeout="$HELM_TIMEOUT" >/dev/null 2>&1 || true
  else
    echo "==> Release ${RELEASE} does not exist, continuing with a fresh install"
  fi
fi

echo "==> Deploying Helm release ${RELEASE} in namespace ${NAMESPACE}"
helm upgrade --install "$RELEASE" deploy/helm/ovra-zero \
  --namespace "$NAMESPACE" \
  --create-namespace \
  --wait \
  --timeout "$HELM_TIMEOUT" \
  --set image.repository="$IMAGE_REPOSITORY" \
  --set image.tag="$IMAGE_TAG" \
  --set config.mountPath="/app/etc/config" \
  --set ingress.className="traefik" \
  --set ingress.host="ovra-zero.localhost" \
  --set mysql.internal.enabled=false \
  --set mysql.external.enabled=true \
  --set mysql.external.ip="$MYSQL_EXTERNAL_IP"

echo "==> Waiting for app rollouts"
kubectl -n "$NAMESPACE" rollout status deployment/auth --timeout="$HELM_TIMEOUT"
kubectl -n "$NAMESPACE" rollout status deployment/system --timeout="$HELM_TIMEOUT"
kubectl -n "$NAMESPACE" rollout status deployment/demo --timeout="$HELM_TIMEOUT"

echo "==> Current pods"
kubectl -n "$NAMESPACE" get pods
