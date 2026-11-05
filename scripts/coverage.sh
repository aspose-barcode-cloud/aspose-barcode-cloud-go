#!/bin/bash
#
# `make cover`: run the tests with line coverage of the SDK library (./barcode/...)
# and enforce the gate (COVERAGE_THRESHOLD, default 80) with the shared
# scripts/coverage.py, which also writes coverage/coverage-{badge,summary}.json.
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
mkdir -p coverage

go test -v -covermode=atomic -coverpkg=./barcode/... -coverprofile=coverage/coverage.out ./barcode/... ./test

python3 scripts/coverage.py go coverage/coverage.out
