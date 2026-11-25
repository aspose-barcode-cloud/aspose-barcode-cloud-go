package main

import (
	"context"
	"os"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
)

func main() {
	configuration := barcode.NewConfiguration()
	token := os.Getenv("TEST_CONFIGURATION_ACCESS_TOKEN")
	if token == "" {
		token = os.Getenv("TEST_JWT_ACCESS_TOKEN")
	}
	configuration.AddDefaultHeader("Authorization", "Bearer "+token)
	client := barcode.NewAPIClient(configuration)
	data, _, err := client.GenerateAPI.GenerateBatch(context.Background(), barcode.BatchGenerateRequest{
		ArchiveFormat: barcode.BatchArchiveFormatZip,
		Items: []barcode.BatchGenerateItem{{FileName: "qr.png", Generation: barcode.GenerateParams{
			BarcodeType: barcode.EncodeBarcodeTypeQR, EncodeData: barcode.EncodeData{Data: "Batch example"},
		}}},
	})
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("barcodes.zip", data, 0644); err != nil {
		panic(err)
	}
}
