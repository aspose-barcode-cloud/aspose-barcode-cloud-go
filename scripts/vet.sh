#!/bin/bash

set -euo pipefail

go vet ./aspose_barcode_cloud/...
go vet ./cmd/...
go vet ./test
