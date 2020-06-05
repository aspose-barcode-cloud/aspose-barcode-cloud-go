# Go API client for aspose_barcode_cloud

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 3.0
- Package version: 20.5.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./aspose_barcode_cloud"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.aspose.cloud/v3.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BarcodeApi* | [**GetBarcodeGenerate**](docs/BarcodeApi.md#getbarcodegenerate) | **Get** /barcode/generate | Generate barcode.             
*BarcodeApi* | [**GetBarcodeRecognize**](docs/BarcodeApi.md#getbarcoderecognize) | **Get** /barcode/{name}/recognize | Recognize barcode from a file on server.             
*BarcodeApi* | [**PostBarcodeRecognizeFromUrlOrContent**](docs/BarcodeApi.md#postbarcoderecognizefromurlorcontent) | **Post** /barcode/recognize | Recognize barcode from an url or from request body. Request body can contain raw data bytes of the image or encoded with base64.             
*BarcodeApi* | [**PostGenerateMultiple**](docs/BarcodeApi.md#postgeneratemultiple) | **Post** /barcode/generateMultiple | Generate multiple barcodes and return in response stream             
*BarcodeApi* | [**PutBarcodeGenerateFile**](docs/BarcodeApi.md#putbarcodegeneratefile) | **Put** /barcode/{name}/generate | Generate barcode and save on server (from query params or from file with json or xml content)             
*BarcodeApi* | [**PutBarcodeRecognizeFromBody**](docs/BarcodeApi.md#putbarcoderecognizefrombody) | **Put** /barcode/{name}/recognize | Recognition of a barcode from file on server with parameters in body.             
*BarcodeApi* | [**PutGenerateMultiple**](docs/BarcodeApi.md#putgeneratemultiple) | **Put** /barcode/{name}/generateMultiple | Generate image with multiple barcodes and put new file on server             
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

 - [AustralianPostParams](docs/AustralianPostParams.md)
 - [AutoSizeMode](docs/AutoSizeMode.md)
 - [AvailableGraphicsUnit](docs/AvailableGraphicsUnit.md)
 - [AztecParams](docs/AztecParams.md)
 - [AztecSymbolMode](docs/AztecSymbolMode.md)
 - [BarCodeErrorResponse](docs/BarCodeErrorResponse.md)
 - [BarcodeResponse](docs/BarcodeResponse.md)
 - [BarcodeResponseList](docs/BarcodeResponseList.md)
 - [BorderDashStyle](docs/BorderDashStyle.md)
 - [CaptionParams](docs/CaptionParams.md)
 - [ChecksumValidation](docs/ChecksumValidation.md)
 - [CodabarChecksumMode](docs/CodabarChecksumMode.md)
 - [CodabarParams](docs/CodabarParams.md)
 - [CodabarSymbol](docs/CodabarSymbol.md)
 - [CodablockParams](docs/CodablockParams.md)
 - [Code16KParams](docs/Code16KParams.md)
 - [CodeLocation](docs/CodeLocation.md)
 - [CouponParams](docs/CouponParams.md)
 - [CustomerInformationInterpretingType](docs/CustomerInformationInterpretingType.md)
 - [DataBarParams](docs/DataBarParams.md)
 - [DataMatrixEccType](docs/DataMatrixEccType.md)
 - [DataMatrixEncodeMode](docs/DataMatrixEncodeMode.md)
 - [DataMatrixParams](docs/DataMatrixParams.md)
 - [DecodeBarcodeType](docs/DecodeBarcodeType.md)
 - [DiscUsage](docs/DiscUsage.md)
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
 - [Itf14BorderType](docs/Itf14BorderType.md)
 - [ItfParams](docs/ItfParams.md)
 - [MaxiCodeParams](docs/MaxiCodeParams.md)
 - [ModelError](docs/ModelError.md)
 - [ObjectExist](docs/ObjectExist.md)
 - [Padding](docs/Padding.md)
 - [Pdf417CompactionMode](docs/Pdf417CompactionMode.md)
 - [Pdf417ErrorLevel](docs/Pdf417ErrorLevel.md)
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
 - [TextAlignment](docs/TextAlignment.md)
 - [FileVersion](docs/FileVersion.md)


## Documentation For Authorization

## JWT
- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Author


