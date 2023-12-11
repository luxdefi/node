#!/usr/bin/env bash

set -euo pipefail

# Directory above this script
LUX_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$LUX_PATH"/scripts/constants.sh

# Ensure execution of fixture unit tests under tests/ but exclude ginkgo tests in tests/e2e and tests/upgrade
go test -shuffle=on -race -timeout=${TIMEOUT:-"120s"} -coverprofile="coverage.out" -covermode="atomic" $(go list ./... | grep -v /mocks | grep -v proto | grep -v tests/e2e | grep -v tests/upgrade)
