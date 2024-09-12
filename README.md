# Aspose.BarCode Cloud SDK for Go

[![License](https://img.shields.io/github/license/aspose-barcode-cloud/aspose-barcode-cloud-go)](LICENSE)
[![Go](https://github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/actions/workflows/go.yml)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/aspose-barcode-cloud/aspose-barcode-cloud-go?label=module&sort=semver)](https://pkg.go.dev/github.com/aspose-barcode-cloud/aspose-barcode-cloud-go)

- API version: 4.0
- SDK version: 1.2407.0

## Demo applications

[Scan QR](https://products.aspose.app/barcode/scanqr) | [Generate Barcode](https://products.aspose.app/barcode/generate) | [Recognize Barcode](https://products.aspose.app/barcode/recognize)
:---: | :---: | :---:
[![ScanQR](https://products.aspose.app/barcode/scanqr/img/aspose_scanqr-app-48.png)](https://products.aspose.app/barcode/scanqr) | [![Generate](https://products.aspose.app/barcode/generate/img/aspose_generate-app-48.png)](https://products.aspose.app/barcode/generate) | [![Recognize](https://products.aspose.app/barcode/recognize/img/aspose_recognize-app-48.png)](https://products.aspose.app/barcode/recognize)
[**Generate Wi-Fi QR**](https://products.aspose.app/barcode/wifi-qr) | [**Embed Barcode**](https://products.aspose.app/barcode/embed) | [**Scan Barcode**](https://products.aspose.app/barcode/scan)
[![Wi-FiQR](https://products.aspose.app/barcode/embed/img/aspose_wifi-qr-app-48.png)](https://products.aspose.app/barcode/wifi-qr) | [![Embed](https://products.aspose.app/barcode/embed/img/aspose_embed-app-48.png)](https://products.aspose.app/barcode/embed) | [![Scan](https://products.aspose.app/barcode/embed/img/aspose_scan-app-48.png)](https://products.aspose.app/barcode/scan)

[Aspose.BarCode for Cloud](https://products.aspose.cloud/barcode/) is a REST API for Linear, 2D and postal barcode generation and recognition in the cloud. API recognizes and generates barcode images in a variety of formats. Barcode REST API allows to specify barcode image attributes like image width, height, border style and output image format in order to customize the generation process. Developers can also specify the barcode type and text attributes such as text location and font styles in order to suit the application requirements.

This repository contains Aspose.BarCode Cloud SDK for Go source code. This SDK allows you to work with Aspose.BarCode for Cloud REST APIs in your Go applications quickly and easily.

To use these SDKs, you will need Client Id and Client Secret which can be looked up at [Aspose Cloud Dashboard](https://dashboard.aspose.cloud/applications) (free registration in Aspose Cloud is required for this).

## Prerequisites

To use Aspose.BarCode Cloud SDK for Go you need to register an account with [Aspose Cloud](https://www.aspose.cloud) and lookup/create Client Secret and SID at [Cloud Dashboard](https://dashboard.aspose.cloud/applications). There is a free quota available. For more details, see [Aspose Cloud Pricing](https://purchase.aspose.cloud/).

## Installation

### Using Go Modules (recommended)

1. Go to existing module directory, or create a new module (see <https://blog.golang.org/using-go-modules>)
1. Run `go get` command

    ```shell script
    go get -u github.com/aspose-barcode-cloud/aspose-barcode-cloud-go@v1.2407.0
    ```

### Using GOPATH (for Go < 1.11 )

1. Run `go get` command outside module directory

    ```shell script
   go get -u github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode
   ```

## Sample usage

### Generate an image with barcode

The examples below show how you can generate QR barcode and save it into a local file using **barcode**:

```go
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

	opts := &barcode.BarcodeApiGetBarcodeGenerateOpts{
		TextLocation: optional.NewString(string(barcode.CodeLocationNone)),
	}

	data, _, err := client.BarcodeApi.GetBarcodeGenerate(authCtx,
		string(barcode.EncodeBarcodeTypeQR),
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

```

### Recognize a barcode on image

The examples below show how you can scan barcode text and type on the image using **barcode**:

```go
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

	imageFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(imageFile)

	client := barcode.NewAPIClient(barcode.NewConfiguration())
	authCtx := context.WithValue(context.Background(),
		barcode.ContextJWT,
		jwtConf.TokenSource(context.Background()))

	optionals := barcode.BarcodeApiScanBarcodeOpts{
		DecodeTypes: optional.NewInterface([]barcode.DecodeBarcodeType{
			barcode.DecodeBarcodeTypeQR,
		}),
	}

	recognized, _, err := client.BarcodeApi.ScanBarcode(
		authCtx,
		imageFile,
		&optionals)
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

```

## Dependencies

- github.com/antihax/optional
- github.com/google/uuid
- golang.org/x/oauth2

## Licensing

All Aspose.BarCode for Cloud SDKs, helper scripts and templates are licensed under [MIT License](LICENSE).

## Resources

- [**Website**](https://www.aspose.cloud)
- [**Product Home**](https://products.aspose.cloud/barcode/)
- [**Documentation**](https://docs.aspose.cloud/barcode/)
- [**Free Support Forum**](https://forum.aspose.cloud/c/barcode)
- [**Paid Support Helpdesk**](https://helpdesk.aspose.cloud/)
- [**Blog**](https://blog.aspose.cloud/categories/aspose.barcode-cloud-product-family/)

## Documentation for API Endpoints

All URIs are relative to *<https://barcode.qa.aspose.cloud/v4.0>*

Class | Method | HTTP request | Description
----- | ------ | ------------ | -----------
*GenerateAPI* | [**BarcodeGenerateBarcodeTypeGet**](docs/GenerateAPI.md#barcodegeneratebarcodetypeget) | **Get** /barcode/generate/{barcodeType} | Generate barcode using GET request with parameters in route and query string.
*GenerateAPI* | [**BarcodeGenerateBodyPost**](docs/GenerateAPI.md#barcodegeneratebodypost) | **Post** /barcode/generate-body | Generate barcode using POST request with parameters in body in json or xml format.
*GenerateAPI* | [**BarcodeGenerateFormPost**](docs/GenerateAPI.md#barcodegenerateformpost) | **Post** /barcode/generate-form | Generate barcode using POST request with parameters in url ecncoded form.
*RecognizeAPI* | [**BarcodeRecognizeBarcodeTypeGet**](docs/RecognizeAPI.md#barcoderecognizebarcodetypeget) | **Get** /barcode/recognize/{barcodeType} | Recognize barcode from file on server using GET requests with parameters in route and query string.
*RecognizeAPI* | [**BarcodeRecognizeBodyPost**](docs/RecognizeAPI.md#barcoderecognizebodypost) | **Post** /barcode/recognize-body | Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.
*RecognizeAPI* | [**BarcodeRecognizeFormPost**](docs/RecognizeAPI.md#barcoderecognizeformpost) | **Post** /barcode/recognize-form | Recognize barcode from file in request body using POST requests with parameters in multipart form.
*ScanAPI* | [**BarcodeScanBodyPost**](docs/ScanAPI.md#barcodescanbodypost) | **Post** /barcode/scan-body | Scan barcode from file in request body using POST requests with parameter in body in json or xml format.
*ScanAPI* | [**BarcodeScanFormPost**](docs/ScanAPI.md#barcodescanformpost) | **Post** /barcode/scan-form | Scan barcode from file in request body using POST requests with parameter in multipart form.
*ScanAPI* | [**BarcodeScanGet**](docs/ScanAPI.md#barcodescanget) | **Get** /barcode/scan | Scan barcode from file on server using GET requests with parameter in query string.

## Documentation For Models

- [ApiError](docs/ApiError.md)
- [ApiErrorResponse](docs/ApiErrorResponse.md)
- [AvailableBarCodeImageFormat](docs/AvailableBarCodeImageFormat.md)
- [AvailableGraphicsUnit](docs/AvailableGraphicsUnit.md)
- [BarcodeImageParams](docs/BarcodeImageParams.md)
- [BarcodeResponse](docs/BarcodeResponse.md)
- [BarcodeResponseList](docs/BarcodeResponseList.md)
- [CodeLocation](docs/CodeLocation.md)
- [DecodeBarcodeType](docs/DecodeBarcodeType.md)
- [EncodeBarcodeType](docs/EncodeBarcodeType.md)
- [EncodeData](docs/EncodeData.md)
- [EncodeDataType](docs/EncodeDataType.md)
- [GenerateParams](docs/GenerateParams.md)
- [RecognitionImageKind](docs/RecognitionImageKind.md)
- [RecognitionMode](docs/RecognitionMode.md)
- [RecognizeBase64Request](docs/RecognizeBase64Request.md)
- [RegionPoint](docs/RegionPoint.md)
- [ScanBase64Request](docs/ScanBase64Request.md)
- [TextAlignment](docs/TextAlignment.md)

