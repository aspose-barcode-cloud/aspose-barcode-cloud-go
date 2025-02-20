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

	fileName, err := filepath.Abs(filepath.Join("testdata", "..", "Code39.png"))

	fileBytes, _, err := client.GenerateAPI.GenerateMultipart(authCtx,
		barcode.EncodeBarcodeTypeCode39, "Aspose",
		&barcode.GenerateAPIGenerateMultipartOpts{
			ImageFormat:     optional.NewInterface(barcode.BarcodeImageFormatGif),
			ForegroundColor: optional.NewString("#008000"),
			BackgroundColor: optional.NewString("#FFFF00"),
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
