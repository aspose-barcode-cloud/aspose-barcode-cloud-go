# Pdf417Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | **float64** | Height/Width ratio of 2D BarCode module. | [optional] [default to null]
**TextEncoding** | **string** | Encoding of codetext. | [optional] [default to null]
**Columns** | **int32** | Columns count. | [optional] [default to null]
**CompactionMode** | [***Pdf417CompactionMode**](Pdf417CompactionMode.md) | Pdf417 symbology type of BarCode&#39;s compaction mode. Default value: Pdf417CompactionMode.Auto. | [optional] [default to null]
**ErrorLevel** | [***Pdf417ErrorLevel**](Pdf417ErrorLevel.md) | Pdf417 symbology type of BarCode&#39;s error correction level ranging from level0 to level8, level0 means no error correction info, level8 means best error correction which means a larger picture. | [optional] [default to null]
**MacroFileID** | **int32** | Macro Pdf417 barcode&#39;s file ID. Used for MacroPdf417. | [optional] [default to null]
**MacroSegmentID** | **int32** | Macro Pdf417 barcode&#39;s segment ID, which starts from 0, to MacroSegmentsCount - 1. | [optional] [default to null]
**MacroSegmentsCount** | **int32** | Macro Pdf417 barcode segments count. | [optional] [default to null]
**Rows** | **int32** | Rows count. | [optional] [default to null]
**Truncate** | **bool** | Whether Pdf417 symbology type of BarCode is truncated (to reduce space). | [optional] [default to null]
**Pdf417ECIEncoding** | [***EciEncodings**](EciEncodings.md) | Extended Channel Interpretation Identifiers. It is used to tell the barcode reader details about the used references for encoding the data in the symbol. Current implementation consists all well known charset encodings. | [optional] [default to null]
**IsReaderInitialization** | **bool** | Used to instruct the reader to interpret the data contained within the symbol as programming for reader initialization | [optional] [default to null]
**MacroTimeStamp** | **time.Time** | Macro Pdf417 barcode time stamp | [optional] [default to null]
**MacroSender** | **string** | Macro Pdf417 barcode sender name | [optional] [default to null]
**MacroFileSize** | **int32** | Macro Pdf417 file size. The file size field contains the size in bytes of the entire source file | [optional] [default to null]
**MacroChecksum** | **int32** | Macro Pdf417 barcode checksum. The checksum field contains the value of the 16-bit (2 bytes) CRC checksum using the CCITT-16 polynomial | [optional] [default to null]
**MacroFileName** | **string** | Macro Pdf417 barcode file name | [optional] [default to null]
**MacroAddressee** | **string** | Macro Pdf417 barcode addressee name | [optional] [default to null]
**MacroECIEncoding** | [***EciEncodings**](EciEncodings.md) | Extended Channel Interpretation Identifiers. Applies for Macro PDF417 text fields. | [optional] [default to null]
**Code128Emulation** | [***Code128Emulation**](Code128Emulation.md) | Function codeword for Code 128 emulation. Applied for MicroPDF417 only. Ignored for PDF417 and MacroPDF417 barcodes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
