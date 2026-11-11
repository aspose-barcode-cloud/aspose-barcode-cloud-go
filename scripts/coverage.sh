#!/bin/bash
#
# `make cover`: run the tests with line coverage of the SDK library (./barcode/...)
# and enforce the gate (COVERAGE_THRESHOLD, default 80) with the shared
# scripts/coverage.py, which also writes coverage/coverage-{badge,summary}.json.
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
mkdir -p coverage

go test -v -covermode=atomic -coverpkg=./barcode/... -coverprofile=coverage/coverage.out ./barcode/... ./test

# `python`, or `python3` where only that exists (e.g. a stock Ubuntu 26.04).
python="$(command -v python || command -v python3)"
"${python}" scripts/coverage.py go coverage/coverage.out
