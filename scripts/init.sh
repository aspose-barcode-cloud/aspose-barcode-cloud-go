#!/bin/bash
set -euo pipefail

go get -v -t -d ./...
# Versions compatible with Go 1.17
go install golang.org/x/tools/cmd/goimports@v0.16.0
