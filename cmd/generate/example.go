package main

import (
	"context"
	"fmt"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/jwt"
	"os"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud"
	models "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/models"
)

func main() {
	jwtConf := jwt.NewConfig(
		"App SID from https://dashboard.aspose.cloud/#/apps",
		"App Key from https://dashboard.aspose.cloud/#/apps",
	)
	fileName := "testdata/generated.png"

	authCtx := context.WithValue(context.Background(), api.ContextJWT, jwtConf.TokenSource(context.Background()))

	client := api.NewAPIClient(api.NewConfiguration())

	data, resp, err := client.BarcodeApi.GetBarcodeGenerate(authCtx, string(models.EncodeBarcodeTypeCode128), "Go SDK", nil)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(data)
	}
	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	written, err := out.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Written %d bytes to file %s\n", written, fileName)
}
