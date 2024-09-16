package test

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/stretchr/testify/require"
)

// TestConfigurationFile with configuration
var TestConfigurationFile = "configuration.json"

// TestEnvPrefix for environment variables
var TestEnvPrefix = "TEST"

//Test data folder

func GetTestDataFolder() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(currentPath, "..", "testdata")
}

// NewClientForTests creates new Client with TestConfig
func NewClientForTests() (*barcode.APIClient, error) {
	testConfig, err := NewTestConfig(TestConfigurationFile, TestEnvPrefix)
	if err != nil {
		return nil, err
	}

	return barcode.NewAPIClient(&testConfig.APIConfig), nil
}

// NewAuthContextForTests context for testing
func NewAuthContextForTests() (context.Context, error) {
	testConfig, err := NewTestConfig(TestConfigurationFile, TestEnvPrefix)
	if err != nil {
		return nil, err
	}

	return context.WithValue(context.Background(), barcode.ContextJWT, testConfig.JwtConfig.TokenSource(context.Background())), nil
}

func setup(t *testing.T) (*barcode.APIClient, context.Context) {
	// Initialize Auth context and Client
	authCtx, err := NewAuthContextForTests()
	require.Nil(t, err)

	client, err := NewClientForTests()
	require.Nil(t, err)
	return client, authCtx
}

func readFileContent(t *testing.T, fileName string) []byte {
	file, err := os.Open(fileName)
	require.Nil(t, err)
	bytes, err := io.ReadAll(io.Reader(file))
	require.Nil(t, err)
	file.Close()
	return bytes
}
