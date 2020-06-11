#!/bin/bash

set -euo pipefail

go clean ./barcode/...
go clean ./cmd/...
go clean ./test
