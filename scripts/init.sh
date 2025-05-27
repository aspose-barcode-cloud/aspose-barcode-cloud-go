#!/bin/bash
set -euo pipefail

go get -v -t ./...

go install golang.org/x/tools/cmd/goimports@v0.33.0
