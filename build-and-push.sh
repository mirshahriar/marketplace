#!/bin/bash

set -e

go build
echo "Docker build with tag: latest"
docker build --platform linux/amd64 -t mirshahriar/marketplace:latest .

echo "Skipping docker push for now"
