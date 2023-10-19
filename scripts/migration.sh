#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

kubectl delete -f k8s/migration.yaml || true

# wait for the pod to be deleted
echo "Waiting for job to be deleted..."
kubectl wait --for=delete job/migration -n marketplace

kubectl apply -f k8s/migration.yaml

# wait for the migration to complete
echo "Waiting for migration to complete..."
kubectl wait --for=condition=complete job/migration -n marketplace
