package test

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBarcodeScanFormPost(t *testing.T) {
	apiClient, authCtx := setup(t)

	filePath := filepath.Join(GetTestDataFolder(), "pdf417Sample.png")
	file, err := os.Open(filePath)
	require.Nil(t, err)
	defer file.Close()

	scanResponse, _, err := apiClient.ScanAPI.BarcodeScanFormPost(authCtx, file)
	require.Nil(t, err)
	require.NotNil(t, scanResponse)

	assert.Equal(t, 1, len(scanResponse.Barcodes))
	assert.Equal(t, string(barcode.DecodeBarcodeTypePdf417), scanResponse.Barcodes[0].Type)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", scanResponse.Barcodes[0].BarcodeValue)
}

func TestBarcodeScanBodyPost(t *testing.T) {
	apiClient, authCtx := setup(t)

	filePath := filepath.Join(GetTestDataFolder(), "QR_and_Code128.png")
	fileContent, err := ioutil.ReadFile(filePath)
	require.Nil(t, err)

	encodedFile := base64.StdEncoding.EncodeToString(fileContent)

	scanBase64Request := barcode.ScanBase64Request{
		FileBase64: encodedFile,
	}

	scanResponse, _, err := apiClient.ScanAPI.BarcodeScanBodyPost(authCtx, scanBase64Request)
	require.Nil(t, err)
	require.NotNil(t, scanResponse)

	assert.Equal(t, 2, len(scanResponse.Barcodes))
	assert.Equal(t, string(barcode.DecodeBarcodeTypeQR), scanResponse.Barcodes[0].Type)
	assert.Equal(t, "QR text", scanResponse.Barcodes[0].BarcodeValue)
	assert.Equal(t, string(barcode.DecodeBarcodeTypeCode128), scanResponse.Barcodes[1].Type)
	assert.Equal(t, "Code128 text", scanResponse.Barcodes[1].BarcodeValue)
}

func TestBarcodeScanGet(t *testing.T) {
	apiClient, authCtx := setup(t)

	urlStr := "https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png"

	scanResponse, _, err := apiClient.ScanAPI.BarcodeScanGet(authCtx, urlStr)
	require.Nil(t, err)
	require.NotNil(t, scanResponse)

	assert.Equal(t, 1, len(scanResponse.Barcodes))
	assert.Equal(t, "QR", scanResponse.Barcodes[0].Type)
	assert.Equal(t, "http://en.m.wikipedia.org", scanResponse.Barcodes[0].BarcodeValue)
}