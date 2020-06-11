#!/bin/bash

set -euo pipefail

go vet ./barcode/...
go vet ./cmd/...
go vet ./test
