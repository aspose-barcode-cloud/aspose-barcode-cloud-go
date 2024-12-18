# barcode\RecognizeAPI

All URIs are relative to *<https://api.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**Recognize**](RecognizeAPI.md#Recognize) | **Get** /barcode/recognize | Recognize barcode from file on server using GET requests with parameters in route and query string.
[**RecognizeBase64**](RecognizeAPI.md#RecognizeBase64) | **Post** /barcode/recognize-body | Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.
[**RecognizeMultipart**](RecognizeAPI.md#RecognizeMultipart) | **Post** /barcode/recognize-multipart | Recognize barcode from file in request body using POST requests with parameters in multipart form.

## Recognize

> BarcodeResponseList Recognize(ctx, barcodeType, fileUrl, optional)
Recognize barcode from file on server using GET requests with parameters in route and query string.

### Recognize Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**DecodeBarcodeType**](.md) | Type of barcode to recognize |
 **fileUrl** | **string** | Url to barcode image |
 **optional** | ***RecognizeAPIRecognizeOpts** | optional parameters | nil if no parameters

### Recognize Optional Parameters

Optional parameters are passed through a pointer to a RecognizeAPIRecognizeOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**RecognitionMode** | [**optional.Interface of RecognitionMode**](.md) | Recognition mode |
**RecognitionImageKind** | [**optional.Interface of RecognitionImageKind**](.md) | Image kind for recognition |

### Recognize Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## RecognizeBase64

> BarcodeResponseList RecognizeBase64(ctx, recognizeBase64Request)
Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.

### RecognizeBase64 Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **recognizeBase64Request** | [**RecognizeBase64Request**](RecognizeBase64Request.md) | Barcode recognition request |

### RecognizeBase64 Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## RecognizeMultipart

> BarcodeResponseList RecognizeMultipart(ctx, barcodeType, file, optional)
Recognize barcode from file in request body using POST requests with parameters in multipart form.

### RecognizeMultipart Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**DecodeBarcodeType**](DecodeBarcodeType.md) |  |
 **file** | ***os.File*****os.File** | Barcode image file |
 **optional** | ***RecognizeAPIRecognizeMultipartOpts** | optional parameters | nil if no parameters

### RecognizeMultipart Optional Parameters

Optional parameters are passed through a pointer to a RecognizeAPIRecognizeMultipartOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**RecognitionMode** | [**optional.Interface of RecognitionMode**](RecognitionMode.md) |  |
**RecognitionImageKind** | [**optional.Interface of RecognitionImageKind**](RecognitionImageKind.md) |  |

### RecognizeMultipart Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
