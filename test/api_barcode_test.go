package test

import (
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBarcodeGenerate(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	imgData, response, err := client.BarcodeApi.GetBarcodeGenerate(authCtx, string(barcode.EncodeBarcodeTypeCode128), "Go SDK", nil)
	require.Nil(t, err)
	require.Equal(t, 200, response.StatusCode)

	assert.Equal(t, "image/png", response.Header.Get("content-type"))
	assert.Greater(t, len(imgData), 0)
}

func TestGetBarcodeRecognize(t *testing.T) {
	uploadedFilename := "TestGetBarcodeRecognize.png"

	client, err := NewClientForTests()
	require.Nil(t, err)

	ctx, err := NewAuthContextForTests()
	require.Nil(t, err)

	uploadTestFile(ctx, t, client, uploadedFilename)

	recognized, _, err := client.BarcodeApi.GetBarcodeRecognize(ctx, uploadedFilename, &barcode.BarcodeApiGetBarcodeRecognizeOpts{
		Folder: optional.NewString(TestFolder),
		Preset: optional.NewString(string(barcode.PresetTypeHighPerformance)),
	})
	require.Nil(t, err)
	require.Equal(t, 1, len(recognized.Barcodes))
	assert.Equal(t, "Pdf417", recognized.Barcodes[0].Type)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", recognized.Barcodes[0].BarcodeValue)
	require.Equal(t, 4, len(recognized.Barcodes[0].Region))
	assert.Greater(t, recognized.Barcodes[0].Region[0].X, int32(0))
	assert.Greater(t, recognized.Barcodes[0].Region[0].Y, int32(0))
}

func TestPostBarcodeRecognizeFromUrlOrContent_File(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	fileName := "../testdata/pdf417Sample.png"

	file, err := os.Open(fileName)
	require.Nil(t, err)
	defer file.Close()

	optionals := barcode.BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts{
		Preset: optional.NewString(string(barcode.PresetTypeHighPerformance)),
		Image:  optional.NewInterface(file),
	}
	recognized, _, err := client.BarcodeApi.PostBarcodeRecognizeFromUrlOrContent(authCtx, &optionals)
	require.Nil(t, err)
	require.Equal(t, 1, len(recognized.Barcodes))

	assert.Equal(t, string(barcode.DecodeBarcodeTypePdf417), recognized.Barcodes[0].Type)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", recognized.Barcodes[0].BarcodeValue)
	require.Greater(t, len(recognized.Barcodes[0].Region), 0)
	assert.Greater(t, recognized.Barcodes[0].Region[0].X, int32(0))
	assert.Greater(t, recognized.Barcodes[0].Region[0].Y, int32(0))
}

func TestPostBarcodeRecognizeFromUrlOrContent_Bytes(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	fileName := "../testdata/pdf417Sample.png"

	bytes := readFileContent(t, fileName)

	optionals := barcode.BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts{
		Preset: optional.NewString(string(barcode.PresetTypeHighPerformance)),
		Image:  optional.NewInterface(bytes),
	}
	recognized, _, err := client.BarcodeApi.PostBarcodeRecognizeFromUrlOrContent(authCtx, &optionals)
	require.Nil(t, err)
	require.Equal(t, 1, len(recognized.Barcodes))

	assert.Equal(t, string(barcode.DecodeBarcodeTypePdf417), recognized.Barcodes[0].Type)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", recognized.Barcodes[0].BarcodeValue)
	require.Greater(t, len(recognized.Barcodes[0].Region), 0)
	assert.Greater(t, recognized.Barcodes[0].Region[0].X, int32(0))
	assert.Greater(t, recognized.Barcodes[0].Region[0].Y, int32(0))
}

func TestPostGenerateMultiple(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	imgData, response, err := client.BarcodeApi.PostGenerateMultiple(
		authCtx,
		barcode.GeneratorParamsList{
			BarcodeBuilders: []barcode.GeneratorParams{{
				TypeOfBarcode: barcode.EncodeBarcodeTypeCode128,
				Text:          "First barcode",
			}, {
				TypeOfBarcode: barcode.EncodeBarcodeTypeQR,
				Text:          "Second barcode",
			},
			},
			XStep: 0,
			YStep: 0,
		},
		&barcode.BarcodeApiPostGenerateMultipleOpts{
			Format: optional.NewString("png"),
		},
	)
	require.Nil(t, err)
	require.Equal(t, 200, response.StatusCode)

	assert.Equal(t, "image/png", response.Header.Get("content-type"))
	assert.Greater(t, len(imgData), 0)

}

func TestPutBarcodeGenerateFile(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	genImgInfo, _, err := client.BarcodeApi.PutBarcodeGenerateFile(
		authCtx,
		"TestPutBarcodeGenerateFile.png",
		string(barcode.EncodeBarcodeTypeCode128),
		"Go SDK",
		&barcode.BarcodeApiPutBarcodeGenerateFileOpts{
			Folder: optional.NewString(TestFolder),
		},
	)
	require.Nil(t, err)

	assert.Greater(t, genImgInfo.FileSize, int64(0))
	assert.Greater(t, genImgInfo.ImageWidth, int32(0))
	assert.Greater(t, genImgInfo.ImageHeight, int32(0))
}

func TestPutBarcodeRecognizeFromBody(t *testing.T) {
	uploadedFilename := "TestPutBarcodeRecognizeFromBody.png"

	client, err := NewClientForTests()
	require.Nil(t, err)
	ctx, err := NewAuthContextForTests()
	require.Nil(t, err)

	uploadTestFile(ctx, t, client, uploadedFilename)

	recognized, _, err := client.BarcodeApi.PutBarcodeRecognizeFromBody(ctx, uploadedFilename, barcode.ReaderParams{
		Preset: barcode.PresetTypeHighPerformance,
	},
		&barcode.BarcodeApiPutBarcodeRecognizeFromBodyOpts{
			Folder: optional.NewString(TestFolder),
		},
	)
	require.Nil(t, err)

	require.Equal(t, 1, len(recognized.Barcodes))
	assert.Equal(t, "Pdf417", recognized.Barcodes[0].Type)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", recognized.Barcodes[0].BarcodeValue)
	require.Equal(t, 4, len(recognized.Barcodes[0].Region))
	assert.Greater(t, recognized.Barcodes[0].Region[0].X, int32(0))
	assert.Greater(t, recognized.Barcodes[0].Region[0].Y, int32(0))
}

func TestPutGenerateMultiple(t *testing.T) {
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)

	genImgInfo, _, err := client.BarcodeApi.PutGenerateMultiple(
		authCtx,
		"TestPutBarcodeGenerateFile.png",
		barcode.GeneratorParamsList{
			BarcodeBuilders: []barcode.GeneratorParams{{
				TypeOfBarcode: barcode.EncodeBarcodeTypeCode128,
				Text:          "First barcode",
			}, {
				TypeOfBarcode: barcode.EncodeBarcodeTypeQR,
				Text:          "Second barcode",
			},
			},
			XStep: 0,
			YStep: 0,
		},
		&barcode.BarcodeApiPutGenerateMultipleOpts{
			Folder: optional.NewString(TestFolder),
		},
	)
	require.Nil(t, err)

	assert.Greater(t, genImgInfo.FileSize, int64(0))
	assert.Greater(t, genImgInfo.ImageWidth, int32(0))
	assert.Greater(t, genImgInfo.ImageHeight, int32(0))
}
