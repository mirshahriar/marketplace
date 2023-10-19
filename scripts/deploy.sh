#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

kubectl create namespace marketplace || true
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# wait for the deployment to be ready
echo "Waiting for deployment to be ready..."
kubectl rollout status deployment/server -n github.com/mirshahriar/marketplace