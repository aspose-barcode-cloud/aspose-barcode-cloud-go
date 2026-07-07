package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

func TestWrongRecognizeType(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.RecognizeAPI.Recognize(
		authCtx,
		barcode.DecodeBarcodeType("wrong"),
		"https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png",
		nil,
	)

	requireAPIError(t, err)
}

func TestWrongRecognizeBase64Content(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.RecognizeAPI.RecognizeBase64(authCtx, barcode.RecognizeBase64Request{
		BarcodeTypes: []barcode.DecodeBarcodeType{barcode.DecodeBarcodeTypeQR},
		FileBase64:   "not-base64",
	})

	requireAPIError(t, err)
}

func TestWrongRecognizeMultipartType(t *testing.T) {
	client, authCtx := setup(t)

	filePath := filepath.Join(GetTestDataFolder(), "qr.png")
	file, err := os.Open(filePath)
	require.NoError(t, err)
	defer file.Close()

	_, _, err = client.RecognizeAPI.RecognizeMultipart(authCtx, barcode.DecodeBarcodeType("wrong"), file, nil)

	requireAPIError(t, err)
}
