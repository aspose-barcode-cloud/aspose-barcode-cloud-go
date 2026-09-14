#!/bin/bash
set -euo pipefail

go get -v -t ./...

go install golang.org/x/tools/cmd/goimports@v0.33.0
go install honnef.co/go/tools/cmd/staticcheck@v0.6.1
