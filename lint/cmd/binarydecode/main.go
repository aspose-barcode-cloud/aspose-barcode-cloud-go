// Command binarydecode runs the binarydecode analyzer as a go vet tool:
//
//	go build -C lint -o bin/binarydecode ./cmd/binarydecode
//	go vet -vettool="$(pwd)/lint/bin/binarydecode" ./barcode/...
package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/lint/binarydecode"
)

func main() {
	unitchecker.Main(binarydecode.Analyzer)
}
