#!/bin/bash
set -euo pipefail

go list -m -u all
go get -u ./...
