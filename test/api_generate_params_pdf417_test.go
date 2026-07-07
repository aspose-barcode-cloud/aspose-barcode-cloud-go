package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

// Live-API coverage for the Pdf417Params grouped-parameter model. Each test
// drives GenerateAPI.GenerateMultipart with a populated Pdf417Params block so
// the PDF417 and MicroPDF417 branches (error level, truncation, ECI, macro
// characters, linked mode, Code128 emulation) are exercised end to end.
//
// The shared generateParamsTestData constant and the requireNoGenerateAPIError
// helper live alongside the query-Generate case in api_generate_params_test.go
// (same package).

func TestGenerateMultipartWithPdf417GroupedParamsOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIGenerateMultipartOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeStringData),
		BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
			ImageFormat:   barcode.BarcodeImageFormatPng,
			TextLocation:  barcode.CodeLocationAbove,
			Units:         barcode.GraphicsUnitPixel,
			Resolution:    150,
			ImageHeight:   240,
			ImageWidth:    360,
			RotationAngle: 180,
		}),
		Pdf417Params: optional.NewInterface(barcode.Pdf417Params{
			Pdf417EncodeMode:      barcode.Pdf417EncodeModeECI,
			Pdf417ErrorLevel:      barcode.Pdf417ErrorLevelLevel2,
			Pdf417Truncate:        true,
			Pdf417Columns:         5,
			Pdf417Rows:            12,
			Pdf417AspectRatio:     3,
			Pdf417ECIEncoding:     barcode.ECIEncodingsUTF8,
			Pdf417MacroCharacters: barcode.MacroCharacterMacro05,
		}),
	}

	fileBytes, response, err := apiClient.GenerateAPI.GenerateMultipart(
		authCtx,
		barcode.EncodeBarcodeTypePdf417,
		generateParamsTestData,
		opts,
	)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}

func TestGenerateMultipartWithMicroPdf417LinkedParamsOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIGenerateMultipartOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeStringData),
		BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
			ImageFormat: barcode.BarcodeImageFormatPng,
			ImageHeight: 160,
			ImageWidth:  240,
		}),
		Pdf417Params: optional.NewInterface(barcode.Pdf417Params{
			Pdf417IsLinked: true,
		}),
	}

	fileBytes, response, err := apiClient.GenerateAPI.GenerateMultipart(
		authCtx,
		barcode.EncodeBarcodeTypeMicroPdf417,
		"1234567890",
		opts,
	)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}

func TestGenerateMultipartWithMicroPdf417Code128EmulationOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIGenerateMultipartOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeStringData),
		BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
			ImageFormat: barcode.BarcodeImageFormatPng,
			ImageHeight: 160,
			ImageWidth:  240,
		}),
		Pdf417Params: optional.NewInterface(barcode.Pdf417Params{
			Pdf417IsCode128Emulation: true,
		}),
	}

	fileBytes, response, err := apiClient.GenerateAPI.GenerateMultipart(
		authCtx,
		barcode.EncodeBarcodeTypeMicroPdf417,
		"1234567890",
		opts,
	)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}
