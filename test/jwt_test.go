package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTokenSource(t *testing.T) {
	source := TestConfig.JwtConfig.TokenSource(context.Background())
	token, err := source.Token()
	require.Nil(t, err)

	assert.NotEmpty(t, token.AccessToken)
	assert.Greater(t, token.Expiry.Unix(), time.Now().Unix())
}
