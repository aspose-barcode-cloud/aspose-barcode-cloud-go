package test

import (
	"regexp"
	"testing"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationBasePath(t *testing.T) {
	want := "https://api.aspose.cloud/v4.0"
	if got := api.NewConfiguration().BasePath; got != want {
		t.Errorf("NewConfiguration().BasePath = %s; want %s", got, want)
	}
}

func TestNewConfigurationUserAgentWithRegex(t *testing.T) {
	wantPattern := `^Aspose-Barcode-SDK/4\.\d{4}\.\d+/go$`
	if got := api.NewConfiguration().UserAgent; !regexp.MustCompile(wantPattern).MatchString(got) {
		t.Errorf("NewConfiguration().UserAgent = %s; want pattern %s", got, wantPattern)
	}
}

func TestAddDefaultHeader(t *testing.T) {
	config := api.NewConfiguration()

	assert.Equal(t, 0, len(config.DefaultHeader))
	config.AddDefaultHeader("CustomHeader", "Value")
	assert.Equal(t, 1, len(config.DefaultHeader))
	assert.Equal(t, "Value", config.DefaultHeader["CustomHeader"])
}
