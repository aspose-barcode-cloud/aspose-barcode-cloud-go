package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock-server (httptest, no live API) coverage that the optional Generate
// parameters are actually serialized onto the wire. Error-response handling for
// the same endpoints lives in api_generate_mock_error_test.go (same package).

// TestGenerateForwardsAllOptionalParams drives GenerateAPI.Generate against a mock
// HTTP server (no live API) with fully-populated QrParams and Pdf417Params so every
// optional query-parameter branch is exercised.
func TestGenerateForwardsAllOptionalParams(t *testing.T) {
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("\x89PNG\r\n\x1a\n"))
	}))
	defer srv.Close()

	cfg := barcode.NewConfiguration()
	cfg.BasePath = srv.URL
	client := barcode.NewAPIClient(cfg)

	opts := &barcode.GenerateAPIGenerateOpts{
		QrParams: optional.NewInterface(barcode.QrParams{
			QrEncodeMode:       barcode.QREncodeModeAuto,
			QrErrorLevel:       barcode.QRErrorLevelLevelL,
			QrVersion:          barcode.QRVersionAuto,
			QrECIEncoding:      barcode.ECIEncodingsNONE,
			QrAspectRatio:      0.75,
			MicroQRVersion:     barcode.MicroQRVersionAuto,
			RectMicroQrVersion: barcode.RectMicroQRVersionAuto,
		}),
		Pdf417Params: optional.NewInterface(barcode.Pdf417Params{
			Pdf417EncodeMode:             barcode.Pdf417EncodeModeAuto,
			Pdf417ErrorLevel:             barcode.Pdf417ErrorLevelLevel0,
			Pdf417Truncate:               true,
			Pdf417Columns:                5,
			Pdf417Rows:                   6,
			Pdf417AspectRatio:            3.0,
			Pdf417ECIEncoding:            barcode.ECIEncodingsNONE,
			Pdf417IsReaderInitialization: true,
			Pdf417MacroCharacters:        barcode.MacroCharacterNone,
			Pdf417IsLinked:               true,
			Pdf417IsCode128Emulation:     true,
		}),
	}

	data, resp, err := client.GenerateAPI.Generate(context.Background(), barcode.EncodeBarcodeTypeQR, "test", opts)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, data)

	want := []string{
		"qrEncodeMode", "qrErrorLevel", "qrVersion", "qrECIEncoding", "qrAspectRatio",
		"microQRVersion", "rectMicroQrVersion",
		"pdf417EncodeMode", "pdf417ErrorLevel", "pdf417Truncate", "pdf417Columns", "pdf417Rows",
		"pdf417AspectRatio", "pdf417ECIEncoding", "pdf417IsReaderInitialization",
		"pdf417MacroCharacters", "pdf417IsLinked", "pdf417IsCode128Emulation",
	}
	for _, p := range want {
		assert.Truef(t, captured.Has(p), "expected query param %q to be forwarded", p)
	}
	assert.Equal(t, "Auto", captured.Get("qrEncodeMode"))
	assert.Equal(t, "5", captured.Get("pdf417Columns"))
}

// TestGenerateMultipartForwardsAllOptionalParams drives GenerateAPI.GenerateMultipart
// against a mock server so the optional multipart form-field branches (foreground/
// background color, QR and PDF417 params) are exercised.
func TestGenerateMultipartForwardsAllOptionalParams(t *testing.T) {
	var form url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseMultipartForm(1 << 20)
		if r.MultipartForm != nil {
			form = url.Values(r.MultipartForm.Value)
		}
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("\x89PNG\r\n\x1a\n"))
	}))
	defer srv.Close()

	cfg := barcode.NewConfiguration()
	cfg.BasePath = srv.URL
	client := barcode.NewAPIClient(cfg)

	opts := &barcode.GenerateAPIGenerateMultipartOpts{
		BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
			ForegroundColor: "Black",
			BackgroundColor: "White",
		}),
		QrParams: optional.NewInterface(barcode.QrParams{
			QrEncodeMode:  barcode.QREncodeModeAuto,
			QrErrorLevel:  barcode.QRErrorLevelLevelL,
			QrVersion:     barcode.QRVersionAuto,
			QrAspectRatio: 0.75,
		}),
		Pdf417Params: optional.NewInterface(barcode.Pdf417Params{
			Pdf417EncodeMode: barcode.Pdf417EncodeModeAuto,
			Pdf417Columns:    5,
		}),
	}

	data, resp, err := client.GenerateAPI.GenerateMultipart(context.Background(), barcode.EncodeBarcodeTypeQR, "test", opts)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, data)

	for _, p := range []string{
		"foregroundColor", "backgroundColor",
		"qrEncodeMode", "qrErrorLevel", "qrVersion", "qrAspectRatio",
		"pdf417EncodeMode", "pdf417Columns",
	} {
		assert.Truef(t, form.Has(p), "expected multipart form field %q", p)
	}
}
