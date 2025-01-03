# Aspose.BarCode Cloud SDK for Go

[![License](https://img.shields.io/github/license/aspose-barcode-cloud/aspose-barcode-cloud-go)](LICENSE)
[![Go](https://github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/actions/workflows/go.yml)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/aspose-barcode-cloud/aspose-barcode-cloud-go?label=module&sort=semver)](https://pkg.go.dev/github.com/aspose-barcode-cloud/aspose-barcode-cloud-go)

- API version: 3.0
- SDK version: 1.2412.0

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
    go get -u github.com/aspose-barcode-cloud/aspose-barcode-cloud-go@v1.2412.0
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

All URIs are relative to *<https://api.aspose.cloud/v3.0>*

Class | Method | HTTP request | Description
----- | ------ | ------------ | -----------
*BarcodeApi* | [**GetBarcodeGenerate**](docs/BarcodeApi.md#getbarcodegenerate) | **Get** /barcode/generate | Generate barcode.
*BarcodeApi* | [**GetBarcodeRecognize**](docs/BarcodeApi.md#getbarcoderecognize) | **Get** /barcode/{name}/recognize | Recognize barcode from a file on server.
*BarcodeApi* | [**PostBarcodeRecognizeFromUrlOrContent**](docs/BarcodeApi.md#postbarcoderecognizefromurlorcontent) | **Post** /barcode/recognize | Recognize barcode from an url or from request body. Request body can contain raw data bytes of the image with content-type \&quot;application/octet-stream\&quot;. An image can also be passed as a form field.
*BarcodeApi* | [**PostGenerateMultiple**](docs/BarcodeApi.md#postgeneratemultiple) | **Post** /barcode/generateMultiple | Generate multiple barcodes and return in response stream
*BarcodeApi* | [**PutBarcodeGenerateFile**](docs/BarcodeApi.md#putbarcodegeneratefile) | **Put** /barcode/{name}/generate | Generate barcode and save on server (from query params or from file with json or xml content)
*BarcodeApi* | [**PutBarcodeRecognizeFromBody**](docs/BarcodeApi.md#putbarcoderecognizefrombody) | **Put** /barcode/{name}/recognize | Recognition of a barcode from file on server with parameters in body.
*BarcodeApi* | [**PutGenerateMultiple**](docs/BarcodeApi.md#putgeneratemultiple) | **Put** /barcode/{name}/generateMultiple | Generate image with multiple barcodes and put new file on server
*BarcodeApi* | [**ScanBarcode**](docs/BarcodeApi.md#scanbarcode) | **Post** /barcode/scan | Quickly scan a barcode from an image.
*FileApi* | [**CopyFile**](docs/FileApi.md#copyfile) | **Put** /barcode/storage/file/copy/{srcPath} | Copy file
*FileApi* | [**DeleteFile**](docs/FileApi.md#deletefile) | **Delete** /barcode/storage/file/{path} | Delete file
*FileApi* | [**DownloadFile**](docs/FileApi.md#downloadfile) | **Get** /barcode/storage/file/{path} | Download file
*FileApi* | [**MoveFile**](docs/FileApi.md#movefile) | **Put** /barcode/storage/file/move/{srcPath} | Move file
*FileApi* | [**UploadFile**](docs/FileApi.md#uploadfile) | **Put** /barcode/storage/file/{path} | Upload file
*FolderApi* | [**CopyFolder**](docs/FolderApi.md#copyfolder) | **Put** /barcode/storage/folder/copy/{srcPath} | Copy folder
*FolderApi* | [**CreateFolder**](docs/FolderApi.md#createfolder) | **Put** /barcode/storage/folder/{path} | Create the folder
*FolderApi* | [**DeleteFolder**](docs/FolderApi.md#deletefolder) | **Delete** /barcode/storage/folder/{path} | Delete folder
*FolderApi* | [**GetFilesList**](docs/FolderApi.md#getfileslist) | **Get** /barcode/storage/folder/{path} | Get all files and folders within a folder
*FolderApi* | [**MoveFolder**](docs/FolderApi.md#movefolder) | **Put** /barcode/storage/folder/move/{srcPath} | Move folder
*StorageApi* | [**GetDiscUsage**](docs/StorageApi.md#getdiscusage) | **Get** /barcode/storage/disc | Get disc usage
*StorageApi* | [**GetFileVersions**](docs/StorageApi.md#getfileversions) | **Get** /barcode/storage/version/{path} | Get file versions
*StorageApi* | [**ObjectExists**](docs/StorageApi.md#objectexists) | **Get** /barcode/storage/exist/{path} | Check if file or folder exists
*StorageApi* | [**StorageExists**](docs/StorageApi.md#storageexists) | **Get** /barcode/storage/{storageName}/exist | Check if storage exists

## Documentation For Models

- [ApiError](docs/ApiError.md)
- [ApiErrorResponse](docs/ApiErrorResponse.md)
- [AustralianPostParams](docs/AustralianPostParams.md)
- [AutoSizeMode](docs/AutoSizeMode.md)
- [AvailableGraphicsUnit](docs/AvailableGraphicsUnit.md)
- [AztecEncodeMode](docs/AztecEncodeMode.md)
- [AztecParams](docs/AztecParams.md)
- [AztecSymbolMode](docs/AztecSymbolMode.md)
- [BarcodeResponse](docs/BarcodeResponse.md)
- [BarcodeResponseList](docs/BarcodeResponseList.md)
- [BorderDashStyle](docs/BorderDashStyle.md)
- [CaptionParams](docs/CaptionParams.md)
- [ChecksumValidation](docs/ChecksumValidation.md)
- [CodabarChecksumMode](docs/CodabarChecksumMode.md)
- [CodabarParams](docs/CodabarParams.md)
- [CodabarSymbol](docs/CodabarSymbol.md)
- [CodablockParams](docs/CodablockParams.md)
- [Code128Emulation](docs/Code128Emulation.md)
- [Code128EncodeMode](docs/Code128EncodeMode.md)
- [Code128Params](docs/Code128Params.md)
- [Code16KParams](docs/Code16KParams.md)
- [CodeLocation](docs/CodeLocation.md)
- [CouponParams](docs/CouponParams.md)
- [CustomerInformationInterpretingType](docs/CustomerInformationInterpretingType.md)
- [DataBarParams](docs/DataBarParams.md)
- [DataMatrixEccType](docs/DataMatrixEccType.md)
- [DataMatrixEncodeMode](docs/DataMatrixEncodeMode.md)
- [DataMatrixParams](docs/DataMatrixParams.md)
- [DataMatrixVersion](docs/DataMatrixVersion.md)
- [DecodeBarcodeType](docs/DecodeBarcodeType.md)
- [DiscUsage](docs/DiscUsage.md)
- [DotCodeEncodeMode](docs/DotCodeEncodeMode.md)
- [DotCodeParams](docs/DotCodeParams.md)
- [EciEncodings](docs/EciEncodings.md)
- [EnableChecksum](docs/EnableChecksum.md)
- [EncodeBarcodeType](docs/EncodeBarcodeType.md)
- [ErrorDetails](docs/ErrorDetails.md)
- [FileVersions](docs/FileVersions.md)
- [FilesList](docs/FilesList.md)
- [FilesUploadResult](docs/FilesUploadResult.md)
- [FontMode](docs/FontMode.md)
- [FontParams](docs/FontParams.md)
- [FontStyle](docs/FontStyle.md)
- [GeneratorParams](docs/GeneratorParams.md)
- [GeneratorParamsList](docs/GeneratorParamsList.md)
- [HanXinEncodeMode](docs/HanXinEncodeMode.md)
- [HanXinErrorLevel](docs/HanXinErrorLevel.md)
- [HanXinParams](docs/HanXinParams.md)
- [HanXinVersion](docs/HanXinVersion.md)
- [Itf14BorderType](docs/Itf14BorderType.md)
- [ItfParams](docs/ItfParams.md)
- [MacroCharacter](docs/MacroCharacter.md)
- [MaxiCodeEncodeMode](docs/MaxiCodeEncodeMode.md)
- [MaxiCodeMode](docs/MaxiCodeMode.md)
- [MaxiCodeParams](docs/MaxiCodeParams.md)
- [ModelError](docs/ModelError.md)
- [ObjectExist](docs/ObjectExist.md)
- [Padding](docs/Padding.md)
- [PatchCodeParams](docs/PatchCodeParams.md)
- [PatchFormat](docs/PatchFormat.md)
- [Pdf417CompactionMode](docs/Pdf417CompactionMode.md)
- [Pdf417ErrorLevel](docs/Pdf417ErrorLevel.md)
- [Pdf417MacroTerminator](docs/Pdf417MacroTerminator.md)
- [Pdf417Params](docs/Pdf417Params.md)
- [PostalParams](docs/PostalParams.md)
- [PresetType](docs/PresetType.md)
- [QrEncodeMode](docs/QrEncodeMode.md)
- [QrEncodeType](docs/QrEncodeType.md)
- [QrErrorLevel](docs/QrErrorLevel.md)
- [QrParams](docs/QrParams.md)
- [QrVersion](docs/QrVersion.md)
- [ReaderParams](docs/ReaderParams.md)
- [RegionPoint](docs/RegionPoint.md)
- [ResultImageInfo](docs/ResultImageInfo.md)
- [StorageExist](docs/StorageExist.md)
- [StorageFile](docs/StorageFile.md)
- [StructuredAppend](docs/StructuredAppend.md)
- [TextAlignment](docs/TextAlignment.md)
- [FileVersion](docs/FileVersion.md)

