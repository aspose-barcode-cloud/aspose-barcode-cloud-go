package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

func TestGenerateBatchBinaryResponse(t *testing.T) {
	for _, contentType := range []string{"application/zip", "application/x-tar", "application/gzip"} {
		t.Run(contentType, func(t *testing.T) {
			expected := []byte{0, 1, 0xff, 0x80}
			var captured barcode.BatchGenerateRequest
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodPost, r.Method)
				require.Equal(t, "/barcode/generate-batch", r.URL.Path)
				require.NoError(t, json.NewDecoder(r.Body).Decode(&captured))
				w.Header().Set("Content-Type", contentType)
				_, _ = w.Write(expected)
			}))
			defer server.Close()
			cfg := barcode.NewConfiguration()
			cfg.BasePath = server.URL
			client := barcode.NewAPIClient(cfg)
			batch := barcode.BatchGenerateRequest{ArchiveFormat: barcode.BatchArchiveFormatZip,
				Items: []barcode.BatchGenerateItem{{FileName: "qr.png", Generation: barcode.GenerateParams{
					BarcodeType: barcode.EncodeBarcodeTypeQR, EncodeData: barcode.EncodeData{Data: "Batch example"},
				}}},
			}
			actual, response, err := client.GenerateAPI.GenerateBatch(context.Background(), batch)
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, response.StatusCode)
			require.Equal(t, expected, actual)
			require.Equal(t, batch, captured)
		})
	}
}

func TestGenerateBatchErrors(t *testing.T) {
	for _, scenario := range []struct {
		name   string
		status int
		body   string
	}{
		{"validation", 400, `{"error":{"code":"invalidRequest","message":"Invalid batch"}}`},
		{"invalid-error-json", 401, "not-json"},
		{"server-error", 500, "unavailable"},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(scenario.status)
				_, _ = w.Write([]byte(scenario.body))
			}))
			defer server.Close()
			cfg := barcode.NewConfiguration()
			cfg.BasePath = server.URL
			client := barcode.NewAPIClient(cfg)
			_, response, err := client.GenerateAPI.GenerateBatch(context.Background(), barcode.BatchGenerateRequest{})
			require.Error(t, err)
			require.Equal(t, scenario.status, response.StatusCode)
		})
	}
	t.Run("invalid-url", func(t *testing.T) {
		cfg := barcode.NewConfiguration()
		cfg.BasePath = "://invalid"
		_, _, err := barcode.NewAPIClient(cfg).GenerateAPI.GenerateBatch(context.Background(), barcode.BatchGenerateRequest{})
		require.Error(t, err)
	})
	t.Run("canceled-context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, _, err := barcode.NewAPIClient(barcode.NewConfiguration()).GenerateAPI.GenerateBatch(ctx, barcode.BatchGenerateRequest{})
		require.Error(t, err)
	})
}
