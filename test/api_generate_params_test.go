package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

// Live-API coverage for grouped optional parameters on the Generate endpoints.
// This file holds the query-parameter Generate case (Code128Params) plus the
// shared generateParamsTestData constant and requireNoGenerateAPIError helper
// used across the grouped-params suite. The symbology-specific bodies live in
// api_generate_params_qr_test.go (GenerateBody / QrParams) and
// api_generate_params_pdf417_test.go (GenerateMultipart / Pdf417Params).

const generateParamsTestData = "Aspose.BarCode.Cloud"

func TestGenerateWithGroupedOptionalParamsOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIGenerateOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeStringData),
		BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
			ImageFormat:     barcode.BarcodeImageFormatPng,
			TextLocation:    barcode.CodeLocationAbove,
			ForegroundColor: "#FF000000",
			BackgroundColor: "#FFFFFFFF",
			Units:           barcode.GraphicsUnitPixel,
			Resolution:      150,
			ImageHeight:     360,
			ImageWidth:      640,
			RotationAngle:   90,
		}),
		Code128Params: optional.NewInterface(barcode.Code128Params{
			Code128EncodeMode: barcode.Code128EncodeModeCodeB,
		}),
	}

	fileBytes, response, err := apiClient.GenerateAPI.Generate(
		authCtx,
		barcode.EncodeBarcodeTypeCode128,
		generateParamsTestData,
		opts,
	)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}

func requireNoGenerateAPIError(t *testing.T, err error) {
	t.Helper()
	if apiError, ok := err.(barcode.GenericAPIError); ok {
		require.NoErrorf(t, err, "status=%d text=%s model=%#v", apiError.StatusCode, apiError.Text(), apiError.Model())
	}
	require.NoError(t, err)
}
