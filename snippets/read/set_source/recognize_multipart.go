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

	filePath, err := filepath.Abs(filepath.Join("testdata", "qr.png"))

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	opts := &barcode.RecognizeAPIRecognizeMultipartOpts{}
	response, _, err := client.RecognizeAPI.RecognizeMultipart(authCtx, barcode.DecodeBarcodeTypeQR, file, opts)

	if err != nil {
		panic(err)
	}

	if len(response.Barcodes) > 0 {
		fmt.Printf("File '%s' recognized, result: '%s'\n", filePath, response.Barcodes[0].BarcodeValue)
	} else {
		fmt.Printf("No barcodes found in '%s'.\n", filePath)
	}
}
