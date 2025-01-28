package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/antihax/optional"

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

	fileName, err := filepath.Abs(filepath.Join("testdata", "qr.png"))

	fileBytes, _, err := client.GenerateAPI.Generate(authCtx, barcode.EncodeBarcodeTypeQR, "Aspose.BarCode.Cloud",
		&barcode.GenerateAPIGenerateOpts{
			ImageHeight: optional.NewFloat32(200),
			ImageWidth:  optional.NewFloat32(200),
			Resolution:  optional.NewFloat32(300),
			ImageFormat: optional.NewInterface(barcode.BarcodeImageFormatPng),
		})
	if err != nil {
		fmt.Printf("Error generating barcode: %v\n", err)
		return
	}

	err = os.WriteFile(fileName, fileBytes, 0644)
	if err != nil {
		fmt.Printf("Error saving barcode to file: %v\n", err)
		return
	}

	fmt.Printf("File '%s' generated.\n", fileName)
}
