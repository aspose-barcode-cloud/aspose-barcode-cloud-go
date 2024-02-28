package main

import (
	"context"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
)

func main() {
	conf := jwt.NewConfig(
		"Client Id from https://dashboard.aspose.cloud/applications",
		"Client Secret from https://dashboard.aspose.cloud/applications",
	)

	ctx := context.Background()
	source := conf.TokenSource(ctx)
	token, err := source.Token()
	if err != nil {
		panic(err)
	}

	println(token.AccessToken)
}
