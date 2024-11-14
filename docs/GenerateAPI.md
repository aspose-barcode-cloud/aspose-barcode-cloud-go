# barcode\GenerateAPI

All URIs are relative to *<https://barcode.qa.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**BarcodeGenerateBarcodeTypeGet**](GenerateAPI.md#BarcodeGenerateBarcodeTypeGet) | **Get** /barcode/generate/{barcodeType} | Generate barcode using GET request with parameters in route and query string.
[**BarcodeGenerateBodyPost**](GenerateAPI.md#BarcodeGenerateBodyPost) | **Post** /barcode/generate-body | Generate barcode using POST request with parameters in body in json or xml format.
[**BarcodeGenerateMultipartPost**](GenerateAPI.md#BarcodeGenerateMultipartPost) | **Post** /barcode/generate-multipart | Generate barcode using POST request with parameters in multipart form.

## BarcodeGenerateBarcodeTypeGet

> *os.File BarcodeGenerateBarcodeTypeGet(ctx, barcodeType, data, optional)
Generate barcode using GET request with parameters in route and query string.

### BarcodeGenerateBarcodeTypeGet Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**EncodeBarcodeType**](.md) | Type of barcode to generate. |
 **data** | **string** | String represents data to encode |
 **optional** | ***GenerateAPIBarcodeGenerateBarcodeTypeGetOpts** | optional parameters | nil if no parameters

### BarcodeGenerateBarcodeTypeGet Optional Parameters

Optional parameters are passed through a pointer to a GenerateAPIBarcodeGenerateBarcodeTypeGetOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**DataType** | [**optional.Interface of EncodeDataType**](.md) | Type of data to encode.  Default value:  EncodeDataType.StringData. |
**ImageFormat** | [**optional.Interface of BarcodeImageFormat**](.md) | Barcode output image format.  Default value: png |
**TextLocation** | [**optional.Interface of CodeLocation**](.md) | Specify the displaying Text Location, set to CodeLocation.None to hide CodeText.  Default value: CodeLocation.Below. |
**ForegroundColor** | **optional.** | Specify the displaying bars and content Color.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: Black. |
**BackgroundColor** | **optional.** | Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: White. |
**Units** | [**optional.Interface of GraphicsUnit**](.md) | Common Units for all measuring in query. Default units: pixel. |
**Resolution** | **optional.** | Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.  Decimal separator is dot. |
**ImageHeight** | **optional.** | Height of the barcode image in given units. Default units: pixel.  Decimal separator is dot. |
**ImageWidth** | **optional.** | Width of the barcode image in given units. Default units: pixel.  Decimal separator is dot. |
**RotationAngle** | **optional.** | BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0. |

### BarcodeGenerateBarcodeTypeGet Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## BarcodeGenerateBodyPost

> *os.File BarcodeGenerateBodyPost(ctx, generateParams)
Generate barcode using POST request with parameters in body in json or xml format.

### BarcodeGenerateBodyPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **generateParams** | [**GenerateParams**](GenerateParams.md) | Parameters of generation |

### BarcodeGenerateBodyPost Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## BarcodeGenerateMultipartPost

> *os.File BarcodeGenerateMultipartPost(ctx, barcodeType, data, optional)
Generate barcode using POST request with parameters in multipart form.

### BarcodeGenerateMultipartPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**EncodeBarcodeType**](EncodeBarcodeType.md) |  |
 **data** | **string** | String represents data to encode |
 **optional** | ***GenerateAPIBarcodeGenerateMultipartPostOpts** | optional parameters | nil if no parameters

### BarcodeGenerateMultipartPost Optional Parameters

Optional parameters are passed through a pointer to a GenerateAPIBarcodeGenerateMultipartPostOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**DataType** | [**optional.Interface of EncodeDataType**](EncodeDataType.md) |  |
**ImageFormat** | [**optional.Interface of BarcodeImageFormat**](BarcodeImageFormat.md) |  |
**TextLocation** | [**optional.Interface of CodeLocation**](CodeLocation.md) |  |
**ForegroundColor** | **optional.** | Specify the displaying bars and content Color.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: Black. |
**BackgroundColor** | **optional.** | Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: White. |
**Units** | [**optional.Interface of GraphicsUnit**](GraphicsUnit.md) |  |
**Resolution** | **optional.** | Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.  Decimal separator is dot. |
**ImageHeight** | **optional.** | Height of the barcode image in given units. Default units: pixel.  Decimal separator is dot. |
**ImageWidth** | **optional.** | Width of the barcode image in given units. Default units: pixel.  Decimal separator is dot. |
**RotationAngle** | **optional.** | BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0. |

### BarcodeGenerateMultipartPost Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
