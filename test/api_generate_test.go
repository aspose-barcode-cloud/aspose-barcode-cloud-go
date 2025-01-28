package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIGenerateOpts{
		ImageFormat: optional.NewInterface(barcode.BarcodeImageFormatSvg),
	}

	fileBytes, _, err := apiClient.GenerateAPI.Generate(authCtx, barcode.EncodeBarcodeTypeCode128, "Hello", opts)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")
}

func TestGenerateBody(t *testing.T) {
	apiClient, authCtx := setup(t)

	// Test case for GenerateBody
	imageParams := barcode.BarcodeImageParams{
		ImageFormat: barcode.BarcodeImageFormatJpeg,
	}

	encodeData := barcode.EncodeData{
		Data:     "VGVzdA==",
		DataType: barcode.EncodeDataTypeBase64Bytes,
	}

	generatorParams := barcode.GenerateParams{
		BarcodeType:        barcode.EncodeBarcodeTypeQR,
		EncodeData:         encodeData,
		BarcodeImageParams: imageParams,
	}

	fileBytes, _, err := apiClient.GenerateAPI.GenerateBody(authCtx, generatorParams)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	// Check the content length and file name
	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")

}

func TestGenerateMultipart(t *testing.T) {
	apiClient, authCtx := setup(t)

	// Test case for GenerateMultipart
	opts := &barcode.GenerateAPIGenerateMultipartOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeHexBytes),
	}

	fileBytes, _, err := apiClient.GenerateAPI.GenerateMultipart(authCtx, barcode.EncodeBarcodeTypeQR, "54657374", opts)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	// Check the content length and file name
	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")

}
