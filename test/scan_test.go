package test

import (
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanCode39(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	fileName := "../testdata/Code39.jpg"

	imageFile, err := os.Open(fileName)
	require.Nil(t, err)

	optionals := barcode.BarcodeApiScanBarcodeOpts{
		DecodeTypes: optional.NewInterface([]barcode.DecodeBarcodeType{
			barcode.DecodeBarcodeTypeCode39Extended,
		}),
		ChecksumValidation: optional.NewString(string(barcode.ChecksumValidationOff)),
	}
	recognized, _, err := client.BarcodeApi.ScanBarcode(authCtx, imageFile, &optionals)
	require.Nil(t, err)
	require.Equal(t, 1, len(recognized.Barcodes))

	assert.Equal(t, string(barcode.DecodeBarcodeTypeCode39Extended), recognized.Barcodes[0].Type)
	assert.Equal(t, "8M93", recognized.Barcodes[0].BarcodeValue)
	require.Greater(t, len(recognized.Barcodes[0].Region), 0)
	assert.Greater(t, recognized.Barcodes[0].Region[0].X, int32(0))
	assert.Greater(t, recognized.Barcodes[0].Region[0].Y, int32(0))
}
