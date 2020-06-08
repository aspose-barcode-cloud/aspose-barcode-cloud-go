#!/bin/bash

set -euo pipefail

go clean ./aspose_barcode_cloud/...
go clean ./cmd/...
go clean ./test
