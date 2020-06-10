package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTokenSource(t *testing.T) {
	config, err := NewConfig("configuration.json")
	require.Nil(t, err)

	source := config.JwtConfig.TokenSource(context.Background())
	token, err := source.Token()
	require.Nil(t, err)

	assert.NotEmpty(t, token.AccessToken)
	assert.Greater(t, token.Expiry.Unix(), time.Now().Unix())
}
