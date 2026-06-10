# Pdf417Params

Optional PDF417 barcode generation parameters. Applies to Pdf417, MacroPdf417, MicroPdf417, and GS1MicroPdf417 barcode types.

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**Pdf417EncodeMode** | [***Pdf417EncodeMode**](Pdf417EncodeMode.md) |  | [optional]
**Pdf417ErrorLevel** | [***Pdf417ErrorLevel**](Pdf417ErrorLevel.md) |  | [optional]
**Pdf417Truncate** | **NullableBool** | Whether to use truncated PDF417 format (removes right-side stop pattern). | [optional]
**Pdf417Columns** | **NullableInt32** | Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto. | [optional]
**Pdf417Rows** | **NullableInt32** | Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic. | [optional]
**Pdf417AspectRatio** | **NullableFloat32** | PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417. | [optional]
**Pdf417ECIEncoding** | [***ECIEncodings**](ECIEncodings.md) |  | [optional]
**Pdf417IsReaderInitialization** | **NullableBool** | Whether the barcode is used for reader initialization (programming). | [optional]
**Pdf417MacroCharacters** | [***MacroCharacter**](MacroCharacter.md) |  | [optional]
**Pdf417IsLinked** | **NullableBool** | Whether to use linked mode (for MicroPdf417). | [optional]
**Pdf417IsCode128Emulation** | **NullableBool** | Whether to use Code128 emulation for MicroPdf417. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
