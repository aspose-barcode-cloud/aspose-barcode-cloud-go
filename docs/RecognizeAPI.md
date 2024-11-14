# barcode\RecognizeAPI

All URIs are relative to *<https://barcode.qa.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**BarcodeRecognizeBodyPost**](RecognizeAPI.md#BarcodeRecognizeBodyPost) | **Post** /barcode/recognize-body | Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.
[**BarcodeRecognizeGet**](RecognizeAPI.md#BarcodeRecognizeGet) | **Get** /barcode/recognize | Recognize barcode from file on server using GET requests with parameters in route and query string.
[**BarcodeRecognizeMultipartPost**](RecognizeAPI.md#BarcodeRecognizeMultipartPost) | **Post** /barcode/recognize-multipart | Recognize barcode from file in request body using POST requests with parameters in multipart form.

## BarcodeRecognizeBodyPost

> BarcodeResponseList BarcodeRecognizeBodyPost(ctx, recognizeBase64Request)
Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.

### BarcodeRecognizeBodyPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **recognizeBase64Request** | [**RecognizeBase64Request**](RecognizeBase64Request.md) | Barcode recognition request |

### BarcodeRecognizeBodyPost Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## BarcodeRecognizeGet

> BarcodeResponseList BarcodeRecognizeGet(ctx, barcodeType, fileUrl, optional)
Recognize barcode from file on server using GET requests with parameters in route and query string.

### BarcodeRecognizeGet Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**DecodeBarcodeType**](.md) | Type of barcode to recognize |
 **fileUrl** | **string** | Url to barcode image |
 **optional** | ***RecognizeAPIBarcodeRecognizeGetOpts** | optional parameters | nil if no parameters

### BarcodeRecognizeGet Optional Parameters

Optional parameters are passed through a pointer to a RecognizeAPIBarcodeRecognizeGetOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**RecognitionMode** | [**optional.Interface of RecognitionMode**](.md) | Recognition mode |
**RecognitionImageKind** | [**optional.Interface of RecognitionImageKind**](.md) | Image kind for recognition |

### BarcodeRecognizeGet Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## BarcodeRecognizeMultipartPost

> BarcodeResponseList BarcodeRecognizeMultipartPost(ctx, barcodeType, file, optional)
Recognize barcode from file in request body using POST requests with parameters in multipart form.

### BarcodeRecognizeMultipartPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**DecodeBarcodeType**](DecodeBarcodeType.md) |  |
 **file** | ***os.File*****os.File** | Barcode image file |
 **optional** | ***RecognizeAPIBarcodeRecognizeMultipartPostOpts** | optional parameters | nil if no parameters

### BarcodeRecognizeMultipartPost Optional Parameters

Optional parameters are passed through a pointer to a RecognizeAPIBarcodeRecognizeMultipartPostOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**RecognitionMode** | [**optional.Interface of RecognitionMode**](RecognitionMode.md) |  |
**RecognitionImageKind** | [**optional.Interface of RecognitionImageKind**](RecognitionImageKind.md) |  |

### BarcodeRecognizeMultipartPost Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
