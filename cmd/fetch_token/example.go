package main

import (
	"context"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud/jwt"
)

func main() {
	conf := jwt.NewConfig(
		"App SID from https://dashboard.aspose.cloud/#/apps",
		"App Key from https://dashboard.aspose.cloud/#/apps",
	)

	ctx := context.Background()
	source := conf.TokenSource(ctx)
	token, err := source.Token()
	if err != nil {
		panic(err)
	}

	println(token.AccessToken)
}
