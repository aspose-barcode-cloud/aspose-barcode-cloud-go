package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/url"
	"time"
)

const defaultGrantType = "client_credentials"
const defaultTokenURL = "https://api.aspose.cloud/connect/token"

type Config struct {
	ClientId     string `json:"ClientId"`
	ClientSecret string `json:"ClientSecret"`
	TokenURL     string `json:"TokenURL"`
}

func NewConfig(clientID string, clientSecret string) Config {
	config := Config{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     defaultTokenURL,
	}

	return config
}

// TokenSource returns a JWT TokenSource using the configuration
// in c and the HTTP client from the provided context.
func (c *Config) TokenSource(ctx context.Context) oauth2.TokenSource {
	return oauth2.ReuseTokenSource(nil, jwtSource{ctx, c})
}

// jwtSource is a source that always does a signed JWT request for a token.
// It should typically be wrapped with a reuseTokenSource.
type jwtSource struct {
	ctx  context.Context
	conf *Config
}

func (js jwtSource) Token() (*oauth2.Token, error) {
	v := url.Values{}
	v.Set("grant_type", defaultGrantType)
	v.Set("client_id", js.conf.ClientId)
	v.Set("client_secret", js.conf.ClientSecret)
	hc := oauth2.NewClient(js.ctx, nil)
	resp, err := hc.PostForm(js.conf.TokenURL, v)
	if err != nil {
		return nil, fmt.Errorf("jwt: cannot fetch token: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("jwt: cannot read token: %v", err)
	}
	if c := resp.StatusCode; c < 200 || c > 299 {
		return nil, &oauth2.RetrieveError{
			Response: resp,
			Body:     body,
		}
	}
	// tokenRes is the JSON response body.
	var tokenRes struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"` // relative seconds from now
	}
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return nil, fmt.Errorf("oauth2: cannot unmarshal token: %v", err)
	}
	token := &oauth2.Token{
		AccessToken: tokenRes.AccessToken,
		TokenType:   tokenRes.TokenType,
	}

	if secs := tokenRes.ExpiresIn; secs > 0 {
		token.Expiry = time.Now().Add(time.Duration(secs) * time.Second)
	}

	return token, nil
}
