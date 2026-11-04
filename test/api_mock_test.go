package test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

// The tests in this file run every API operation against a local httptest server, so
// they need no credentials and no network. They keep the SDK's line coverage above the
// CI gate on their own; the live tests in the other files check the real service.

const mockBarcodeListJSON = `{"barcodes":[{"barcodeValue":"Aspose","type":"QR","checksum":"","region":[{"x":1,"y":2}]}]}`

// The models carry JSON tags only, so encoding/xml matches elements by Go field name.
const mockBarcodeListXML = `<BarcodeResponseList><Barcodes><BarcodeValue>Aspose</BarcodeValue><Type>QR</Type></Barcodes></BarcodeResponseList>`

const mockAPIErrorJSON = `{"requestId":"req-1","error":{"code":"BadRequest","message":"Invalid barcode type"}}`

// capturedRequest is what the mock server saw for the last request.
type capturedRequest struct {
	baseURL string
	host    string
	method  string
	path    string
	query   url.Values
	header  http.Header
	body    []byte
	form    url.Values
	files   map[string]string
}

// newMockServer starts a server that records the request and answers with the given
// status, content type and body.
func newMockServer(t *testing.T, status int, contentType string, body string) (*barcode.APIClient, *capturedRequest) {
	t.Helper()
	captured := &capturedRequest{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured.host = r.Host
		captured.method = r.Method
		captured.path = r.URL.Path
		captured.query = r.URL.Query()
		captured.header = r.Header.Clone()
		if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/") {
			require.NoError(t, r.ParseMultipartForm(1<<20))
			captured.form = url.Values(r.MultipartForm.Value)
			captured.files = map[string]string{}
			for field, headers := range r.MultipartForm.File {
				captured.files[field] = headers[0].Filename
			}
		} else {
			captured.body, _ = io.ReadAll(r.Body)
		}
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	captured.baseURL = srv.URL

	cfg := barcode.NewConfiguration()
	cfg.BasePath = srv.URL
	return barcode.NewAPIClient(cfg), captured
}

func openTestImage(t *testing.T) *os.File {
	t.Helper()
	file, err := os.Open(filepath.Join(GetTestDataFolder(), "Code128.png"))
	require.NoError(t, err)
	t.Cleanup(func() { _ = file.Close() })
	return file
}

func requireMockBarcodes(t *testing.T, result barcode.BarcodeResponseList) {
	t.Helper()
	require.Len(t, result.Barcodes, 1)
	assert.Equal(t, "Aspose", result.Barcodes[0].BarcodeValue)
	assert.Equal(t, "QR", result.Barcodes[0].Type)
}

func TestRecognizeMock(t *testing.T) {
	client, captured := newMockServer(t, http.StatusOK, "application/json; charset=utf-8", mockBarcodeListJSON)

	result, resp, err := client.RecognizeAPI.Recognize(context.Background(), barcode.DecodeBarcodeTypeQR,
		"https://example.com/qr.png", &barcode.RecognizeAPIRecognizeOpts{
			RecognitionMode:      optional.NewInterface(barcode.RecognitionModeFast),
			RecognitionImageKind: optional.NewInterface(barcode.RecognitionImageKindPhoto),
		})
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	requireMockBarcodes(t, result)

	assert.Equal(t, http.MethodGet, captured.method)
	assert.Equal(t, "/barcode/recognize", captured.path)
	assert.Equal(t, "QR", captured.query.Get("barcodeType"))
	assert.Equal(t, "https://example.com/qr.png", captured.query.Get("fileUrl"))
	assert.Equal(t, "Fast", captured.query.Get("recognitionMode"))
	assert.Equal(t, "Photo", captured.query.Get("recognitionImageKind"))
	assert.Equal(t, "application/json", captured.header.Get("Accept"))
	assert.Equal(t, "go sdk", captured.header.Get("x-aspose-client"))
	assert.NotEmpty(t, captured.header.Get("x-aspose-client-version"))
}

func TestRecognizeBase64Mock(t *testing.T) {
	client, captured := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)

	request := barcode.RecognizeBase64Request{
		BarcodeTypes:         []barcode.DecodeBarcodeType{barcode.DecodeBarcodeTypeQR, barcode.DecodeBarcodeTypeCode128},
		FileBase64:           base64.StdEncoding.EncodeToString([]byte("image")),
		RecognitionMode:      barcode.RecognitionModeFast,
		RecognitionImageKind: barcode.RecognitionImageKindClearImage,
	}
	result, _, err := client.RecognizeAPI.RecognizeBase64(context.Background(), request)
	require.NoError(t, err)
	requireMockBarcodes(t, result)

	assert.Equal(t, http.MethodPost, captured.method)
	assert.Equal(t, "/barcode/recognize-body", captured.path)
	assert.Equal(t, "application/json", captured.header.Get("Content-Type"))
	var sent barcode.RecognizeBase64Request
	require.NoError(t, json.Unmarshal(captured.body, &sent))
	assert.Equal(t, request, sent)
}

