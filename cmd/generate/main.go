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
	fileName := "testdata/generated.png"

	apiCfg := api.NewConfiguration()
	apiCfg.BasePath = "https://api-qa.aspose.cloud/v3.0"
	client := api.NewAPIClient(apiCfg)


	jwtConf := jwt.NewConfig(
		"barcode.cloud",
		"gW9S5WZX8xQ5ArVM",
	)
	// TODO: Use base URL from apiCfg
	jwtConf.TokenURL = "https://api-qa.aspose.cloud/connect/token"
	tokenSource := jwtConf.TokenSource(context.TODO())
	authCtx := context.WithValue(context.TODO(), api.ContextJWT, tokenSource)

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
