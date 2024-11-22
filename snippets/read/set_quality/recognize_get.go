package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

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

	fileURL := "https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png"
	imageBytes, err := ioutil.ReadFile(fileURL)
	if err != nil {
		panic(err)
	}

	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)

	base64Request := barcode.RecognizeBase64Request{
		BarcodeTypes:           []barcode.DecodeBarcodeType{barcode.DecodeBarcodeTypeQR},
		FileBase64:            imageBase64,
		RecognitionMode:       barcode.RecognitionModeFast,
		RecognitionImageKind:  barcode.RecognitionImageKindPhoto,
	}


	result, _, err := client.RecognizeAPI.BarcodeRecognizeBodyPost(authCtx, base64Request)
	if err != nil {
		panic(err)
	}

	if len(result.Barcodes) > 0 {
		fmt.Printf("File recognized, result: '%s'\n", result.Barcodes[0].BarcodeValue)
	} else {
		fmt.Println("File recognized, but no barcodes found.")
	}
}