func TestRecognizeMultipartMock(t *testing.T) {
	// An XML answer also exercises the XML branch of the response decoder.
	client, captured := newMockServer(t, http.StatusOK, "application/xml", mockBarcodeListXML)

	result, _, err := client.RecognizeAPI.RecognizeMultipart(context.Background(), barcode.DecodeBarcodeTypeCode128,
		openTestImage(t), &barcode.RecognizeAPIRecognizeMultipartOpts{
			RecognitionMode:      optional.NewInterface(barcode.RecognitionModeFast),
			RecognitionImageKind: optional.NewInterface(barcode.RecognitionImageKindScannedDocument),
		})
	require.NoError(t, err)
	requireMockBarcodes(t, result)

	assert.Equal(t, "/barcode/recognize-multipart", captured.path)
	assert.Equal(t, "Code128", captured.form.Get("barcodeType"))
	assert.Equal(t, "Fast", captured.form.Get("recognitionMode"))
	assert.Equal(t, "ScannedDocument", captured.form.Get("recognitionImageKind"))
	assert.Equal(t, "Code128.png", captured.files["file"])
}

func TestScanMock(t *testing.T) {
	client, captured := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)

	result, _, err := client.ScanAPI.Scan(context.Background(), "https://example.com/qr.png")
	require.NoError(t, err)
	requireMockBarcodes(t, result)

	assert.Equal(t, http.MethodGet, captured.method)
	assert.Equal(t, "/barcode/scan", captured.path)
	assert.Equal(t, "https://example.com/qr.png", captured.query.Get("fileUrl"))
}

func TestScanBase64Mock(t *testing.T) {
	client, captured := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)

	request := barcode.ScanBase64Request{FileBase64: base64.StdEncoding.EncodeToString([]byte("image"))}
	result, _, err := client.ScanAPI.ScanBase64(context.Background(), request)
	require.NoError(t, err)
	requireMockBarcodes(t, result)

	assert.Equal(t, http.MethodPost, captured.method)
	assert.Equal(t, "/barcode/scan-body", captured.path)
	var sent barcode.ScanBase64Request
	require.NoError(t, json.Unmarshal(captured.body, &sent))
	assert.Equal(t, request, sent)
}

func TestScanMultipartMock(t *testing.T) {
	client, captured := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)

	result, _, err := client.ScanAPI.ScanMultipart(context.Background(), openTestImage(t))
	require.NoError(t, err)
	requireMockBarcodes(t, result)

	assert.Equal(t, "/barcode/scan-multipart", captured.path)
	assert.Equal(t, "Code128.png", captured.files["file"])
}

func TestGenerateBodyMock(t *testing.T) {
	client, captured := newMockServer(t, http.StatusOK, "image/png", "\x89PNG\r\n\x1a\n")

	params := barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeQR,
		EncodeData:  barcode.EncodeData{Data: "Aspose", DataType: barcode.EncodeDataTypeStringData},
	}
	data, _, err := client.GenerateAPI.GenerateBody(context.Background(), params)
	require.NoError(t, err)
	assert.Equal(t, []byte("\x89PNG\r\n\x1a\n"), data)

	assert.Equal(t, "/barcode/generate-body", captured.path)
	var sent barcode.GenerateParams
	require.NoError(t, json.Unmarshal(captured.body, &sent))
	assert.Equal(t, params.EncodeData, sent.EncodeData)
}

// apiCall runs one API operation against the given client and returns its error.
type apiCall struct {
	name string
	call func(t *testing.T, client *barcode.APIClient) error
}

