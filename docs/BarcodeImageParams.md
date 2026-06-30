# BarcodeImageParams

Optional barcode image parameters.

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ImageFormat** | [***BarcodeImageFormat**](BarcodeImageFormat.md) | Barcode output image format. Default value: png. | [optional] [default to PNG]
**TextLocation** | [***CodeLocation**](CodeLocation.md) | Specify the displayed text location. Set to CodeLocation.None to hide CodeText. Default value depends on BarcodeType: CodeLocation.Below for 1D barcodes and CodeLocation.None for 2D barcodes. | [optional]
**ForegroundColor** | **NullableString** | Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black. | [optional] [default to "Black"]
**BackgroundColor** | **NullableString** | Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White. | [optional] [default to "White"]
**Units** | [***GraphicsUnit**](GraphicsUnit.md) | Common units for all measurements. Default units: pixels. | [optional]
**Resolution** | **NullableFloat32** | Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot. | [optional]
**ImageHeight** | **NullableFloat32** | Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot. | [optional]
**ImageWidth** | **NullableFloat32** | Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot. | [optional]
**RotationAngle** | **NullableInt32** | Barcode image rotation angle, measured in degrees. For example, RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
