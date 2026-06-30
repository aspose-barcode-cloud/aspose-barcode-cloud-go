package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTestConfigFromFile(t *testing.T) {
	config, err := NewTestConfig("configuration.example.json", TestEnvPrefix)
	require.Nil(t, err)

	assert.Equal(t, "https://id.aspose.cloud/connect/token", config.JwtConfig.TokenURL)
	assert.Equal(t, "https://api.aspose.cloud/v4.0", config.APIConfig.BasePath)
}

func TestNewTestConfigFileNotExists(t *testing.T) {
	uniqPrefix := uuid.New().String()
	err := os.Setenv(fmt.Sprintf("%s_JWT_ACCESS_TOKEN", uniqPrefix),
		"jwt access token")
	require.Nil(t, err)
	t.Cleanup(func() {
		require.NoError(t, os.Unsetenv(fmt.Sprintf("%s_JWT_ACCESS_TOKEN", uniqPrefix)))
	})

	config, err := NewTestConfig("not a file", uniqPrefix)
	require.Nil(t, err)

	assert.Equal(t, "jwt access token", config.JwtConfig.AccessToken)
}

func TestNewTestConfigEnvOverridesFile(t *testing.T) {
	uniqPrefix := uuid.New().String()
	configFile := filepath.Join(t.TempDir(), "configuration.json")
	err := os.WriteFile(configFile, []byte(`{"jwt":{"clientId":"file id","clientSecret":"file secret"},"api":{"basePath":"https://file.example"}}`), 0600)
	require.Nil(t, err)

	err = os.Setenv(fmt.Sprintf("%s_JWT_ACCESS_TOKEN", uniqPrefix), "env access token")
	require.Nil(t, err)
	err = os.Setenv(fmt.Sprintf("%s_API_BASE_PATH", uniqPrefix), "https://env.example")
	require.Nil(t, err)
	t.Cleanup(func() {
		require.NoError(t, os.Unsetenv(fmt.Sprintf("%s_JWT_ACCESS_TOKEN", uniqPrefix)))
		require.NoError(t, os.Unsetenv(fmt.Sprintf("%s_API_BASE_PATH", uniqPrefix)))
	})

	config, err := NewTestConfig(configFile, uniqPrefix)
	require.Nil(t, err)

	assert.Equal(t, "file id", config.JwtConfig.ClientID)
	assert.Equal(t, "file secret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "env access token", config.JwtConfig.AccessToken)
	assert.Equal(t, "https://env.example", config.APIConfig.BasePath)
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
	assert.Equal(t, "https://id.aspose.cloud/connect/token", config.JwtConfig.TokenURL)

	assert.Equal(t, "https://api.aspose.cloud/v4.0", config.APIConfig.BasePath)
	assert.Equal(t, "", config.APIConfig.Host)
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
	t.Cleanup(func() {
		require.NoError(t, os.Unsetenv(fmt.Sprintf("%s_JWT_CLIENT_ID", uniqPrefix)))
		require.NoError(t, os.Unsetenv(fmt.Sprintf("%s_JWT_CLIENT_SECRET", uniqPrefix)))
	})

	config, err := newConfigFromEnv(uniqPrefix)
	require.Nil(t, err)

	assert.Equal(t, "jwt client id", config.JwtConfig.ClientID)
	assert.Equal(t, "jwt client secret", config.JwtConfig.ClientSecret)
	assert.Equal(t, "https://id.aspose.cloud/connect/token", config.JwtConfig.TokenURL)

	assert.Equal(t, "https://api.aspose.cloud/v4.0", config.APIConfig.BasePath)
	assert.Equal(t, "", config.APIConfig.Host)
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