func allAPICalls() []apiCall {
	ctx := context.Background()
	return []apiCall{
		{"Generate", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.GenerateAPI.Generate(ctx, barcode.EncodeBarcodeTypeQR, "Aspose", nil)
			return err
		}},
		{"GenerateBody", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.GenerateAPI.GenerateBody(ctx, barcode.GenerateParams{
				BarcodeType: barcode.EncodeBarcodeTypeQR,
				EncodeData:  barcode.EncodeData{Data: "Aspose"},
			})
			return err
		}},
		{"GenerateMultipart", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.GenerateAPI.GenerateMultipart(ctx, barcode.EncodeBarcodeTypeQR, "Aspose", nil)
			return err
		}},
		{"Recognize", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.RecognizeAPI.Recognize(ctx, barcode.DecodeBarcodeTypeQR, "https://example.com/qr.png", nil)
			return err
		}},
		{"RecognizeBase64", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.RecognizeAPI.RecognizeBase64(ctx, barcode.RecognizeBase64Request{
				BarcodeTypes: []barcode.DecodeBarcodeType{barcode.DecodeBarcodeTypeQR},
				FileBase64:   "aW1hZ2U=",
			})
			return err
		}},
		{"RecognizeMultipart", func(t *testing.T, c *barcode.APIClient) error {
			_, _, err := c.RecognizeAPI.RecognizeMultipart(ctx, barcode.DecodeBarcodeTypeQR, openTestImage(t), nil)
			return err
		}},
		{"Scan", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.ScanAPI.Scan(ctx, "https://example.com/qr.png")
			return err
		}},
		{"ScanBase64", func(_ *testing.T, c *barcode.APIClient) error {
			_, _, err := c.ScanAPI.ScanBase64(ctx, barcode.ScanBase64Request{FileBase64: "aW1hZ2U="})
			return err
		}},
		{"ScanMultipart", func(t *testing.T, c *barcode.APIClient) error {
			_, _, err := c.ScanAPI.ScanMultipart(ctx, openTestImage(t))
			return err
		}},
	}
}

// TestAPIErrorResponsesMock checks how every operation reports failures: a 4xx with an
// ApiErrorResponse body, a 4xx body that cannot be decoded, a 5xx and a 2xx answer
// with an unknown content type.
func TestAPIErrorResponsesMock(t *testing.T) {
	for _, op := range allAPICalls() {
		t.Run(op.name+"/4xx with ApiErrorResponse", func(t *testing.T) {
			client, _ := newMockServer(t, http.StatusBadRequest, "application/json", mockAPIErrorJSON)

			apiErr := requireAPIError(t, op.call(t, client))
			assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
			assert.Equal(t, "400 Bad Request", apiErr.Error())
			assert.JSONEq(t, mockAPIErrorJSON, apiErr.Text())
			model, ok := apiErr.Model().(barcode.ApiErrorResponse)
			require.Truef(t, ok, "expected ApiErrorResponse model, got %T", apiErr.Model())
			assert.Equal(t, "req-1", model.RequestId)
			assert.Equal(t, "Invalid barcode type", model.Error.Message)
		})

		t.Run(op.name+"/4xx with undecodable body", func(t *testing.T) {
			client, _ := newMockServer(t, http.StatusNotFound, "application/json", "{ not json")

			apiErr := requireAPIError(t, op.call(t, client))
			assert.Equal(t, http.StatusNotFound, apiErr.StatusCode)
			assert.Nil(t, apiErr.Model())
		})

		t.Run(op.name+"/5xx", func(t *testing.T) {
			client, _ := newMockServer(t, http.StatusBadGateway, "text/plain", "upstream failed")

			err := op.call(t, client)
			var apiErr barcode.GenericAPIError
			require.ErrorAs(t, err, &apiErr)
			assert.Equal(t, http.StatusBadGateway, apiErr.StatusCode)
			assert.Equal(t, "upstream failed", apiErr.Text())
		})

		t.Run(op.name+"/2xx with unknown content type", func(t *testing.T) {
			client, _ := newMockServer(t, http.StatusOK, "text/plain", "hello")

			require.EqualError(t, op.call(t, client), "undefined response type")
		})

		t.Run(op.name+"/unreachable server", func(t *testing.T) {
			cfg := barcode.NewConfiguration()
			cfg.BasePath = "http://127.0.0.1:1"
			require.Error(t, op.call(t, barcode.NewAPIClient(cfg)))
		})
	}
}

