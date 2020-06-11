package test

import (
	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	models "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWrongFormat(t *testing.T) {
	_, _, err := NewClientForTests().BarcodeApi.GetBarcodeGenerate(
		NewAuthContextForTests(),
		string(models.EncodeBarcodeTypeCode128),
		"text",
		&barcode.BarcodeApiGetBarcodeGenerateOpts{
			Format: optional.NewString("wrong"),
		},
	)
	require.NotNil(t, err)
	apiError := err.(barcode.GenericApiError)
	model := apiError.Model().(models.BarCodeErrorResponse)
	assert.Equal(t, "Format 'wrong' is not supported. Available formats are: jpeg, jpg, png, gif, bmp, tiff, svg", model.Error.Message)
}
