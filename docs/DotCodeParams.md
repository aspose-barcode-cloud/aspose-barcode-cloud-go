# DotCodeParams

DotCode parameters.

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**AspectRatio** | **float64** | Height/Width ratio of 2D BarCode module. | [optional] [default to null]
**Columns** | **int32** | Identifies columns count. Sum of the number of rows plus the number of columns of a DotCode symbol must be odd. Number of columns must be at least 5. | [optional] [default to null]
**EncodeMode** | [***DotCodeEncodeMode**](DotCodeEncodeMode.md) | Identifies DotCode encode mode. Default value: Auto. | [optional] [default to null]
**ECIEncoding** | [***EciEncodings**](EciEncodings.md) | Identifies ECI encoding. Used when DotCodeEncodeMode is Auto. Default value: ISO-8859-1. | [optional] [default to null]
**IsReaderInitialization** | **bool** | Indicates whether code is used for instruct reader to interpret the following data as instructions for initialization or reprogramming of the bar code reader. Default value is false. | [optional] [default to null]
**Rows** | **int32** | Identifies rows count. Sum of the number of rows plus the number of columns of a DotCode symbol must be odd. Number of rows must be at least 5. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
