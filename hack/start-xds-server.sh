#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

make clean
make build
./build/bin/envoy-xds-server -watchDirectoryFileName=test/configs