#!/bin/bash
set -euo pipefail

"$(go env GOPATH)/bin/goimports" -w .
