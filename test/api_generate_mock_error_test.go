package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateBodyErrorResponses exercises GenerateAPI.GenerateBody's error-response
// handling for a 4xx body that can't be decoded and for a generic 5xx, using a mock
// HTTP server (no live API). The happy-path param-forwarding mock tests live in
// api_generate_mock_test.go (same package).
func TestGenerateBodyErrorResponses(t *testing.T) {
	newClient := func(baseURL string) *barcode.APIClient {
		cfg := barcode.NewConfiguration()
		cfg.BasePath = baseURL
		return barcode.NewAPIClient(cfg)
	}
	params := barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeQR,
		EncodeData:  barcode.EncodeData{Data: "test"},
	}

	t.Run("4xx with undecodable body", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{ not valid json`))
		}))
		defer srv.Close()

		_, resp, err := newClient(srv.URL).GenerateAPI.GenerateBody(context.Background(), params)
		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("5xx", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("boom"))
		}))
		defer srv.Close()

		_, resp, err := newClient(srv.URL).GenerateAPI.GenerateBody(context.Background(), params)
		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var apiErr barcode.GenericAPIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, http.StatusInternalServerError, apiErr.StatusCode)
	})
}
