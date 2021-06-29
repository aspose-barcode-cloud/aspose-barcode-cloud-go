package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTestConfigFromFile(t *testing.T) {
	config, err := NewTestConfig("configuration.example.json", TestEnvPrefix)
	require.Nil(t, err)

	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenURL)
	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.APIConfig.BasePath)
}

func TestNewTestConfigFileNotExists(t *testing.T) {
	uniqPrefix := uuid.New().String()
	err := os.Setenv(fmt.Sprintf("%s_JWT_ACCESS_TOKEN", uniqPrefix),
		"jwt access token")
	require.Nil(t, err)

	config, err := NewTestConfig("not a file", uniqPrefix)
	require.Nil(t, err)

	assert.Equal(t, "jwt access token", config.JwtConfig.AccessToken)
}

func TestNewConfigFromJson(t *testing.T) {
	config, err := newConfigFromJSON([]byte("{\"jwt\":{\"clientId\":\"ClientID\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenURL\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}"))
	require.Nil(t, err)

	assert.Equal(t, "ClientID", config.JwtConfig.ClientID)
	assert.Equal(t, "ClientSecret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "TokenURL", config.JwtConfig.TokenURL)

	assert.Equal(t, "BasePath", config.APIConfig.BasePath)
	assert.Equal(t, "Host", config.APIConfig.Host)
	assert.Equal(t, "UserAgent", config.APIConfig.UserAgent)
}

func TestNewConfigFromEnvDefaults(t *testing.T) {
	config, err := newConfigFromEnv(uuid.New().String())
	require.Nil(t, err)

	assert.Equal(t, "", config.JwtConfig.ClientID)
	assert.Equal(t, "", config.JwtConfig.ClientSecret)
	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenURL)

	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.APIConfig.BasePath)
	assert.Equal(t, "", config.APIConfig.Host)
	assert.Equal(t, "Aspose-Barcode-SDK/0.2106.0/go", config.APIConfig.UserAgent)
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

	config, err := newConfigFromEnv(uniqPrefix)
	require.Nil(t, err)

	assert.Equal(t, "jwt client id", config.JwtConfig.ClientID)
	assert.Equal(t, "jwt client secret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "https://api.aspose.cloud/connect/token", config.JwtConfig.TokenURL)

	assert.Equal(t, "https://api.aspose.cloud/v3.0", config.APIConfig.BasePath)
	assert.Equal(t, "", config.APIConfig.Host)
	assert.Equal(t, "Aspose-Barcode-SDK/0.2106.0/go", config.APIConfig.UserAgent)
}

func TestConfigMarshal(t *testing.T) {
	config := Config{
		JwtConfig: jwt.Config{
			ClientID:     "ClientID",
			ClientSecret: "ClientSecret",
			TokenURL:     "TokenURL",
		},
		APIConfig: barcode.Configuration{
			BasePath:  "BasePath",
			Host:      "Host",
			UserAgent: "UserAgent",
		},
	}

	bytes, err := json.Marshal(config)
	require.Nil(t, err)

	assert.Equal(t, "{\"jwt\":{\"clientId\":\"ClientID\",\"clientSecret\":\"ClientSecret\",\"tokenUrl\":\"TokenURL\",\"accessToken\":\"\"},\"api\":{\"basePath\":\"BasePath\",\"host\":\"Host\",\"userAgent\":\"UserAgent\"}}", string(bytes))
}
