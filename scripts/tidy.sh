#!/bin/bash
set -euo pipefail

go mod tidy -go=1.16 && go mod tidy -go=1.17
