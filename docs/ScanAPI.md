# barcode\ScanAPI

All URIs are relative to *<https://barcode.qa.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**BarcodeScanBodyPost**](ScanAPI.md#BarcodeScanBodyPost) | **Post** /barcode/scan-body | Scan barcode from file in request body using POST requests with parameter in body in json or xml format.
[**BarcodeScanFormPost**](ScanAPI.md#BarcodeScanFormPost) | **Post** /barcode/scan-form | Scan barcode from file in request body using POST requests with parameter in multipart form.
[**BarcodeScanGet**](ScanAPI.md#BarcodeScanGet) | **Get** /barcode/scan | Scan barcode from file on server using GET requests with parameter in query string.

## BarcodeScanBodyPost

> BarcodeResponseList BarcodeScanBodyPost(ctx, scanBase64Request)
Scan barcode from file in request body using POST requests with parameter in body in json or xml format.

### BarcodeScanBodyPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **scanBase64Request** | [**ScanBase64Request**](ScanBase64Request.md) | Barcode scan request |

### BarcodeScanBodyPost Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## BarcodeScanFormPost

> BarcodeResponseList BarcodeScanFormPost(ctx, file)
Scan barcode from file in request body using POST requests with parameter in multipart form.

### BarcodeScanFormPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **file** | ***os.File*****os.File** |  |

### BarcodeScanFormPost Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## BarcodeScanGet

> BarcodeResponseList BarcodeScanGet(ctx, fileUrl)
Scan barcode from file on server using GET requests with parameter in query string.

### BarcodeScanGet Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **fileUrl** | **string** | Url to barcode image |

### BarcodeScanGet Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
