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

	fileName := filepath.Join(filepath.Dir("."), "..", "..", "..", "Code128.jpeg")

	generateParams := barcode.GenerateParams{
		BarcodeType: barcode.EncodeBarcodeTypeCode128,
		EncodeData: barcode.EncodeData{
			Data:     "Aspose.BarCode.Cloud",
			DataType: barcode.EncodeDataTypeStringData,
		},
	}

	fileBytes, _, err := client.GenerateAPI.BarcodeGenerateBodyPost(authCtx, generateParams)
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