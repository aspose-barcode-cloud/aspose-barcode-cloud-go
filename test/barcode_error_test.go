package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrongFormat(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	_, _, err = client.GenerateAPI.Generate(
		authCtx,
		barcode.EncodeBarcodeTypeCode128,
		"text",
		&barcode.GenerateAPIGenerateOpts{
			ImageFormat: optional.NewInterface("wrong"),
		},
	)
	require.NotNil(t, err)
	apiError := err.(barcode.GenericAPIError)
	model := apiError.Model().(barcode.ApiErrorResponse)
	assert.Equal(t, "Error: Field name: 'ImageFormat' errors: The value 'wrong' is not valid for ImageFormat.", model.Error.Message)
}
