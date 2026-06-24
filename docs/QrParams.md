# QrParams

Optional QR barcode generation parameters. Applies to QR, GS1QR, MicroQR, and RectMicroQR barcode types.

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | -----
**QrEncodeMode** | [***QREncodeMode**](QREncodeMode.md) | QR barcode encode mode. | [optional]
**QrErrorLevel** | [***QRErrorLevel**](QRErrorLevel.md) | QR barcode error correction level. | [optional]
**QrVersion** | [***QRVersion**](QRVersion.md) | QR barcode version. Automatically selects the smallest version that fits the data. | [optional]
**QrECIEncoding** | [***ECIEncodings**](ECIEncodings.md) | ECI encoding for QR barcode data. | [optional]
**QrAspectRatio** | **NullableFloat32** | QR barcode aspect ratio. Values: 0 to 1. | [optional]
**MicroQRVersion** | [***MicroQRVersion**](MicroQRVersion.md) | MicroQR barcode version. Used when BarcodeType is MicroQR. | [optional]
**RectMicroQrVersion** | [***RectMicroQRVersion**](RectMicroQRVersion.md) | RectMicroQR barcode version. Used when BarcodeType is RectMicroQR. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
