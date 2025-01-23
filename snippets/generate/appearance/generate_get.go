package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/antihax/optional"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
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

func Generate(client *barcode.APIClient, authCtx context.Context) error {
	fileName, err := filepath.Abs(filepath.Join("testdata", "qr.png"))

	opts := &barcode.GenerateAPIGenerateOpts{
		ImageFormat: optional.NewInterface(barcode.BarcodeImageFormatPng),
	}

	fileBytes, _, err := client.GenerateAPI.Generate(authCtx, barcode.EncodeBarcodeTypeQR, "Aspose.BarCode.Cloud", opts)
	if err != nil {
		return fmt.Errorf("error generating barcode: %v", err)
	}

	err = os.WriteFile(fileName, fileBytes, 0644)
	if err != nil {
		return fmt.Errorf("error saving barcode to file: %v", err)
	}

	fmt.Printf("File '%s' generated.\n", fileName)
	return nil
}

func main() {
	client, authCtx, err := makeConfiguration()
	if err != nil {
		fmt.Printf("Error setting up configuration: %v\n", err)
		return
	}

	err = Generate(client, authCtx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
