#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <tag>"
  exit 0
fi

go build
echo "Docker build with tag: $1"
docker build --platform linux/amd64 -t mirshahriar/github.com/mirshahriar/marketplace:latest .
