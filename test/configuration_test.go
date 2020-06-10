package test

import (
	"encoding/json"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/jwt"
	"testing"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud"
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
	want := "Aspose-Barcode-SDK/20.5.0/go"
	if got := api.NewConfiguration().UserAgent; got != want {
		t.Errorf("NewConfiguration().UserAgent = %s; want %s", got, want)
	}
}

func TestNewConfig(t *testing.T) {
	config, err := NewConfig("configuration.example.json")
	require.Nil(t, err)

	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenUrl)
	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.ApiConfig.BasePath)
}

func TestConfigMarshal(t *testing.T) {
	config := Config{
		JwtConfig: jwt.Config{
			ClientId:     "ClientId",
			ClientSecret: "ClientSecret",
			TokenUrl:     "TokenUrl",
		},
		ApiConfig: api.Configuration{
			BasePath:  "BasePath",
			Host:      "Host",
			UserAgent: "UserAgent",
		},
	}

	bytes, err := json.Marshal(config)
	require.Nil(t, err)

	assert.Equal(t, "{\"jwt\":{\"clientId\":\"ClientId\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenUrl\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}", string(bytes))
}

func TestNewConfigFromJson(t *testing.T) {
	config, err := NewConfigFromJson([]byte("{\"jwt\":{\"clientId\":\"ClientId\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenUrl\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}"))
	require.Nil(t, err)

	assert.Equal(t, "ClientId", config.JwtConfig.ClientId)
	assert.Equal(t, "ClientSecret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "TokenUrl", config.JwtConfig.TokenUrl)

	assert.Equal(t, "BasePath", config.ApiConfig.BasePath)
	assert.Equal(t, "Host", config.ApiConfig.Host)
	assert.Equal(t, "UserAgent", config.ApiConfig.UserAgent)
}
