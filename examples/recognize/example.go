package main

import (
	"context"
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
)

func main() {
	jwtConf := jwt.NewConfig(
		"Client Id from https://dashboard.aspose.cloud/applications",
		"Client Secret from https://dashboard.aspose.cloud/applications",
	)
	fileName := "testdata/generated.png"

	imageFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(imageFile)

	client := barcode.NewAPIClient(barcode.NewConfiguration())
	authCtx := context.WithValue(context.Background(),
		barcode.ContextJWT,
		jwtConf.TokenSource(context.Background()))

	optionals := barcode.BarcodeApiScanBarcodeOpts{
		DecodeTypes: optional.NewInterface([]barcode.DecodeBarcodeType{
			barcode.DecodeBarcodeTypeQR,
		}),
	}

	recognized, _, err := client.BarcodeApi.ScanBarcode(
		authCtx,
		imageFile,
		&optionals)
	if err != nil {
		panic(err)
	}

	if len(recognized.Barcodes) == 0 {
		fmt.Printf("No barcodes in %s", fileName)
	}

	for i, oneBarcode := range recognized.Barcodes {
		fmt.Printf("Recognized #%d: %s %s", i+1, oneBarcode.Type, oneBarcode.BarcodeValue)
	}
}
