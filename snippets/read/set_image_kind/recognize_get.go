package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"io/ioutil"

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

	imageURL := "https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png"

	base64Request := barcode.RecognizeBase64Request{
		BarcodeTypes: []barcode.DecodeBarcodeType{barcode.DecodeBarcodeTypeQR},
		ImageURL:     imageURL,
	}



	result, _, err := client.RecognizeAPI.BarcodeRecognizeGet(authCtx, base64Request)
	if err != nil {
		panic(err)
	}

	if len(result.Barcodes) > 0 {
		fmt.Printf("File '%s' recognized, result: '%s'\n", imageURL, result.Barcodes[0].BarcodeValue)
	} else {
		fmt.Printf("File '%s' recognized, but no barcodes found.\n", imageURL)
	}
}