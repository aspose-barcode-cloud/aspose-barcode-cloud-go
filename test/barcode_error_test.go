package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrongFormat(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	_, _, err = client.BarcodeApi.GetBarcodeGenerate(
		authCtx,
		string(barcode.EncodeBarcodeTypeCode128),
		"text",
		&barcode.BarcodeApiGetBarcodeGenerateOpts{
			Format: optional.NewString("wrong"),
		},
	)
	require.NotNil(t, err)
	apiError := err.(barcode.GenericAPIError)
	model := apiError.Model().(barcode.ApiErrorResponse)
	assert.Equal(t, "Format 'wrong' is not supported. Available formats are: jpeg, jpg, png, gif, bmp, tiff, svg", model.Error.Message)
}
