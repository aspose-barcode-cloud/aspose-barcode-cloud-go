package test

import (
	"context"
	"testing"
	"time"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenSource(t *testing.T) {
	testConfig, err := NewTestConfig(TestConfigurationFile, TestEnvPrefix)
	require.Nil(t, err)

	source := testConfig.JwtConfig.TokenSource(context.Background())
	token, err := source.Token()
	require.Nil(t, err)

	assert.NotEmpty(t, token.AccessToken)
	assert.Greater(t, token.Expiry.Unix(), time.Now().Unix())
}

func TestValidateAccessToken(t *testing.T) {
	config := jwt.NewConfig("", "")
	config.AccessToken = "access token"

	err := config.Validate()
	require.Nil(t, err)
}

func TestValidateTokenURL(t *testing.T) {
	config := jwt.NewConfig("id", "secret")
	config.TokenURL = ""

	err := config.Validate()
	require.EqualError(t, err, "incorrect TokenURL value ''")
}

func TestValidateNoClientID(t *testing.T) {
	err := jwt.NewConfig("", "secret").Validate()

	require.EqualError(t, err, "incorrect ClientID value ''")
}

func TestValidateNoClientSecret(t *testing.T) {
	err := jwt.NewConfig("id", "").Validate()

	require.EqualError(t, err, "incorrect ClientSecret value ''")
}
