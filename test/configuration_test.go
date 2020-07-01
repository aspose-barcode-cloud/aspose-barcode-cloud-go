package test

import (
	"encoding/json"
	"fmt"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
	"github.com/google/uuid"
	"os"
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
	want := "Aspose-Barcode-SDK/0.2006.1/go"
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

	assert.Equal(t, "{\"jwt\":{\"clientId\":\"ClientID\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenURL\",\"accessToken\":\"\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}", string(bytes))
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

func TestNewConfigFromEnvDefaults(t *testing.T) {
	config, err := NewConfigFromEnv(uuid.New().String())
	require.Nil(t, err)

	assert.Equal(t, "", config.JwtConfig.ClientID)
	assert.Equal(t, "", config.JwtConfig.ClientSecret)
	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenURL)

	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.APIConfig.BasePath)
	assert.Equal(t, "", config.APIConfig.Host)
	assert.Equal(t, "Aspose-Barcode-SDK/0.2006.1/go", config.APIConfig.UserAgent)
}

func TestNewConfigFromEnvValues(t *testing.T) {
	uniqPrefix := uuid.New().String()
	var err error
	err = os.Setenv(fmt.Sprintf("%s_JWT_CLIENT_ID", uniqPrefix),
		"jwt client id")
	require.Nil(t, err)
	err = os.Setenv(fmt.Sprintf("%s_JWT_CLIENT_SECRET", uniqPrefix),
		"jwt client secret")
	require.Nil(t, err)

	config, err := NewConfigFromEnv(uniqPrefix)
	require.Nil(t, err)

	assert.Equal(t, "jwt client id", config.JwtConfig.ClientID)
	assert.Equal(t, "jwt client secret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenURL)

	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.APIConfig.BasePath)
	assert.Equal(t, "", config.APIConfig.Host)
	assert.Equal(t, "Aspose-Barcode-SDK/0.2006.1/go", config.APIConfig.UserAgent)
}
