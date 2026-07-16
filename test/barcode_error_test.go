package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrongFormat(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.GenerateAPI.Generate(
		authCtx,
		barcode.EncodeBarcodeTypeCode128,
		"text",
		&barcode.GenerateAPIGenerateOpts{
			BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
				ImageFormat: barcode.BarcodeImageFormat("wrong"),
			}),
		},
	)
	apiError := requireAPIError(t, err)
	model := apiError.Model().(barcode.ApiErrorResponse)
	assert.Equal(t, "Error: Field name: 'ImageFormat' errors: The value 'wrong' is not valid for ImageFormat.", model.Error.Message)
}

func TestWrongFormatGenerateBody(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.GenerateAPI.GenerateBody(authCtx, barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeCode128,
		EncodeData: barcode.EncodeData{
			DataType: barcode.EncodeDataTypeStringData,
			Data:     "text",
		},
		BarcodeImageParams: barcode.BarcodeImageParams{
			ImageFormat: barcode.BarcodeImageFormat("wrong"),
		},
	})

	requireAPIError(t, err)
}

func TestWrongFormatGenerateMultipart(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.GenerateAPI.GenerateMultipart(
		authCtx,
		barcode.EncodeBarcodeTypeCode128,
		"text",
		&barcode.GenerateAPIGenerateMultipartOpts{
			BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
				ImageFormat: barcode.BarcodeImageFormat("wrong"),
			}),
		},
	)

	requireAPIError(t, err)
}

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

func TestWrongScanUrl(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.ScanAPI.Scan(authCtx, "not-a-url")

	requireAPIError(t, err)
}

func TestWrongScanBase64Content(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.ScanAPI.ScanBase64(authCtx, barcode.ScanBase64Request{
		FileBase64: "not-base64",
	})

	requireAPIError(t, err)
}

func TestWrongScanMultipartContent(t *testing.T) {
	client, authCtx := setup(t)

	file, err := os.CreateTemp(t.TempDir(), "not-image-*.txt")
	require.NoError(t, err)
	_, err = file.WriteString("not an image")
	require.NoError(t, err)
	_, err = file.Seek(0, 0)
	require.NoError(t, err)
	defer file.Close()

	_, _, err = client.ScanAPI.ScanMultipart(authCtx, file)

	requireAPIError(t, err)
}

func requireAPIError(t *testing.T, err error) barcode.GenericAPIError {
	t.Helper()
	require.Error(t, err)

	apiError, ok := err.(barcode.GenericAPIError)
	require.Truef(t, ok, "expected GenericAPIError, got %T: %v", err, err)
	assert.GreaterOrEqual(t, apiError.StatusCode, 400)
	assert.NotEmpty(t, apiError.Text())

	return apiError
}
