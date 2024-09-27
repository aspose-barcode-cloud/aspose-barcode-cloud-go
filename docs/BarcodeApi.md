# barcode\BarcodeApi

All URIs are relative to *<https://api.aspose.cloud/v3.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**GetBarcodeGenerate**](BarcodeApi.md#GetBarcodeGenerate) | **Get** /barcode/generate | Generate barcode.
[**GetBarcodeRecognize**](BarcodeApi.md#GetBarcodeRecognize) | **Get** /barcode/{name}/recognize | Recognize barcode from a file on server.
[**PostBarcodeRecognizeFromUrlOrContent**](BarcodeApi.md#PostBarcodeRecognizeFromUrlOrContent) | **Post** /barcode/recognize | Recognize barcode from an url or from request body. Request body can contain raw data bytes of the image with content-type \&quot;application/octet-stream\&quot;. An image can also be passed as a form field.
[**PostGenerateMultiple**](BarcodeApi.md#PostGenerateMultiple) | **Post** /barcode/generateMultiple | Generate multiple barcodes and return in response stream
[**PutBarcodeGenerateFile**](BarcodeApi.md#PutBarcodeGenerateFile) | **Put** /barcode/{name}/generate | Generate barcode and save on server (from query params or from file with json or xml content)
[**PutBarcodeRecognizeFromBody**](BarcodeApi.md#PutBarcodeRecognizeFromBody) | **Put** /barcode/{name}/recognize | Recognition of a barcode from file on server with parameters in body.
[**PutGenerateMultiple**](BarcodeApi.md#PutGenerateMultiple) | **Put** /barcode/{name}/generateMultiple | Generate image with multiple barcodes and put new file on server
[**ScanBarcode**](BarcodeApi.md#ScanBarcode) | **Post** /barcode/scan | Quickly scan a barcode from an image.

## GetBarcodeGenerate

> *os.File GetBarcodeGenerate(ctx, type_, text, optional)
Generate barcode.

### GetBarcodeGenerate Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **type_** | **string** | Type of barcode to generate. |
 **text** | **string** | Text to encode. |
 **optional** | ***BarcodeApiGetBarcodeGenerateOpts** | optional parameters | nil if no parameters

### GetBarcodeGenerate Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiGetBarcodeGenerateOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**TwoDDisplayText** | **optional.String** | Text that will be displayed instead of codetext in 2D barcodes. Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode |
**TextLocation** | **optional.String** | Specify the displaying Text Location, set to CodeLocation.None to hide CodeText. Default value: CodeLocation.Below. |
**TextAlignment** | **optional.String** | Text alignment. |
**TextColor** | **optional.String** | Specify the displaying CodeText&#39;s Color. Default value: black. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**NoWrap** | **optional.Bool** | Specify word wraps (line breaks) within text. Default value: false. |
**Resolution** | **optional.Float64** | Resolution of the BarCode image. One value for both dimensions. Default value: 96 dpi. |
**ResolutionX** | **optional.Float64** | DEPRECATED: Use &#39;Resolution&#39; instead. |
**ResolutionY** | **optional.Float64** | DEPRECATED: Use &#39;Resolution&#39; instead. |
**DimensionX** | **optional.Float64** | The smallest width of the unit of BarCode bars or spaces. Increase this will increase the whole barcode image width. Ignored if AutoSizeMode property is set to AutoSizeMode.Nearest or AutoSizeMode.Interpolation. |
**TextSpace** | **optional.Float64** | Space between the CodeText and the BarCode in Unit value. Default value: 2pt. Ignored for EAN8, EAN13, UPCE, UPCA, ISBN, ISMN, ISSN, UpcaGs1DatabarCoupon. |
**Units** | **optional.String** | Common Units for all measuring in query. Default units: pixel. |
**SizeMode** | **optional.String** | Specifies the different types of automatic sizing modes. Default value: AutoSizeMode.None. |
**BarHeight** | **optional.Float64** | Height of the barcode in given units. Default units: pixel. |
**ImageHeight** | **optional.Float64** | Height of the barcode image in given units. Default units: pixel. |
**ImageWidth** | **optional.Float64** | Width of the barcode image in given units. Default units: pixel. |
**RotationAngle** | **optional.Float64** | BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image. Default value: 0. |
**BackColor** | **optional.String** | Background color of the barcode image. Default value: white. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**BarColor** | **optional.String** | Bars color. Default value: black. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**BorderColor** | **optional.String** | Border color. Default value: black. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**BorderWidth** | **optional.Float64** | Border width. Default value: 0. Ignored if Visible is set to false. |
**BorderDashStyle** | **optional.String** | Border dash style. Default value: BorderDashStyle.Solid. |
**BorderVisible** | **optional.Bool** | Border visibility. If false than parameter Width is always ignored (0). Default value: false. |
**EnableChecksum** | **optional.String** | Enable checksum during generation 1D barcodes. Default is treated as Yes for symbology which must contain checksum, as No where checksum only possible. Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN, Codabar Checksum always used: Rest symbology |
**EnableEscape** | **optional.Bool** | Indicates whether explains the character \&quot;\\\&quot; as an escape character in CodeText property. Used for Pdf417, DataMatrix, Code128 only If the EnableEscape is true, \&quot;\\\&quot; will be explained as a special escape character. Otherwise, \&quot;\\\&quot; acts as normal characters. Aspose.BarCode supports input decimal ascii code and mnemonic for ASCII control-code characters. For example, \\013 and \\\\CR stands for CR. |
**FilledBars** | **optional.Bool** | Value indicating whether bars are filled. Only for 1D barcodes. Default value: true. |
**AlwaysShowChecksum** | **optional.Bool** | Always display checksum digit in the human readable text for Code128 and GS1Code128 barcodes. |
**WideNarrowRatio** | **optional.Float64** | Wide bars to Narrow bars ratio. Default value: 3, that is, wide bars are 3 times as wide as narrow bars. Used for ITF, PZN, PharmaCode, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, IATA2of5, VIN, DeutschePost, OPC, Code32, DataLogic2of5, PatchCode, Code39Extended, Code39Standard |
**ValidateText** | **optional.Bool** | Only for 1D barcodes. If codetext is incorrect and value set to true - exception will be thrown. Otherwise codetext will be corrected to match barcode&#39;s specification. Exception always will be thrown for: Databar symbology if codetext is incorrect. Exception always will not be thrown for: AustraliaPost, SingaporePost, Code39Extended, Code93Extended, Code16K, Code128 symbology if codetext is incorrect. |
**SupplementData** | **optional.String** | Supplement parameters. Used for Interleaved2of5, Standard2of5, EAN13, EAN8, UPCA, UPCE, ISBN, ISSN, ISMN. |
**SupplementSpace** | **optional.Float64** | Space between main the BarCode and supplement BarCode. |
**BarWidthReduction** | **optional.Float64** | Bars reduction value that is used to compensate ink spread while printing. |
**UseAntiAlias** | **optional.Bool** | Indicates whether is used anti-aliasing mode to render image. Anti-aliasing mode is applied to barcode and text drawing. |
**Format** | **optional.String** | Result image format. |

### GetBarcodeGenerate Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## GetBarcodeRecognize

> BarcodeResponseList GetBarcodeRecognize(ctx, name, optional)
Recognize barcode from a file on server.

### GetBarcodeRecognize Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **name** | **string** | The image file name. |
 **optional** | ***BarcodeApiGetBarcodeRecognizeOpts** | optional parameters | nil if no parameters

### GetBarcodeRecognize Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiGetBarcodeRecognizeOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**Type_** | **optional.String** | The type of barcode to read. |
**Types** | **optional.Interface of []DecodeBarcodeType** | Multiple barcode types to read. |
**ChecksumValidation** | **optional.String** | Enable checksum validation during recognition for 1D barcodes. Default is treated as Yes for symbologies which must contain checksum, as No where checksum only possible. Checksum never used: Codabar Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN Checksum always used: Rest symbologies |
**DetectEncoding** | **optional.Bool** | A flag which force engine to detect codetext encoding for Unicode. |
**Preset** | **optional.String** | Preset allows to configure recognition quality and speed manually. You can quickly set up Preset by embedded presets: HighPerformance, NormalQuality, HighQuality, MaxBarCodes or you can manually configure separate options. Default value of Preset is NormalQuality. |
**RectX** | **optional.Int32** | Set X of top left corner of area for recognition. |
**RectY** | **optional.Int32** | Set Y of top left corner of area for recognition. |
**RectWidth** | **optional.Int32** | Set Width of area for recognition. |
**RectHeight** | **optional.Int32** | Set Height of area for recognition. |
**StripFNC** | **optional.Bool** | Value indicating whether FNC symbol strip must be done. |
**Timeout** | **optional.Int32** | Timeout of recognition process in milliseconds. Default value is 15_000 (15 seconds). Maximum value is 30_000 (1/2 minute). In case of a timeout RequestTimeout (408) status will be returned. Try reducing the image size to avoid timeout. |
**MedianSmoothingWindowSize** | **optional.Int32** | Window size for median smoothing. Typical values are 3 or 4. Default value is 3. AllowMedianSmoothing must be set. |
**AllowMedianSmoothing** | **optional.Bool** | Allows engine to enable median smoothing as additional scan. Mode helps to recognize noised barcodes. |
**AllowComplexBackground** | **optional.Bool** | Allows engine to recognize color barcodes on color background as additional scan. Extremely slow mode. |
**AllowDatamatrixIndustrialBarcodes** | **optional.Bool** | Allows engine for Datamatrix to recognize dashed industrial Datamatrix barcodes. Slow mode which helps only for dashed barcodes which consist from spots. |
**AllowDecreasedImage** | **optional.Bool** | Allows engine to recognize decreased image as additional scan. Size for decreasing is selected by internal engine algorithms. Mode helps to recognize barcodes which are noised and blurred but captured with high resolution. |
**AllowDetectScanGap** | **optional.Bool** | Allows engine to use gap between scans to increase recognition speed. Mode can make recognition problems with low height barcodes. |
**AllowIncorrectBarcodes** | **optional.Bool** | Allows engine to recognize barcodes which has incorrect checksum or incorrect values. Mode can be used to recognize damaged barcodes with incorrect text. |
**AllowInvertImage** | **optional.Bool** | Allows engine to recognize inverse color image as additional scan. Mode can be used when barcode is white on black background. |
**AllowMicroWhiteSpotsRemoving** | **optional.Bool** | Allows engine for Postal barcodes to recognize slightly noised images. Mode helps to recognize slightly damaged Postal barcodes. |
**AllowOneDFastBarcodesDetector** | **optional.Bool** | Allows engine for 1D barcodes to quickly recognize high quality barcodes which fill almost whole image. Mode helps to quickly recognize generated barcodes from Internet. |
**AllowOneDWipedBarsRestoration** | **optional.Bool** | Allows engine for 1D barcodes to recognize barcodes with single wiped/glued bars in pattern. |
**AllowQRMicroQrRestoration** | **optional.Bool** | Allows engine for QR/MicroQR to recognize damaged MicroQR barcodes. |
**AllowRegularImage** | **optional.Bool** | Allows engine to recognize regular image without any restorations as main scan. Mode to recognize image as is. |
**AllowSaltAndPepperFiltering** | **optional.Bool** | Allows engine to recognize barcodes with salt and pepper noise type. Mode can remove small noise with white and black dots. |
**AllowWhiteSpotsRemoving** | **optional.Bool** | Allows engine to recognize image without small white spots as additional scan. Mode helps to recognize noised image as well as median smoothing filtering. |
**CheckMore1DVariants** | **optional.Bool** | Allows engine to recognize 1D barcodes with checksum by checking more recognition variants. Default value: False. |
**FastScanOnly** | **optional.Bool** | Allows engine for 1D barcodes to quickly recognize middle slice of an image and return result without using any time-consuming algorithms. Default value: False. |
**AllowAdditionalRestorations** | **optional.Bool** | Allows engine using additional image restorations to recognize corrupted barcodes. At this time, it is used only in MicroPdf417 barcode type. Default value: False. |
**RegionLikelihoodThresholdPercent** | **optional.Float64** | Sets threshold for detected regions that may contain barcodes. Value 0.7 means that bottom 70% of possible regions are filtered out and not processed further. Region likelihood threshold must be between [0.05, 0.9] Use high values for clear images with few barcodes. Use low values for images with many barcodes or for noisy images. Low value may lead to a bigger recognition time. |
**ScanWindowSizes** | **optional.Interface of []int32** | Scan window sizes in pixels. Allowed sizes are 10, 15, 20, 25, 30. Scanning with small window size takes more time and provides more accuracy but may fail in detecting very big barcodes. Combining of several window sizes can improve detection quality. |
**Similarity** | **optional.Float64** | Similarity coefficient depends on how homogeneous barcodes are. Use high value for clear barcodes. Use low values to detect barcodes that ara partly damaged or not lighten evenly. Similarity coefficient must be between [0.5, 0.9] |
**SkipDiagonalSearch** | **optional.Bool** | Allows detector to skip search for diagonal barcodes. Setting it to false will increase detection time but allow to find diagonal barcodes that can be missed otherwise. Enabling of diagonal search leads to a bigger detection time. |
**ReadTinyBarcodes** | **optional.Bool** | Allows engine to recognize tiny barcodes on large images. Ignored if AllowIncorrectBarcodes is set to True. Default value: False. |
**AustralianPostEncodingTable** | **optional.String** | Interpreting Type for the Customer Information of AustralianPost BarCode.Default is CustomerInformationInterpretingType.Other. |
**IgnoreEndingFillingPatternsForCTable** | **optional.Bool** | The flag which force AustraliaPost decoder to ignore last filling patterns in Customer Information Field during decoding as CTable method. CTable encoding method does not have any gaps in encoding table and sequence \&quot;333\&quot; of filling patterns is decoded as letter \&quot;z\&quot;. |
**Storage** | **optional.String** | The image storage. |
**Folder** | **optional.String** | The image folder. |

### GetBarcodeRecognize Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## PostBarcodeRecognizeFromUrlOrContent

> BarcodeResponseList PostBarcodeRecognizeFromUrlOrContent(ctx, optional)
Recognize barcode from an url or from request body. Request body can contain raw data bytes of the image with content-type \"application/octet-stream\". An image can also be passed as a form field.

### PostBarcodeRecognizeFromUrlOrContent Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts** | optional parameters | nil if no parameters

### PostBarcodeRecognizeFromUrlOrContent Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**Type_** | **optional.String** | The type of barcode to read. |
**Types** | **optional.Interface of []DecodeBarcodeType** | Multiple barcode types to read. |
**ChecksumValidation** | **optional.String** | Enable checksum validation during recognition for 1D barcodes. Default is treated as Yes for symbologies which must contain checksum, as No where checksum only possible. Checksum never used: Codabar Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN Checksum always used: Rest symbologies |
**DetectEncoding** | **optional.Bool** | A flag which force engine to detect codetext encoding for Unicode. |
**Preset** | **optional.String** | Preset allows to configure recognition quality and speed manually. You can quickly set up Preset by embedded presets: HighPerformance, NormalQuality, HighQuality, MaxBarCodes or you can manually configure separate options. Default value of Preset is NormalQuality. |
**RectX** | **optional.Int32** | Set X of top left corner of area for recognition. |
**RectY** | **optional.Int32** | Set Y of top left corner of area for recognition. |
**RectWidth** | **optional.Int32** | Set Width of area for recognition. |
**RectHeight** | **optional.Int32** | Set Height of area for recognition. |
**StripFNC** | **optional.Bool** | Value indicating whether FNC symbol strip must be done. |
**Timeout** | **optional.Int32** | Timeout of recognition process in milliseconds. Default value is 15_000 (15 seconds). Maximum value is 30_000 (1/2 minute). In case of a timeout RequestTimeout (408) status will be returned. Try reducing the image size to avoid timeout. |
**MedianSmoothingWindowSize** | **optional.Int32** | Window size for median smoothing. Typical values are 3 or 4. Default value is 3. AllowMedianSmoothing must be set. |
**AllowMedianSmoothing** | **optional.Bool** | Allows engine to enable median smoothing as additional scan. Mode helps to recognize noised barcodes. |
**AllowComplexBackground** | **optional.Bool** | Allows engine to recognize color barcodes on color background as additional scan. Extremely slow mode. |
**AllowDatamatrixIndustrialBarcodes** | **optional.Bool** | Allows engine for Datamatrix to recognize dashed industrial Datamatrix barcodes. Slow mode which helps only for dashed barcodes which consist from spots. |
**AllowDecreasedImage** | **optional.Bool** | Allows engine to recognize decreased image as additional scan. Size for decreasing is selected by internal engine algorithms. Mode helps to recognize barcodes which are noised and blurred but captured with high resolution. |
**AllowDetectScanGap** | **optional.Bool** | Allows engine to use gap between scans to increase recognition speed. Mode can make recognition problems with low height barcodes. |
**AllowIncorrectBarcodes** | **optional.Bool** | Allows engine to recognize barcodes which has incorrect checksum or incorrect values. Mode can be used to recognize damaged barcodes with incorrect text. |
**AllowInvertImage** | **optional.Bool** | Allows engine to recognize inverse color image as additional scan. Mode can be used when barcode is white on black background. |
**AllowMicroWhiteSpotsRemoving** | **optional.Bool** | Allows engine for Postal barcodes to recognize slightly noised images. Mode helps to recognize slightly damaged Postal barcodes. |
**AllowOneDFastBarcodesDetector** | **optional.Bool** | Allows engine for 1D barcodes to quickly recognize high quality barcodes which fill almost whole image. Mode helps to quickly recognize generated barcodes from Internet. |
**AllowOneDWipedBarsRestoration** | **optional.Bool** | Allows engine for 1D barcodes to recognize barcodes with single wiped/glued bars in pattern. |
**AllowQRMicroQrRestoration** | **optional.Bool** | Allows engine for QR/MicroQR to recognize damaged MicroQR barcodes. |
**AllowRegularImage** | **optional.Bool** | Allows engine to recognize regular image without any restorations as main scan. Mode to recognize image as is. |
**AllowSaltAndPepperFiltering** | **optional.Bool** | Allows engine to recognize barcodes with salt and pepper noise type. Mode can remove small noise with white and black dots. |
**AllowWhiteSpotsRemoving** | **optional.Bool** | Allows engine to recognize image without small white spots as additional scan. Mode helps to recognize noised image as well as median smoothing filtering. |
**CheckMore1DVariants** | **optional.Bool** | Allows engine to recognize 1D barcodes with checksum by checking more recognition variants. Default value: False. |
**FastScanOnly** | **optional.Bool** | Allows engine for 1D barcodes to quickly recognize middle slice of an image and return result without using any time-consuming algorithms. Default value: False. |
**AllowAdditionalRestorations** | **optional.Bool** | Allows engine using additional image restorations to recognize corrupted barcodes. At this time, it is used only in MicroPdf417 barcode type. Default value: False. |
**RegionLikelihoodThresholdPercent** | **optional.Float64** | Sets threshold for detected regions that may contain barcodes. Value 0.7 means that bottom 70% of possible regions are filtered out and not processed further. Region likelihood threshold must be between [0.05, 0.9] Use high values for clear images with few barcodes. Use low values for images with many barcodes or for noisy images. Low value may lead to a bigger recognition time. |
**ScanWindowSizes** | **optional.Interface of []int32** | Scan window sizes in pixels. Allowed sizes are 10, 15, 20, 25, 30. Scanning with small window size takes more time and provides more accuracy but may fail in detecting very big barcodes. Combining of several window sizes can improve detection quality. |
**Similarity** | **optional.Float64** | Similarity coefficient depends on how homogeneous barcodes are. Use high value for clear barcodes. Use low values to detect barcodes that ara partly damaged or not lighten evenly. Similarity coefficient must be between [0.5, 0.9] |
**SkipDiagonalSearch** | **optional.Bool** | Allows detector to skip search for diagonal barcodes. Setting it to false will increase detection time but allow to find diagonal barcodes that can be missed otherwise. Enabling of diagonal search leads to a bigger detection time. |
**ReadTinyBarcodes** | **optional.Bool** | Allows engine to recognize tiny barcodes on large images. Ignored if AllowIncorrectBarcodes is set to True. Default value: False. |
**AustralianPostEncodingTable** | **optional.String** | Interpreting Type for the Customer Information of AustralianPost BarCode.Default is CustomerInformationInterpretingType.Other. |
**IgnoreEndingFillingPatternsForCTable** | **optional.Bool** | The flag which force AustraliaPost decoder to ignore last filling patterns in Customer Information Field during decoding as CTable method. CTable encoding method does not have any gaps in encoding table and sequence \&quot;333\&quot; of filling patterns is decoded as letter \&quot;z\&quot;. |
**Url** | **optional.String** | The image file url. |
**Image** | **optional.Interface of *os.File** | Image data |

### PostBarcodeRecognizeFromUrlOrContent Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## PostGenerateMultiple

> *os.File PostGenerateMultiple(ctx, generatorParamsList, optional)
Generate multiple barcodes and return in response stream

### PostGenerateMultiple Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **generatorParamsList** | [**GeneratorParamsList**](GeneratorParamsList.md) | List of barcodes |
 **optional** | ***BarcodeApiPostGenerateMultipleOpts** | optional parameters | nil if no parameters

### PostGenerateMultiple Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiPostGenerateMultipleOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**Format** | **optional.String** | Format to return stream in | [default to png]

### PostGenerateMultiple Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## PutBarcodeGenerateFile

> ResultImageInfo PutBarcodeGenerateFile(ctx, name, type_, text, optional)
Generate barcode and save on server (from query params or from file with json or xml content)

### PutBarcodeGenerateFile Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **name** | **string** | The image file name. |
 **type_** | **string** | Type of barcode to generate. |
 **text** | **string** | Text to encode. |
 **optional** | ***BarcodeApiPutBarcodeGenerateFileOpts** | optional parameters | nil if no parameters

### PutBarcodeGenerateFile Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiPutBarcodeGenerateFileOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**TwoDDisplayText** | **optional.String** | Text that will be displayed instead of codetext in 2D barcodes. Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode |
**TextLocation** | **optional.String** | Specify the displaying Text Location, set to CodeLocation.None to hide CodeText. Default value: CodeLocation.Below. |
**TextAlignment** | **optional.String** | Text alignment. |
**TextColor** | **optional.String** | Specify the displaying CodeText&#39;s Color. Default value: black. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**NoWrap** | **optional.Bool** | Specify word wraps (line breaks) within text. Default value: false. |
**Resolution** | **optional.Float64** | Resolution of the BarCode image. One value for both dimensions. Default value: 96 dpi. |
**ResolutionX** | **optional.Float64** | DEPRECATED: Use &#39;Resolution&#39; instead. |
**ResolutionY** | **optional.Float64** | DEPRECATED: Use &#39;Resolution&#39; instead. |
**DimensionX** | **optional.Float64** | The smallest width of the unit of BarCode bars or spaces. Increase this will increase the whole barcode image width. Ignored if AutoSizeMode property is set to AutoSizeMode.Nearest or AutoSizeMode.Interpolation. |
**TextSpace** | **optional.Float64** | Space between the CodeText and the BarCode in Unit value. Default value: 2pt. Ignored for EAN8, EAN13, UPCE, UPCA, ISBN, ISMN, ISSN, UpcaGs1DatabarCoupon. |
**Units** | **optional.String** | Common Units for all measuring in query. Default units: pixel. |
**SizeMode** | **optional.String** | Specifies the different types of automatic sizing modes. Default value: AutoSizeMode.None. |
**BarHeight** | **optional.Float64** | Height of the barcode in given units. Default units: pixel. |
**ImageHeight** | **optional.Float64** | Height of the barcode image in given units. Default units: pixel. |
**ImageWidth** | **optional.Float64** | Width of the barcode image in given units. Default units: pixel. |
**RotationAngle** | **optional.Float64** | BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image. Default value: 0. |
**BackColor** | **optional.String** | Background color of the barcode image. Default value: white. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**BarColor** | **optional.String** | Bars color. Default value: black. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**BorderColor** | **optional.String** | Border color. Default value: black. Use named colors like: red, green, blue Or HTML colors like: #FF0000, #00FF00, #0000FF |
**BorderWidth** | **optional.Float64** | Border width. Default value: 0. Ignored if Visible is set to false. |
**BorderDashStyle** | **optional.String** | Border dash style. Default value: BorderDashStyle.Solid. |
**BorderVisible** | **optional.Bool** | Border visibility. If false than parameter Width is always ignored (0). Default value: false. |
**EnableChecksum** | **optional.String** | Enable checksum during generation 1D barcodes. Default is treated as Yes for symbology which must contain checksum, as No where checksum only possible. Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN, Codabar Checksum always used: Rest symbology |
**EnableEscape** | **optional.Bool** | Indicates whether explains the character \&quot;\\\&quot; as an escape character in CodeText property. Used for Pdf417, DataMatrix, Code128 only If the EnableEscape is true, \&quot;\\\&quot; will be explained as a special escape character. Otherwise, \&quot;\\\&quot; acts as normal characters. Aspose.BarCode supports input decimal ascii code and mnemonic for ASCII control-code characters. For example, \\013 and \\\\CR stands for CR. |
**FilledBars** | **optional.Bool** | Value indicating whether bars are filled. Only for 1D barcodes. Default value: true. |
**AlwaysShowChecksum** | **optional.Bool** | Always display checksum digit in the human readable text for Code128 and GS1Code128 barcodes. |
**WideNarrowRatio** | **optional.Float64** | Wide bars to Narrow bars ratio. Default value: 3, that is, wide bars are 3 times as wide as narrow bars. Used for ITF, PZN, PharmaCode, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, IATA2of5, VIN, DeutschePost, OPC, Code32, DataLogic2of5, PatchCode, Code39Extended, Code39Standard |
**ValidateText** | **optional.Bool** | Only for 1D barcodes. If codetext is incorrect and value set to true - exception will be thrown. Otherwise codetext will be corrected to match barcode&#39;s specification. Exception always will be thrown for: Databar symbology if codetext is incorrect. Exception always will not be thrown for: AustraliaPost, SingaporePost, Code39Extended, Code93Extended, Code16K, Code128 symbology if codetext is incorrect. |
**SupplementData** | **optional.String** | Supplement parameters. Used for Interleaved2of5, Standard2of5, EAN13, EAN8, UPCA, UPCE, ISBN, ISSN, ISMN. |
**SupplementSpace** | **optional.Float64** | Space between main the BarCode and supplement BarCode. |
**BarWidthReduction** | **optional.Float64** | Bars reduction value that is used to compensate ink spread while printing. |
**UseAntiAlias** | **optional.Bool** | Indicates whether is used anti-aliasing mode to render image. Anti-aliasing mode is applied to barcode and text drawing. |
**Storage** | **optional.String** | Image&#39;s storage. |
**Folder** | **optional.String** | Image&#39;s folder. |
**Format** | **optional.String** | The image format. |

### PutBarcodeGenerateFile Return type

[**ResultImageInfo**](ResultImageInfo.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## PutBarcodeRecognizeFromBody

> BarcodeResponseList PutBarcodeRecognizeFromBody(ctx, name, readerParams, optional)
Recognition of a barcode from file on server with parameters in body.

### PutBarcodeRecognizeFromBody Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **name** | **string** | The image file name. |
 **readerParams** | [**ReaderParams**](ReaderParams.md) | BarcodeReader object with parameters. |
 **optional** | ***BarcodeApiPutBarcodeRecognizeFromBodyOpts** | optional parameters | nil if no parameters

### PutBarcodeRecognizeFromBody Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiPutBarcodeRecognizeFromBodyOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**Type_** | **optional.String** |  |
**Storage** | **optional.String** | The storage name |
**Folder** | **optional.String** | The image folder. |

### PutBarcodeRecognizeFromBody Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## PutGenerateMultiple

> ResultImageInfo PutGenerateMultiple(ctx, name, generatorParamsList, optional)
Generate image with multiple barcodes and put new file on server

### PutGenerateMultiple Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **name** | **string** | New filename |
 **generatorParamsList** | [**GeneratorParamsList**](GeneratorParamsList.md) | List of barcodes |
 **optional** | ***BarcodeApiPutGenerateMultipleOpts** | optional parameters | nil if no parameters

### PutGenerateMultiple Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiPutGenerateMultipleOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**Format** | **optional.String** | Format of file | [default to png]
**Folder** | **optional.String** | Folder to place file to |
**Storage** | **optional.String** | The storage name |

### PutGenerateMultiple Return type

[**ResultImageInfo**](ResultImageInfo.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## ScanBarcode

> BarcodeResponseList ScanBarcode(ctx, imageFile, optional)
Quickly scan a barcode from an image.

### ScanBarcode Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **imageFile** | ***os.File** | Image as file |
 **optional** | ***BarcodeApiScanBarcodeOpts** | optional parameters | nil if no parameters

### ScanBarcode Optional Parameters

Optional parameters are passed through a pointer to a BarcodeApiScanBarcodeOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**DecodeTypes** | **optional.Interface of []DecodeBarcodeType** | Types of barcode to recognize |
**Timeout** | **optional.Int32** | Timeout of recognition process in milliseconds.  Default value is 15_000 (15 seconds).  Maximum value is 30_000 (1/2 minute).  In case of a timeout RequestTimeout (408) status will be returned.  Try reducing the image size to avoid timeout. |
**ChecksumValidation** | **optional.String** | Checksum validation setting. Default is ON. |

### ScanBarcode Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
