package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode/jwt"
)

func makeConfiguration() (*barcode.APIClient, context.Context, error) {
	jwtToken := os.Getenv("TEST_JWT_ACCESS_TOKEN")
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

	fileName, err := filepath.Abs(filepath.Join("testdata", "Pdf417.png"))

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	result, _, err := client.RecognizeAPI.RecognizeMultipart(authCtx, barcode.DecodeBarcodeTypePdf417, file, nil)
	if err != nil {
		panic(err)
	}

	if len(result.Barcodes) > 0 {
		fmt.Printf("File '%s' recognized, result: '%s'\n", fileName, result.Barcodes[0].BarcodeValue)
	} else {
		fmt.Printf("File '%s' recognized, but no barcodes found.\n", fileName)
	}
}
