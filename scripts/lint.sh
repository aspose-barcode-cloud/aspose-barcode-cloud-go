#!/bin/bash
set -euo pipefail

golint -set_exit_status ./examples/ ./test/
