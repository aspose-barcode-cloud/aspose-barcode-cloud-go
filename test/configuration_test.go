package test

import (
	"testing"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud"
)

func TestNewConfigurationBasePath(t *testing.T) {
	want := "https://api.aspose.cloud/v3.0"
	if got := api.NewConfiguration().BasePath; got != want {
		t.Errorf("NewConfiguration().BasePath = %s; want %s", got, want)
	}
}

func TestNewConfigurationUserAgent(t *testing.T) {
	want := "Aspose-Barcode-SDK/20.5.0/go"
	if got := api.NewConfiguration().UserAgent; got != want {
		t.Errorf("NewConfiguration().UserAgent = %s; want %s", got, want)
	}
}
