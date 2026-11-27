package binarydecode_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/lint/binarydecode"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), binarydecode.Analyzer, "decode")
}
