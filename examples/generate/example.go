package main

import (
	"context"
	"fmt"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
	"os"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
)

func main() {
	jwtConf := jwt.NewConfig(
		"App SID from https://dashboard.aspose.cloud/#/apps",
		"App Key from https://dashboard.aspose.cloud/#/apps",
	)
	fileName := "testdata/generated.png"

	authCtx := context.WithValue(context.Background(), barcode.ContextJWT, jwtConf.TokenSource(context.Background()))

	client := barcode.NewAPIClient(barcode.NewConfiguration())

	data, _, err := client.BarcodeApi.GetBarcodeGenerate(authCtx, string(barcode.EncodeBarcodeTypeQR), "Go SDK example", nil)
	if err != nil {
		panic(err)
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
