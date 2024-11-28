#!/bin/bash
set -euo pipefail

#Run all tests except for snippets that have duplicate function names
go test -v $(go list ./... | grep -v snippets)
