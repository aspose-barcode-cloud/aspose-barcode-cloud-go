package main

import (
	"context"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/jwt"
)

func main() {
	conf := jwt.NewConfig(
		"f60b7a02-6306-4cc9-b423-96c7591b699b",
		"4bff3dd4c702380388d9e0eb620edbf4",
	)

	ctx := context.Background()
	source := conf.TokenSource(ctx)
	token, err := source.Token()
	if err != nil {
		panic(err)
	}

	println(token.AccessToken)
}
