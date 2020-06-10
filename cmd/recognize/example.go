package main

import (
	"context"
	"fmt"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/jwt"
	"os"

	"github.com/antihax/optional"
	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud"
	models "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/models"
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

	client := api.NewAPIClient(api.NewConfiguration())
	authCtx := context.WithValue(context.Background(), api.ContextJWT, jwtConf.TokenSource(context.Background()))

	optionals := &api.BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts{
		Preset: optional.NewString(string(models.PresetTypeHighPerformance)),
		Image:  optional.NewInterface(file),
	}

	recognized, _, err := client.BarcodeApi.PostBarcodeRecognizeFromUrlOrContent(authCtx, optionals)
	if err != nil {
		panic(err)
	}

	if len(recognized.Barcodes) == 0 {
		fmt.Printf("No barcodes in %s", fileName)
	}

	for i, barcode := range recognized.Barcodes {
		fmt.Printf("Recognized #%d: %s %s", i+1, barcode.Type_, barcode.BarcodeValue)
	}
}
