package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
)

func makeConfiguration() (*barcode.APIClient, context.Context, error) {
	jwtToken := os.Getenv("TEST_CONFIGURATION_JWT_TOKEN")
	if jwtToken != "" {
		config := barcode.NewConfiguration()
		config.AddDefaultHeader("Authorization", "Bearer "+jwtToken)
		client := barcode.NewAPIClient(config)
		authCtx := context.Background()
		return client, authCtx, nil
	}

	jwtConf := jwt.NewConfig(
		"Client Id from https://dashboard.aspose.cloud/applications",
		"Client Secret from https://dashboard.aspose.cloud/applications",
	)

	authCtx := context.WithValue(context.Background(),
		barcode.ContextJWT,
		jwtConf.TokenSource(context.Background()))

	client := barcode.NewAPIClient(barcode.NewConfiguration())

	return client, authCtx, nil
}

func main() {
	client, authCtx, err := makeConfiguration()
	if err != nil {
		fmt.Printf("Error setting up configuration: %v\n", err)
		return
	}

	fileName, err := filepath.Abs(filepath.Join("testdata", "qr.png"))

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	response, _, err := client.ScanAPI.ScanMultipart(authCtx, file)
	if err != nil {
		panic(err)
	}

	if len(response.Barcodes) > 0 {
		fmt.Printf("File '%s' recognized, results: value: '%s', type: %s\n",
			fileName, response.Barcodes[0].BarcodeValue, response.Barcodes[0].Type)
	} else {
		fmt.Printf("File '%s' recognized, but no barcodes found.\n", fileName)
	}

}
