# BarcodeImageParams

Barcode image optional parameters

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ImageFormat** | [***BarcodeImageFormat**](BarcodeImageFormat.md) |  | [optional]
**TextLocation** | [***CodeLocation**](CodeLocation.md) |  | [optional]
**ForegroundColor** | **NullableString** | Specify the displaying bars and content Color.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: Black. | [optional] [default to "Black"]
**BackgroundColor** | **NullableString** | Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: White. | [optional] [default to "White"]
**Units** | [***GraphicsUnit**](GraphicsUnit.md) |  | [optional]
**Resolution** | **NullableFloat32** | Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.  Decimal separator is dot. | [optional]
**ImageHeight** | **NullableFloat32** | Height of the barcode image in given units. Default units: pixel.  Decimal separator is dot. | [optional]
**ImageWidth** | **NullableFloat32** | Width of the barcode image in given units. Default units: pixel.  Decimal separator is dot. | [optional]
**RotationAngle** | **NullableInt32** | BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
