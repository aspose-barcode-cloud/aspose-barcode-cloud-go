#!/bin/bash

set -euo pipefail

go fmt ./aspose_barcode_cloud/...
go fmt ./cmd/...
go fmt ./test
