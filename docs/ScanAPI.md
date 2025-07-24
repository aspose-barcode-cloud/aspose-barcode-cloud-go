# barcode\ScanAPI

All URIs are relative to *<https://api.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**Scan**](ScanAPI.md#Scan) | **Get** /barcode/scan | Scan barcode from file on server in the Internet using GET requests with parameter in query string. For scaning files from your hard drive use &#x60;scan-body&#x60; or &#x60;scan-multipart&#x60; endpoints instead.
[**ScanBase64**](ScanAPI.md#ScanBase64) | **Post** /barcode/scan-body | Scan barcode from file in request body using POST requests with parameter in body in json or xml format.
[**ScanMultipart**](ScanAPI.md#ScanMultipart) | **Post** /barcode/scan-multipart | Scan barcode from file in request body using POST requests with parameter in multipart form.

## Scan

> BarcodeResponseList Scan(ctx, fileUrl)
Scan barcode from file on server in the Internet using GET requests with parameter in query string. For scaning files from your hard drive use `scan-body` or `scan-multipart` endpoints instead.

### Scan Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **fileUrl** | **string** | Url to barcode image |

### Scan Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## ScanBase64

> BarcodeResponseList ScanBase64(ctx, scanBase64Request)
Scan barcode from file in request body using POST requests with parameter in body in json or xml format.

### ScanBase64 Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **scanBase64Request** | [**ScanBase64Request**](ScanBase64Request.md) | Barcode scan request |

### ScanBase64 Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## ScanMultipart

> BarcodeResponseList ScanMultipart(ctx, file)
Scan barcode from file in request body using POST requests with parameter in multipart form.

### ScanMultipart Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **file** | ***os.File*****os.File** | Barcode image file |

### ScanMultipart Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
