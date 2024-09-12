# barcode\GenerateAPI

All URIs are relative to *<https://barcode.qa.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**BarcodeGenerateBarcodeTypeGet**](GenerateAPI.md#BarcodeGenerateBarcodeTypeGet) | **Get** /barcode/generate/{barcodeType} | Generate barcode using GET request with parameters in route and query string.
[**BarcodeGenerateBodyPost**](GenerateAPI.md#BarcodeGenerateBodyPost) | **Post** /barcode/generate-body | Generate barcode using POST request with parameters in body in json or xml format.
[**BarcodeGenerateFormPost**](GenerateAPI.md#BarcodeGenerateFormPost) | **Post** /barcode/generate-form | Generate barcode using POST request with parameters in url ecncoded form.

## BarcodeGenerateBarcodeTypeGet

> *os.File BarcodeGenerateBarcodeTypeGet(ctx, barcodeType, dataType, data, optional)
Generate barcode using GET request with parameters in route and query string.

### BarcodeGenerateBarcodeTypeGet Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**EncodeBarcodeType**](.md) | Type of barcode to generate. |
 **dataType** | [**EncodeDataType**](.md) | Type of data to encode. |
 **data** | **string** | String represents data to encode |
 **optional** | ***GenerateAPIBarcodeGenerateBarcodeTypeGetOpts** | optional parameters | nil if no parameters

### BarcodeGenerateBarcodeTypeGet Optional Parameters

Optional parameters are passed through a pointer to a GenerateAPIBarcodeGenerateBarcodeTypeGetOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**ImageFormat** | [**optional.Interface of AvailableBarCodeImageFormat**](.md) | Barcode output image format.  Default value: png |
**TwoDDisplayText** | **optional.** | Text that will be displayed instead of codetext in 2D barcodes.  Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode |
**TextLocation** | [**optional.Interface of CodeLocation**](.md) | Specify the displaying Text Location, set to CodeLocation.None to hide CodeText.  Default value: CodeLocation.Below. |
**TextAlignment** | [**optional.Interface of TextAlignment**](.md) | Text alignment.  Default value: TextAligment.Left |
**ForegroundColor** | **optional.** | Specify the displaying bars and content Color.   Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.Black. |
**BackgroundColor** | **optional.** | Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.White. |
**Units** | [**optional.Interface of AvailableGraphicsUnit**](.md) | Common Units for all measuring in query. Default units: pixel. |
**Resolution** | **optional.** | Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi. |
**ImageHeight** | **optional.** | Height of the barcode image in given units. Default units: pixel. |
**ImageWidth** | **optional.** | Width of the barcode image in given units. Default units: pixel. |
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

## BarcodeGenerateFormPost

> *os.File BarcodeGenerateFormPost(ctx, barcodeType, dataType, data, optional)
Generate barcode using POST request with parameters in url ecncoded form.

### BarcodeGenerateFormPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**EncodeBarcodeType**](EncodeBarcodeType.md) |  |
 **dataType** | [**EncodeDataType**](EncodeDataType.md) |  |
 **data** | **string** | String represents data to encode |
 **optional** | ***GenerateAPIBarcodeGenerateFormPostOpts** | optional parameters | nil if no parameters

### BarcodeGenerateFormPost Optional Parameters

Optional parameters are passed through a pointer to a GenerateAPIBarcodeGenerateFormPostOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**ImageFormat** | [**optional.Interface of AvailableBarCodeImageFormat**](AvailableBarCodeImageFormat.md) |  |
**TwoDDisplayText** | **optional.** | Text that will be displayed instead of codetext in 2D barcodes.  Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode |
**TextLocation** | [**optional.Interface of CodeLocation**](CodeLocation.md) |  |
**TextAlignment** | [**optional.Interface of TextAlignment**](TextAlignment.md) |  |
**ForegroundColor** | **optional.** | Specify the displaying bars and content Color.   Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.Black. |
**BackgroundColor** | **optional.** | Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.White. |
**Units** | [**optional.Interface of AvailableGraphicsUnit**](AvailableGraphicsUnit.md) |  |
**Resolution** | **optional.** | Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi. |
**ImageHeight** | **optional.** | Height of the barcode image in given units. Default units: pixel. |
**ImageWidth** | **optional.** | Width of the barcode image in given units. Default units: pixel. |
**RotationAngle** | **optional.** | BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0. |

### BarcodeGenerateFormPost Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
