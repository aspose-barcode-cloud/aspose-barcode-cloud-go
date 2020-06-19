package test

import (
	"encoding/json"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
	"testing"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfigurationBasePath(t *testing.T) {
	want := "https://api.aspose.cloud/v3.0"
	if got := api.NewConfiguration().BasePath; got != want {
		t.Errorf("NewConfiguration().BasePath = %s; want %s", got, want)
	}
}

func TestNewConfigurationUserAgent(t *testing.T) {
	want := "Aspose-Barcode-SDK/0.20.6.1/go"
	if got := api.NewConfiguration().UserAgent; got != want {
		t.Errorf("NewConfiguration().UserAgent = %s; want %s", got, want)
	}
}

func TestNewConfig(t *testing.T) {
	config, err := NewConfig("configuration.example.json")
	require.Nil(t, err)

	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenURL)
	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.APIConfig.BasePath)
}

func TestConfigMarshal(t *testing.T) {
	config := Config{
		JwtConfig: jwt.Config{
			ClientID:     "ClientID",
			ClientSecret: "ClientSecret",
			TokenURL:     "TokenURL",
		},
		APIConfig: api.Configuration{
			BasePath:  "BasePath",
			Host:      "Host",
			UserAgent: "UserAgent",
		},
	}

	bytes, err := json.Marshal(config)
	require.Nil(t, err)

	assert.Equal(t, "{\"jwt\":{\"clientId\":\"ClientID\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenURL\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}", string(bytes))
}

func TestNewConfigFromJson(t *testing.T) {
	config, err := NewConfigFromJSON([]byte("{\"jwt\":{\"clientId\":\"ClientID\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenURL\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}"))
	require.Nil(t, err)

	assert.Equal(t, "ClientID", config.JwtConfig.ClientID)
	assert.Equal(t, "ClientSecret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "TokenURL", config.JwtConfig.TokenURL)

	assert.Equal(t, "BasePath", config.APIConfig.BasePath)
	assert.Equal(t, "Host", config.APIConfig.Host)
	assert.Equal(t, "UserAgent", config.APIConfig.UserAgent)
}
