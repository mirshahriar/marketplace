#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

tmp_dir=$(mktemp -d -t goinstall_XXXXXXXXXX)
function clean {
  rm -rf "${tmp_dir}"
}
trap clean EXIT

rm "${GOBIN}/${2}"* || true

cd "${tmp_dir}"

# create a new module in the tmp directory
go mod init fake/mod

# install the golang module specified as the first argument
go install "${1}@${3}"
mv "${GOBIN}/${2}" "${GOBIN}/${2}-${3}"
ln -sf "${GOBIN}/${2}-${3}" "${GOBIN}/${2}"
