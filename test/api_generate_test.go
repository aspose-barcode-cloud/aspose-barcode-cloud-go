package test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	apiClient, authCtx := setup(t)

	opts := &barcode.GenerateAPIGenerateOpts{
		DataType: optional.NewInterface(barcode.EncodeDataTypeStringData),
		BarcodeImageParams: optional.NewInterface(barcode.BarcodeImageParams{
			ImageFormat:     barcode.BarcodeImageFormatSvg,
			TextLocation:    barcode.CodeLocationBelow,
			ForegroundColor: "#FF000000",
			BackgroundColor: "#FFFFFFFF",
			Units:           barcode.GraphicsUnitPixel,
			Resolution:      150,
			ImageHeight:     120,
			ImageWidth:      320,
			RotationAngle:   90,
		}),
		Code128Params: optional.NewInterface(barcode.Code128Params{
			Code128EncodeMode: barcode.Code128EncodeModeCodeB,
		}),
	}

	fileBytes, _, err := apiClient.GenerateAPI.Generate(authCtx, barcode.EncodeBarcodeTypeCode128, "Hello", opts)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	assertGeneratedContent(t, fileBytes)
}

func TestGenerateBody(t *testing.T) {
	apiClient, authCtx := setup(t)

	// Test case for GenerateBody
	imageParams := barcode.BarcodeImageParams{
		ImageFormat:     barcode.BarcodeImageFormatJpeg,
		TextLocation:    barcode.CodeLocationNone,
		ForegroundColor: "#FF000000",
		BackgroundColor: "#FFFFFFFF",
		Units:           barcode.GraphicsUnitPixel,
		Resolution:      150,
		ImageHeight:     240,
		ImageWidth:      240,
		RotationAngle:   90,
	}

	encodeData := barcode.EncodeData{
		Data:     "VGVzdA==",
		DataType: barcode.EncodeDataTypeBase64Bytes,
	}

	generatorParams := barcode.GenerateParams{
		BarcodeType:        barcode.EncodeBarcodeTypeQR,
		EncodeData:         encodeData,
		BarcodeImageParams: imageParams,
		QrParams: barcode.QrParams{
			QrEncodeMode:  barcode.QREncodeModeAuto,
			QrErrorLevel:  barcode.QRErrorLevelLevelM,
			QrVersion:     barcode.QRVersionAuto,
			QrAspectRatio: 0.75,
		},
	}

	fileBytes, _, err := apiClient.GenerateAPI.GenerateBody(authCtx, generatorParams)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	assertGeneratedContent(t, fileBytes)
}

func TestGenerateMultipart(t *testing.T) {
	apiClient, authCtx := setup(t)

	// Test case for GenerateMultipart
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
			Pdf417EncodeMode:  barcode.Pdf417EncodeModeAuto,
			Pdf417ErrorLevel:  barcode.Pdf417ErrorLevelLevel2,
			Pdf417Truncate:    true,
			Pdf417Columns:     5,
			Pdf417Rows:        12,
			Pdf417AspectRatio: 3,
		}),
	}

	fileBytes, _, err := apiClient.GenerateAPI.GenerateMultipart(authCtx, barcode.EncodeBarcodeTypePdf417, "Aspose.BarCode.Cloud", opts)
	require.Nil(t, err)
	require.NotNil(t, fileBytes)

	assertGeneratedContent(t, fileBytes)
}

func assertGeneratedContent(t *testing.T, fileBytes []byte) {
	t.Helper()
	assert.True(t, len(fileBytes) > 0, "Content length is zero or negative")
}
