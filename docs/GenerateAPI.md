# barcode\GenerateAPI

All URIs are relative to *<https://api.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**Generate**](GenerateAPI.md#Generate) | **Get** /barcode/generate/{barcodeType} | Generate a barcode using a GET request with parameters in the route and query string.
[**GenerateBody**](GenerateAPI.md#GenerateBody) | **Post** /barcode/generate-body | Generate a barcode using a POST request with parameters in the request body in JSON or XML format.
[**GenerateMultipart**](GenerateAPI.md#GenerateMultipart) | **Post** /barcode/generate-multipart | Generate a barcode using a POST request with parameters in a multipart form.

## Generate

> *os.File Generate(ctx, barcodeType, data, optional)
Generate a barcode using a GET request with parameters in the route and query string.

### Generate Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**EncodeBarcodeType**](.md) | Type of barcode to generate. |
 **data** | **string** | String that represents the data to encode. |
 **optional** | ***GenerateAPIGenerateOpts** | optional parameters | nil if no parameters

### Generate Optional Parameters

Optional parameters are passed through a pointer to a GenerateAPIGenerateOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**DataType** | [**optional.Interface of EncodeDataType**](.md) | Type of data to encode. Default value: StringData. | [default to &quot;StringData&quot;]
**ImageFormat** | [**optional.Interface of BarcodeImageFormat**](.md) | Barcode output image format. Default value: png. | [default to &quot;Png&quot;]
**TextLocation** | [**optional.Interface of CodeLocation**](.md) | Specify the displayed text location. Set to CodeLocation.None to hide CodeText. Default value depends on BarcodeType: CodeLocation.Below for 1D barcodes and CodeLocation.None for 2D barcodes. |
**ForegroundColor** | **optional.** | Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black. | [default to &quot;Black&quot;]
**BackgroundColor** | **optional.** | Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White. | [default to &quot;White&quot;]
**Units** | [**optional.Interface of GraphicsUnit**](.md) | Common units for all measurements. Default units: pixels. |
**Resolution** | **optional.** | Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot. |
**ImageHeight** | **optional.** | Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot. |
**ImageWidth** | **optional.** | Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot. |
**RotationAngle** | **optional.** | Barcode image rotation angle, measured in degrees. For example, RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0. |
**QrEncodeMode** | [**optional.Interface of QREncodeMode**](.md) | QR barcode encode mode. |
**QrErrorLevel** | [**optional.Interface of QRErrorLevel**](.md) | QR barcode error correction level. |
**QrVersion** | [**optional.Interface of QRVersion**](.md) | QR barcode version. Automatically selects the smallest version that fits the data. |
**QrECIEncoding** | [**optional.Interface of ECIEncodings**](.md) | ECI encoding for QR barcode data. |
**QrAspectRatio** | **optional.** | QR barcode aspect ratio. Values: 0 to 1. |
**MicroQRVersion** | [**optional.Interface of MicroQRVersion**](.md) | MicroQR barcode version. Used when BarcodeType is MicroQR. |
**RectMicroQrVersion** | [**optional.Interface of RectMicroQRVersion**](.md) | RectMicroQR barcode version. Used when BarcodeType is RectMicroQR. |
**Code128EncodeMode** | [**optional.Interface of Code128EncodeMode**](.md) | Code128 barcode encode mode. Controls which Code 128 subset (A, B, C, or mix) is used. |
**Pdf417EncodeMode** | [**optional.Interface of Pdf417EncodeMode**](.md) | PDF417 barcode encode mode. |
**Pdf417ErrorLevel** | [**optional.Interface of Pdf417ErrorLevel**](.md) | PDF417 barcode error correction level. |
**Pdf417Truncate** | **optional.** | Whether to use truncated PDF417 format (removes right-side stop pattern). |
**Pdf417Columns** | **optional.** | Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto. |
**Pdf417Rows** | **optional.** | Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic. |
**Pdf417AspectRatio** | **optional.** | PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417. |
**Pdf417ECIEncoding** | [**optional.Interface of ECIEncodings**](.md) | ECI encoding for PDF417 barcode data. |
**Pdf417IsReaderInitialization** | **optional.** | Whether the barcode is used for reader initialization (programming). |
**Pdf417MacroCharacters** | [**optional.Interface of MacroCharacter**](.md) | Macro character to prepend (structured append). |
**Pdf417IsLinked** | **optional.** | Whether to use linked mode (for MicroPdf417). |
**Pdf417IsCode128Emulation** | **optional.** | Whether to use Code128 emulation for MicroPdf417. |

### Generate Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## GenerateBody

> *os.File GenerateBody(ctx, generateParams)
Generate a barcode using a POST request with parameters in the request body in JSON or XML format.

### GenerateBody Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **generateParams** | [**GenerateParams**](GenerateParams.md) | Generation parameters. |

### GenerateBody Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## GenerateMultipart

> *os.File GenerateMultipart(ctx, barcodeType, data, optional)
Generate a barcode using a POST request with parameters in a multipart form.

### GenerateMultipart Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**EncodeBarcodeType**](EncodeBarcodeType.md) | See https://reference.aspose.com/barcode/net/aspose.barcode.generation/encodetypes/ |
 **data** | **string** | String that represents the data to encode. |
 **optional** | ***GenerateAPIGenerateMultipartOpts** | optional parameters | nil if no parameters

### GenerateMultipart Optional Parameters

Optional parameters are passed through a pointer to a GenerateAPIGenerateMultipartOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**DataType** | [**optional.Interface of EncodeDataType**](EncodeDataType.md) | Type of data to encode. Default value: StringData. | [default to &quot;StringData&quot;]
**ImageFormat** | [**optional.Interface of BarcodeImageFormat**](BarcodeImageFormat.md) | Barcode output image format. Default value: png. | [default to &quot;Png&quot;]
**TextLocation** | [**optional.Interface of CodeLocation**](CodeLocation.md) | Specify the displayed text location. Set to CodeLocation.None to hide CodeText. Default value depends on BarcodeType: CodeLocation.Below for 1D barcodes and CodeLocation.None for 2D barcodes. |
**ForegroundColor** | **optional.** | Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black. | [default to &quot;Black&quot;]
**BackgroundColor** | **optional.** | Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White. | [default to &quot;White&quot;]
**Units** | [**optional.Interface of GraphicsUnit**](GraphicsUnit.md) | Common units for all measurements. Default units: pixels. |
**Resolution** | **optional.** | Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot. |
**ImageHeight** | **optional.** | Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot. |
**ImageWidth** | **optional.** | Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot. |
**RotationAngle** | **optional.** | Barcode image rotation angle, measured in degrees. For example, RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0. |
**QrEncodeMode** | [**optional.Interface of QREncodeMode**](QREncodeMode.md) | QR barcode encode mode. |
**QrErrorLevel** | [**optional.Interface of QRErrorLevel**](QRErrorLevel.md) | QR barcode error correction level. |
**QrVersion** | [**optional.Interface of QRVersion**](QRVersion.md) | QR barcode version. Automatically selects the smallest version that fits the data. |
**QrECIEncoding** | [**optional.Interface of ECIEncodings**](ECIEncodings.md) | ECI encoding for QR barcode data. |
**QrAspectRatio** | **optional.** | QR barcode aspect ratio. Values: 0 to 1. |
**MicroQRVersion** | [**optional.Interface of MicroQRVersion**](MicroQRVersion.md) | MicroQR barcode version. Used when BarcodeType is MicroQR. |
**RectMicroQrVersion** | [**optional.Interface of RectMicroQRVersion**](RectMicroQRVersion.md) | RectMicroQR barcode version. Used when BarcodeType is RectMicroQR. |
**Code128EncodeMode** | [**optional.Interface of Code128EncodeMode**](Code128EncodeMode.md) | Code128 barcode encode mode. Controls which Code 128 subset (A, B, C, or mix) is used. |
**Pdf417EncodeMode** | [**optional.Interface of Pdf417EncodeMode**](Pdf417EncodeMode.md) | PDF417 barcode encode mode. |
**Pdf417ErrorLevel** | [**optional.Interface of Pdf417ErrorLevel**](Pdf417ErrorLevel.md) | PDF417 barcode error correction level. |
**Pdf417Truncate** | **optional.** | Whether to use truncated PDF417 format (removes right-side stop pattern). |
**Pdf417Columns** | **optional.** | Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto. |
**Pdf417Rows** | **optional.** | Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic. |
**Pdf417AspectRatio** | **optional.** | PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417. |
**Pdf417ECIEncoding** | [**optional.Interface of ECIEncodings**](ECIEncodings.md) | ECI encoding for PDF417 barcode data. |
**Pdf417IsReaderInitialization** | **optional.** | Whether the barcode is used for reader initialization (programming). |
**Pdf417MacroCharacters** | [**optional.Interface of MacroCharacter**](MacroCharacter.md) | Macro character to prepend (structured append). |
**Pdf417IsLinked** | **optional.** | Whether to use linked mode (for MicroPdf417). |
**Pdf417IsCode128Emulation** | **optional.** | Whether to use Code128 emulation for MicroPdf417. |

### GenerateMultipart Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
