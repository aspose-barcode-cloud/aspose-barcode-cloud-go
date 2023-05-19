package test

import (
	"testing"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationBasePath(t *testing.T) {
	want := "https://api.aspose.cloud/v3.0"
	if got := api.NewConfiguration().BasePath; got != want {
		t.Errorf("NewConfiguration().BasePath = %s; want %s", got, want)
	}
}

func TestNewConfigurationUserAgent(t *testing.T) {
	want := "Aspose-Barcode-SDK/0.2305.0/go"
	if got := api.NewConfiguration().UserAgent; got != want {
		t.Errorf("NewConfiguration().UserAgent = %s; want %s", got, want)
	}
}

func TestAddDefaultHeader(t *testing.T) {
	config := api.NewConfiguration()

	assert.Equal(t, 0, len(config.DefaultHeader))
	config.AddDefaultHeader("CustomHeader", "Value")
	assert.Equal(t, 1, len(config.DefaultHeader))
	assert.Equal(t, "Value", config.DefaultHeader["CustomHeader"])
}
