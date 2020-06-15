package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

//TestFolder Aspose.Storage folder for test files
//noinspection GoNameStartsWithPackageName
var TestFolder = fmt.Sprintf("BarcodeTests/%s", uuid.New())

//TestConfig configuration for testing
//noinspection GoNameStartsWithPackageName
var TestConfig, _ = NewConfig("configuration.json")

//NewClientForTests creates new Client with TestConfig
func NewClientForTests() *barcode.APIClient {
	return barcode.NewAPIClient(&TestConfig.APIConfig)
}

//NewAuthContextForTests context for testing
func NewAuthContextForTests() context.Context {
	return context.WithValue(context.Background(), barcode.ContextJWT, TestConfig.JwtConfig.TokenSource(context.Background()))
}

func uploadTestFile(ctx context.Context, t *testing.T, client *barcode.APIClient, uploadedFilename string) {
	file, err := os.Open("../testdata/pdf417Sample.png")
	require.Nil(t, err)

	uploaded, _, err := client.FileApi.UploadFile(ctx, fmt.Sprintf("%s/%s", TestFolder, uploadedFilename), file, nil)
	require.Nil(t, err)
	require.Empty(t, uploaded.Errors)
	require.NotEmpty(t, uploaded.Uploaded)
	require.Equal(t, uploadedFilename, uploaded.Uploaded[0])
}
