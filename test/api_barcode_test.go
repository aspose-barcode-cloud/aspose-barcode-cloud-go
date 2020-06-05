package test

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var TestFolder = fmt.Sprintf("BarcodeTests/%s", uuid.New())

func NewClientForTests() *api.APIClient {
	cfg := api.NewConfiguration()
	cfg.BasePath = "http://localhost:47972/v3.0"
	client := api.NewAPIClient(cfg)

	return client
}

func NewContextForTests() context.Context {
	ctx := context.WithValue(context.TODO(), api.ContextAccessToken, "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1OTEyNzc1MDMsImV4cCI6MTU5MTM2MzkwMywiaXNzIjoiaHR0cHM6Ly9hcGktcWEuYXNwb3NlLmNsb3VkIiwiYXVkIjpbImh0dHBzOi8vYXBpLXFhLmFzcG9zZS5jbG91ZC9yZXNvdXJjZXMiLCJhcGkucGxhdGZvcm0iLCJhcGkucHJvZHVjdHMiXSwiY2xpZW50X2lkIjoiYmFyY29kZS5jbG91ZCIsImNsaWVudF91c2FnZSI6InBvc3QiLCJjbGllbnRfc3RvcmFnZSI6ImdldCIsImNsaWVudF9yZXN0cmljdGlvbnMiOiJnZXQiLCJjbGllbnRfZGVmYXVsdF9zdG9yYWdlIjoiMEZCOUI2MjctNzY5QS00MjJBLUJBMkUtODhBMUUzNEFDODA1IiwiY2xpZW50X2lkU3J2SWQiOiIxODkyMDEiLCJzY29wZSI6WyJhcGkucGxhdGZvcm0iLCJhcGkucHJvZHVjdHMiXX0.ZXYVEiZncgoEVHaFFg0p4MGGBlak2Z7OKzXhYu6WukEyYa9fCXw6vAoLUJqsam2-FKpXhZtssLBzyPrRP7kqtJ1gTNEMsSEDbK2HwLr6WzXfCn81h7VJiV6IOsqn0KY0K8Yb-Ge27G_325hJKzPp712hN6Cl1bTjG9lnfnPdZ2g1SGYMcQ_6062hTukXHDYTYY_1yr-SM_TfWOzIuQ8z1Wop_lmX2ur3hRPP1OyS33hwbxGbrErzXvlBFPxf5GWCSzZCaBk-JQYMRMH0QKUE-RPXtrwCwsd6ejvGN74JQZXpt__hUjGh7FaKmaOFfSL4H4VnMzCzbvmE9VH7D-Ol4g")

	return ctx
}

func TestGetBarcodeGenerate(t *testing.T) {
	imgData, response, err := NewClientForTests().BarcodeApi.GetBarcodeGenerate(NewContextForTests(), string(api.CODE128_EncodeBarcodeType), "Go SDK", nil)
	if err != nil {
		t.Errorf("Api error: %s", err)
	}

	require.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "image/png", response.Header.Get("content-type"))
	assert.Greater(t, len(imgData), 0)
}

func TestPutBarcodeGenerateFile(t *testing.T) {
	genImgInfo, response, err := NewClientForTests().BarcodeApi.PutBarcodeGenerateFile(
		NewContextForTests(),
		"TestPutBarcodeGenerateFile.png",
		string(api.CODE128_EncodeBarcodeType),
		"Go SDK",
		&api.BarcodeApiPutBarcodeGenerateFileOpts{
			Folder: optional.NewString(TestFolder),
		},
	)
	if err != nil {
		t.Errorf("Api error: %s", err)
	}

	require.Equal(t, 200, response.StatusCode)
	assert.Greater(t, genImgInfo.FileSize, int64(0))
	assert.Greater(t, genImgInfo.ImageWidth, int32(0))
	assert.Greater(t, genImgInfo.ImageHeight, int32(0))
}

func TestPutGenerateMultiple(t *testing.T) {
	genImgInfo, response, err := NewClientForTests().BarcodeApi.PutGenerateMultiple(
		NewContextForTests(),
		"TestPutBarcodeGenerateFile.png",
		api.GeneratorParamsList{
			BarcodeBuilders: []api.GeneratorParams{{
				TypeOfBarcode: api.CODE128_EncodeBarcodeType,
				Text:          "First barcode",
			}, {
				TypeOfBarcode: api.QR_EncodeBarcodeType,
				Text:          "Second barcode",
			},
			},
			XStep: 0,
			YStep: 0,
		},
		&api.BarcodeApiPutGenerateMultipleOpts{
			Folder: optional.NewString(TestFolder),
		},
	)
	if err != nil {
		t.Errorf("Api error: %s", err)
	}

	require.Equal(t, 200, response.StatusCode)
	assert.Greater(t, genImgInfo.FileSize, int64(0))
	assert.Greater(t, genImgInfo.ImageWidth, int32(0))
	assert.Greater(t, genImgInfo.ImageHeight, int32(0))
}

func TestPostGenerateMultiple(t *testing.T) {
	imgData, response, err := NewClientForTests().BarcodeApi.PostGenerateMultiple(
		NewContextForTests(),
		api.GeneratorParamsList{
			BarcodeBuilders: []api.GeneratorParams{{
				TypeOfBarcode: api.CODE128_EncodeBarcodeType,
				Text:          "First barcode",
			}, {
				TypeOfBarcode: api.QR_EncodeBarcodeType,
				Text:          "Second barcode",
			},
			},
			XStep: 0,
			YStep: 0,
		},
		&api.BarcodeApiPostGenerateMultipleOpts{
			Format: optional.NewString("png"),
		},
	)
	if err != nil {
		t.Errorf("Api error: %s", err)
	}

	require.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "image/png", response.Header.Get("content-type"))
	assert.Greater(t, len(imgData), 0)

}

func TestPostBarcodeRecognizeFromUrlOrContent(t *testing.T) {
	fileName := "../testdata/pdf417Sample.png"

	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("File '%s' open error: %s", fileName, err)
		t.FailNow()
	}
	defer file.Close()

	optionals := &api.BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts{
		Preset: optional.NewString(string(api.HIGH_PERFORMANCE_PresetType)),
		Image:  optional.NewInterface(file),
	}
	recognized, response, err := NewClientForTests().BarcodeApi.PostBarcodeRecognizeFromUrlOrContent(NewContextForTests(), optionals)
	if err != nil {
		t.Errorf("Api error: %s", err)
	}

	require.Equal(t, 200, response.StatusCode)
	require.Equal(t, 1, len(recognized.Barcodes))
	assert.Equal(t, string(api.PDF417_DecodeBarcodeType), recognized.Barcodes[0].Type_)
	assert.Equal(t, "Aspose.BarCode for Cloud Sample", recognized.Barcodes[0].BarcodeValue)
	require.Greater(t, len(recognized.Barcodes[0].Region), 0)
	assert.Greater(t, recognized.Barcodes[0].Region[0].X, int32(0))
	assert.Greater(t, recognized.Barcodes[0].Region[0].Y, int32(0))
}
