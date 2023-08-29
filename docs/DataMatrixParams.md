# DataMatrixParams

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**AspectRatio** | **float64** | Height/Width ratio of 2D BarCode module | [optional] [default to null]
**TextEncoding** | **string** | Encoding of codetext. | [optional] [default to null]
**Columns** | **int32** | DEPRECATED: Will be replaced with &#39;DataMatrix.Version&#39; in the next release  Columns count. | [optional] [default to null]
**DataMatrixEcc** | [***DataMatrixEccType**](DataMatrixEccType.md) | Datamatrix ECC type. Default value: DataMatrixEccType.Ecc200. | [optional] [default to null]
**DataMatrixEncodeMode** | [***DataMatrixEncodeMode**](DataMatrixEncodeMode.md) | Encode mode of Datamatrix barcode. Default value: DataMatrixEncodeMode.Auto. | [optional] [default to null]
**Rows** | **int32** | DEPRECATED: Will be replaced with &#39;DataMatrix.Version&#39; in the next release  Rows count. | [optional] [default to null]
**MacroCharacters** | [***MacroCharacter**](MacroCharacter.md) | Macro Characters 05 and 06 values are used to obtain more compact encoding in special modes. Can be used only with DataMatrixEccType.Ecc200 or DataMatrixEccType.EccAuto. Cannot be used with EncodeTypes.GS1DataMatrix Default value: MacroCharacters.None. | [optional] [default to null]
**Version** | [***DataMatrixVersion**](DataMatrixVersion.md) | Sets a Datamatrix symbol size. Default value: DataMatrixVersion.Auto. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