// TestAuthenticationContextsMock checks the Authorization header produced for each
// supported authentication context value.
func TestAuthenticationContextsMock(t *testing.T) {
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, r.ParseForm())
		assert.Equal(t, "client_credentials", r.PostForm.Get("grant_type"))
		if r.PostForm.Get("client_secret") != "secret" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"invalid_client"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"jwt-token","token_type":"Bearer","expires_in":3600}`))
	}))
	defer tokenServer.Close()

	jwtConfig := jwt.NewConfig("client", "secret")
	jwtConfig.TokenURL = tokenServer.URL

	cases := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{"JWT client credentials", context.WithValue(context.Background(), barcode.ContextJWT,
			jwtConfig.TokenSource(context.Background())), "Bearer jwt-token"},
		{"JWT external access token", context.WithValue(context.Background(), barcode.ContextJWT,
			(&jwt.Config{AccessToken: "external"}).TokenSource(context.Background())), "Bearer external"},
		{"OAuth2 token source", context.WithValue(context.Background(), barcode.ContextOAuth2,
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "oauth"})), "Bearer oauth"},
		{"access token", context.WithValue(context.Background(), barcode.ContextAccessToken, "plain"), "Bearer plain"},
		{"basic auth", context.WithValue(context.Background(), barcode.ContextBasicAuth,
			barcode.BasicAuth{UserName: "user", Password: "pass"}), "Basic dXNlcjpwYXNz"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client, captured := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)

			_, _, err := client.ScanAPI.Scan(tc.ctx, "https://example.com/qr.png")
			require.NoError(t, err)
			assert.Equal(t, tc.want, captured.header.Get("Authorization"))
		})
	}

	failures := []struct {
		name   string
		config *jwt.Config
	}{
		{"rejected credentials", &jwt.Config{ClientID: "client", ClientSecret: "wrong", TokenURL: tokenServer.URL}},
		{"unreachable token server", &jwt.Config{ClientID: "client", ClientSecret: "secret", TokenURL: "http://127.0.0.1:1"}},
	}
	for _, tc := range failures {
		t.Run("JWT "+tc.name, func(t *testing.T) {
			client, _ := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)
			ctx := context.WithValue(context.Background(), barcode.ContextJWT, tc.config.TokenSource(context.Background()))

			_, _, err := client.ScanAPI.Scan(ctx, "https://example.com/qr.png")
			require.Error(t, err)
		})
	}

	t.Run("OAuth2 token source error", func(t *testing.T) {
		client, _ := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)
		failing := (&jwt.Config{ClientID: "client", ClientSecret: "wrong", TokenURL: tokenServer.URL}).
			TokenSource(context.Background())
		ctx := context.WithValue(context.Background(), barcode.ContextOAuth2, failing)

		_, _, err := client.ScanAPI.Scan(ctx, "https://example.com/qr.png")
		require.Error(t, err)
	})
}

func TestJWTTokenResponsesMock(t *testing.T) {
	serve := func(status int, body string) string {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
			_, _ = w.Write([]byte(body))
		}))
		t.Cleanup(srv.Close)
		return srv.URL
	}

	t.Run("token without expiry", func(t *testing.T) {
		config := jwt.NewConfig("client", "secret")
		config.TokenURL = serve(http.StatusOK, `{"access_token":"no-expiry","token_type":"Bearer"}`)

		token, err := config.TokenSource(context.Background()).Token()
		require.NoError(t, err)
		assert.Equal(t, "no-expiry", token.AccessToken)
		assert.True(t, token.Expiry.IsZero())
	})

	t.Run("malformed token response", func(t *testing.T) {
		config := jwt.NewConfig("client", "secret")
		config.TokenURL = serve(http.StatusOK, `not json`)

		_, err := config.TokenSource(context.Background()).Token()
		require.ErrorContains(t, err, "cannot unmarshal token")
	})

	t.Run("rejected credentials", func(t *testing.T) {
		config := jwt.NewConfig("client", "secret")
		config.TokenURL = serve(http.StatusUnauthorized, `{"error":"invalid_client"}`)

		_, err := config.TokenSource(context.Background()).Token()
		var retrieveErr *oauth2.RetrieveError
		require.ErrorAs(t, err, &retrieveErr)
		assert.Equal(t, http.StatusUnauthorized, retrieveErr.Response.StatusCode)
	})
}

func TestClientConfigurationMock(t *testing.T) {
	_, captured := newMockServer(t, http.StatusOK, "application/json", mockBarcodeListJSON)

	cfg := barcode.NewConfiguration()
	cfg.BasePath = "http://127.0.0.1:1"
	client := barcode.NewAPIClient(cfg)
	_, _, err := client.ScanAPI.Scan(context.Background(), "https://example.com/qr.png")
	require.Error(t, err, "the client must call the configured base path")

	// ChangeBasePath is how callers point an existing client at another server.
	client.ChangeBasePath(captured.baseURL)
	cfg.Host = "barcode.example.test"
	cfg.AddDefaultHeader("X-Custom", "value")
	_, _, err = client.ScanAPI.Scan(context.Background(), "https://example.com/qr.png")
	require.NoError(t, err)
	assert.Equal(t, "barcode.example.test", captured.host)
	assert.Equal(t, "value", captured.header.Get("X-Custom"))
	assert.Equal(t, cfg.UserAgent, captured.header.Get("User-Agent"))

	assert.Equal(t, "auth jwt", barcode.ContextJWT.String())
}
