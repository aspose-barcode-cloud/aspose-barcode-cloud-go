package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBarcodeGenerateBarcodeTypeGet(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIBarcodeGenerateBarcodeTypeGetOpts{}

	fileBytes, httpResponse, err := apiClient.GenerateAPI.BarcodeGenerateBarcodeTypeGet(authCtx, barcode.EncodeBarcodeTypeCode128, "Hello", opts)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")
	assert.Contains(t, httpResponse.Header.Get("Content-Disposition"), "png")
}

func TestBarcodeGenerateBodyPost(t *testing.T) {
	apiClient, authCtx := setup(t)

	// Test case for barcodeGenerateBodyPost
	imageParams := barcode.BarcodeImageParams{
		ImageFormat: barcode.AvailableBarCodeImageFormatJpeg,
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

	fileBytes, httpResponse, err := apiClient.GenerateAPI.BarcodeGenerateBodyPost(authCtx, generatorParams)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	// Check the content length and file name
	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")
	assert.Contains(t, httpResponse.Header.Get("Content-Disposition"), "jpeg")

}

func TestBarcodeGenerateFormPost(t *testing.T) {
	apiClient, authCtx := setup(t)

	// Test case for barcodeGenerateFormPost
	opts := &barcode.GenerateAPIBarcodeGenerateFormPostOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeHexBytes),
	}

	fileBytes, httpResponse, err := apiClient.GenerateAPI.BarcodeGenerateFormPost(authCtx, barcode.EncodeBarcodeTypeQR, "54657374", opts)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	// Check the content length and file name
	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")
	assert.Contains(t, httpResponse.Header.Get("Content-Disposition"), "png")
}
