package test

import (
	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWrongFormat(t *testing.T) {
	_, _, err := NewClientForTests().BarcodeApi.GetBarcodeGenerate(
		NewAuthContextForTests(),
		string(barcode.EncodeBarcodeTypeCode128),
		"text",
		&barcode.BarcodeApiGetBarcodeGenerateOpts{
			Format: optional.NewString("wrong"),
		},
	)
	require.NotNil(t, err)
	apiError := err.(barcode.GenericAPIError)
	model := apiError.Model().(barcode.BarCodeErrorResponse)
	assert.Equal(t, "Format 'wrong' is not supported. Available formats are: jpeg, jpg, png, gif, bmp, tiff, svg", model.Error.Message)
}
