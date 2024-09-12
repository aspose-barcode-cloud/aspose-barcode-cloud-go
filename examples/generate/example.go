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

	authCtx := context.WithValue(context.Background(),
		barcode.ContextJWT,
		jwtConf.TokenSource(context.Background()))

	client := barcode.NewAPIClient(barcode.NewConfiguration())

	opts := &barcode.GenerateAPIGenerateOpts{
		TextLocation: optional.NewInterface(barcode.CodeLocationNone),
	}

	data, _, err := client.GenerateAPI.Generate(authCtx,
		barcode.EncodeBarcodeTypeQR,
		"Go SDK example",
		opts)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	written, err := out.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Written %d bytes to file %s\n", written, fileName)
}
