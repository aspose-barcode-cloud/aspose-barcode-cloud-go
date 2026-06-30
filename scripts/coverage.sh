#!/bin/bash
set -euo pipefail

COVERAGE_PROFILE="${COVERAGE_PROFILE:-coverage.out}"
COVERAGE_THRESHOLD="${COVERAGE_THRESHOLD:-80}"

go test -v -covermode=atomic -coverpkg=./barcode/... -coverprofile="${COVERAGE_PROFILE}" ./barcode/... ./test

coverage="$(go tool cover -func="${COVERAGE_PROFILE}" | awk '/^total:/ {gsub(/%/, "", $3); print $3}')"

awk -v coverage="${coverage}" -v threshold="${COVERAGE_THRESHOLD}" '
BEGIN {
	if ((coverage + 0) < (threshold + 0)) {
		printf("library coverage %.1f%% is below %.1f%%\n", coverage, threshold)
		exit 1
	}

	printf("library coverage %.1f%% meets %.1f%% threshold\n", coverage, threshold)
}
'
