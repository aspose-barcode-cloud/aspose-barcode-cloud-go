package main

import (
	"context"
	"fmt"
	"os"

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

func Recognize() {
	client, authCtx, err := makeConfiguration()
	if err != nil {
		fmt.Printf("Error setting up configuration: %v\n", err)
		return
	}

	fileUrl := "https://products.aspose.app/barcode/scan/img/how-to/scan/step2.png"
	opts := &barcode.RecognizeAPIRecognizeOpts{}

	result, _, err := client.RecognizeAPI.Recognize(authCtx, barcode.DecodeBarcodeTypeQR, fileUrl, opts)
	if err != nil {
		fmt.Printf("Error recognizing barcode: %v\n", err)
		return
	}

	if len(result.Barcodes) > 0 {
		fmt.Printf("Recognized barcode value: '%s'\n", result.Barcodes[0].BarcodeValue)
	} else {
		fmt.Println("No barcodes recognized.")
	}
}

func main() {
	Recognize()
}
