package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

//TestConfigurationFile with configuration
var TestConfigurationFile = "configuration.json"

//TestEnvPrefix for environment variables
var TestEnvPrefix = "TEST"

//TestFolder Aspose.Storage folder for test files
//noinspection GoNameStartsWithPackageName
var TestFolder = fmt.Sprintf("BarcodeTests/%s", uuid.New())

//NewClientForTests creates new Client with TestConfig
func NewClientForTests() (*barcode.APIClient, error) {
	testConfig, err := NewTestConfig(TestConfigurationFile, TestEnvPrefix)
	if err != nil {
		return nil, err
	}

	return barcode.NewAPIClient(&testConfig.APIConfig), nil
}

//NewAuthContextForTests context for testing
func NewAuthContextForTests() (context.Context, error) {
	testConfig, err := NewTestConfig(TestConfigurationFile, TestEnvPrefix)
	if err != nil {
		return nil, err
	}

	return context.WithValue(context.Background(), barcode.ContextJWT, testConfig.JwtConfig.TokenSource(context.Background())), nil
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

func readFileContent(t *testing.T, fileName string) []byte {
	file, err := os.Open(fileName)
	require.Nil(t, err)
	bytes, err := ioutil.ReadAll(file)
	require.Nil(t, err)
	file.Close()
	return bytes
}
