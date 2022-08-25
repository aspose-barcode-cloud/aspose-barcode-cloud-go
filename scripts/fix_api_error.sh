#!/bin/bash
set -euo pipefail

sed -i -Ee's/InnerError([[:space:]]+)ApiError/InnerError\1\*ApiError/' barcode/model_api_error.go
