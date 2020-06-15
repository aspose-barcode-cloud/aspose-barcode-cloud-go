# Pdf417Params

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | **float64** | Height/Width ratio of 2D BarCode module.  | [optional] [default to null]
**TextEncoding** | **string** | Encoding of codetext.  | [optional] [default to null]
**Columns** | **int32** | Columns count.  | [optional] [default to null]
**CompactionMode** | [***Pdf417CompactionMode**](Pdf417CompactionMode.md) | Pdf417 symbology type of BarCode&#39;s compaction mode. Default value: Pdf417CompactionMode.Auto. | [optional] [default to null]
**ErrorLevel** | [***Pdf417ErrorLevel**](Pdf417ErrorLevel.md) | Pdf417 symbology type of BarCode&#39;s error correction level ranging from level0 to level8, level0 means no error correction info, level8 means best error correction which means a larger picture. | [optional] [default to null]
**MacroFileID** | **int32** | Macro Pdf417 barcode&#39;s file ID. Used for MacroPdf417. | [optional] [default to null]
**MacroSegmentID** | **int32** | Macro Pdf417 barcode&#39;s segment ID, which starts from 0, to MacroSegmentsCount - 1. | [optional] [default to null]
**MacroSegmentsCount** | **int32** | Macro Pdf417 barcode segments count. | [optional] [default to null]
**Rows** | **int32** | Rows count. | [optional] [default to null]
**Truncate** | **bool** | Whether Pdf417 symbology type of BarCode is truncated (to reduce space). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


