#!/bin/bash
set -euo pipefail

go get -v -t -d ./...
go install golang.org/x/tools/cmd/goimports@v0.16.0
go install honnef.co/go/tools/cmd/staticcheck@v0.3.3
