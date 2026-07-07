package test

import (
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

// Live-API coverage for the QrParams grouped-parameter model. Each test drives
// GenerateAPI.GenerateBody with a populated QrParams block so the QR, MicroQR
// and RectMicroQR version/encoding branches are exercised end to end.
//
// The shared generateParamsTestData constant and the requireNoGenerateAPIError
// helper live alongside the query-Generate case in api_generate_params_test.go
// (same package).

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
