package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
	"os"
)

func main() {
	jwtConf := jwt.NewConfig(
		"App SID from https://dashboard.aspose.cloud/#/apps",
		"App Key from https://dashboard.aspose.cloud/#/apps",
	)
	fileName := "testdata/pdf417Sample.png"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	client := barcode.NewAPIClient(barcode.NewConfiguration())
	authCtx := context.WithValue(context.Background(), barcode.ContextJWT, jwtConf.TokenSource(context.Background()))

	optionals := barcode.BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts{
		Preset: optional.NewString(string(barcode.PresetTypeHighPerformance)),
		Image:  optional.NewInterface(file),
	}

	recognized, _, err := client.BarcodeApi.PostBarcodeRecognizeFromUrlOrContent(authCtx, &optionals)
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
