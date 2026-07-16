package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

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

func TestGenerateBodyWithQrGroupedParamsOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	params := barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeQR,
		EncodeData: barcode.EncodeData{
			DataType: barcode.EncodeDataTypeStringData,
			Data:     generateParamsTestData,
		},
		BarcodeImageParams: barcode.BarcodeImageParams{
			ImageFormat:     barcode.BarcodeImageFormatPng,
			TextLocation:    barcode.CodeLocationNone,
			ForegroundColor: "#FF000000",
			BackgroundColor: "#FFFFFFFF",
			Units:           barcode.GraphicsUnitPixel,
			Resolution:      150,
			ImageHeight:     240,
			ImageWidth:      240,
			RotationAngle:   90,
		},
		QrParams: barcode.QrParams{
			QrEncodeMode:  barcode.QREncodeModeECI,
			QrErrorLevel:  barcode.QRErrorLevelLevelM,
			QrVersion:     barcode.QRVersionVersion04,
			QrECIEncoding: barcode.ECIEncodingsUTF8,
			QrAspectRatio: 0.75,
		},
	}

	fileBytes, response, err := apiClient.GenerateAPI.GenerateBody(authCtx, params)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}

func TestGenerateBodyWithMicroQrVersionOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	params := barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeMicroQR,
		EncodeData: barcode.EncodeData{
			DataType: barcode.EncodeDataTypeStringData,
			Data:     "ABC123",
		},
		BarcodeImageParams: barcode.BarcodeImageParams{
			ImageFormat: barcode.BarcodeImageFormatPng,
			ImageHeight: 160,
			ImageWidth:  160,
		},
		QrParams: barcode.QrParams{
			MicroQRVersion: barcode.MicroQRVersionM4,
		},
	}

	fileBytes, response, err := apiClient.GenerateAPI.GenerateBody(authCtx, params)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}

func TestGenerateBodyWithRectMicroQrVersionOnline(t *testing.T) {
	apiClient, authCtx := setup(t)

	params := barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeRectMicroQR,
		EncodeData: barcode.EncodeData{
			DataType: barcode.EncodeDataTypeStringData,
			Data:     "ABC123",
		},
		BarcodeImageParams: barcode.BarcodeImageParams{
			ImageFormat: barcode.BarcodeImageFormatPng,
			ImageHeight: 160,
			ImageWidth:  320,
		},
		QrParams: barcode.QrParams{
			RectMicroQrVersion: barcode.RectMicroQRVersionR13x59,
		},
	}

	fileBytes, response, err := apiClient.GenerateAPI.GenerateBody(authCtx, params)

	requireNoGenerateAPIError(t, err)
	require.NotNil(t, response)
	assertGeneratedContent(t, fileBytes)
}

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

func requireNoGenerateAPIError(t *testing.T, err error) {
	t.Helper()
	if apiError, ok := err.(barcode.GenericAPIError); ok {
		require.NoErrorf(t, err, "status=%d text=%s model=%#v", apiError.StatusCode, apiError.Text(), apiError.Model())
	}
	require.NoError(t, err)
}
