package main

import (
	"context"
	"fmt"
	"os"

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
	imageUrl := "https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png"

	result, _, err := client.RecognizeAPI.Recognize(authCtx, barcode.DecodeBarcodeTypeQR, imageUrl, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File '%s' recognized, result: '%s'\n", imageUrl, result.Barcodes[0].BarcodeValue)
}
