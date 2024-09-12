package test

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecognize(t *testing.T) {
	apiClient, authCtx := setup(t)

	fileUrl := "https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png"

	opts := &barcode.RecognizeAPIRecognizeOpts{}

	response, _, err := apiClient.RecognizeAPI.Recognize(authCtx, barcode.DecodeBarcodeTypeQR, fileUrl, opts)
	require.Nil(t, err)
	require.NotNil(t, response)

	assert.Equal(t, 1, len(response.Barcodes))
	barcode := response.Barcodes[0]
	assert.Equal(t, "QR", barcode.Type)
	assert.Equal(t, "http://en.m.wikipedia.org", barcode.BarcodeValue)
}

func TestRecognizeBase64(t *testing.T) {
	apiClient, authCtx := setup(t)

	filePath := filepath.Join(GetTestDataFolder(), "pdf417Sample.png")
	fileContent, err := ioutil.ReadFile(filePath)
	require.Nil(t, err)

	encodedString := base64.StdEncoding.EncodeToString(fileContent)
	request := barcode.RecognizeBase64Request{
		BarcodeTypes:         []barcode.DecodeBarcodeType{barcode.DecodeBarcodeTypePdf417},
		FileBase64:           encodedString,
		RecognitionImageKind: barcode.RecognitionImageKindClearImage,
		RecognitionMode:      barcode.RecognitionModeFast,
	}

	response, _, err := apiClient.RecognizeAPI.RecognizeBase64(authCtx, request)
	require.Nil(t, err)
	require.NotNil(t, response)

	assert.Equal(t, 1, len(response.Barcodes))
	barcode := response.Barcodes[0]
	assert.Equal(t, "Pdf417", barcode.Type)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", barcode.BarcodeValue)
	require.NotNil(t, barcode.Region)
	assert.True(t, barcode.Region[0].X > 0)
	assert.True(t, barcode.Region[0].Y > 0)
}

func TestRecognizeMultipart(t *testing.T) {
	apiClient, authCtx := setup(t)

	filePath := filepath.Join(GetTestDataFolder(), "ManyTypes.png")
	file, err := os.Open(filePath)
	require.Nil(t, err)
	defer file.Close()

	opts := &barcode.RecognizeAPIRecognizeMultipartOpts{
		RecognitionImageKind: optional.NewInterface(barcode.RecognitionImageKindClearImage),
		RecognitionMode:      optional.NewInterface(barcode.RecognitionModeNormal),
	}

	response, _, err := apiClient.RecognizeAPI.RecognizeMultipart(authCtx, barcode.DecodeBarcodeTypeMostCommonlyUsed, file, opts)
	require.Nil(t, err)
	require.NotNil(t, response)

	assert.Equal(t, 3, len(response.Barcodes))
}
