#!/bin/bash
set -euo pipefail

go build $(go list ./... | grep -v snippets)
