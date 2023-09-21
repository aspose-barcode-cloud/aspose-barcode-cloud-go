# AztecParams

Aztec parameters.

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**AspectRatio** | **float64** | Height/Width ratio of 2D BarCode module. | [optional] [default to null]
**ErrorLevel** | **int32** | Level of error correction of Aztec types of barcode. Value should between 10 to 95. | [optional] [default to null]
**SymbolMode** | [***AztecSymbolMode**](AztecSymbolMode.md) | Aztec Symbol mode. Default value: AztecSymbolMode.Auto. | [optional] [default to null]
**TextEncoding** | **string** | DEPRECATED: This property is obsolete and will be removed in future releases. Unicode symbols detection and encoding will be processed in Auto mode with Extended Channel Interpretation charset designator. Using of own encodings requires manual CodeText encoding into byte[] array.  Sets the encoding of codetext. | [optional] [default to null]
**EncodeMode** | [***AztecEncodeMode**](AztecEncodeMode.md) | Encoding mode for Aztec barcodes. Default value: Auto | [optional] [default to null]
**ECIEncoding** | [***EciEncodings**](EciEncodings.md) | Identifies ECI encoding. Used when AztecEncodeMode is Auto. Default value: ISO-8859-1. | [optional] [default to null]
**IsReaderInitialization** | **bool** | Used to instruct the reader to interpret the data contained within the symbol as programming for reader initialization. | [optional] [default to null]
**LayersCount** | **int32** | Gets or sets layers count of Aztec symbol. Layers count should be in range from 1 to 3 for Compact mode and in range from 1 to 32 for Full Range mode. Default value: 0 (auto). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
