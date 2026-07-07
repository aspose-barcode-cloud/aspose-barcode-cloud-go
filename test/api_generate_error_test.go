package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
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
