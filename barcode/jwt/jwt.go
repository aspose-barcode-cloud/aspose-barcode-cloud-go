package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

// Config - JWT auth configuration
type Config struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	TokenURL     string `json:"tokenUrl"`
	AccessToken  string `json:"accessToken"`
}

// NewConfig creates new Config
func NewConfig(clientID string, clientSecret string) *Config {
	config := Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.aspose.cloud/connect/token",
	}

	return &config
}

func (c *Config) Validate() error {
	if len(c.AccessToken) > 0 {
		return nil
	}

	if len(c.ClientID) == 0 {
		return fmt.Errorf("incorrect ClientID value '%s'", c.ClientID)
	}

	if len(c.ClientSecret) == 0 {
		return fmt.Errorf("incorrect ClientSecret value '%s'", c.ClientSecret)
	}

	if len(c.TokenURL) == 0 {
		return fmt.Errorf("incorrect TokenURL value '%s'", c.TokenURL)
	}

	return nil
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
	if js.conf.AccessToken != "" {
		return &oauth2.Token{
				AccessToken: js.conf.AccessToken,
				// Never expires
				// Because there is no way to update existing token
				Expiry: time.Unix(1<<63-1, 0),
			},
			nil
	}

	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	v.Set("client_id", js.conf.ClientID)
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
