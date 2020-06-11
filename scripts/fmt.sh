#!/bin/bash

set -euo pipefail

go fmt ./barcode/...
go fmt ./cmd/...
go fmt ./test
