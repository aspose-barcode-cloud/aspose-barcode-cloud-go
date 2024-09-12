# barcode\RecognizeAPI

All URIs are relative to *<https://barcode.qa.aspose.cloud/v4.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**BarcodeRecognizeBarcodeTypeGet**](RecognizeAPI.md#BarcodeRecognizeBarcodeTypeGet) | **Get** /barcode/recognize/{barcodeType} | Recognize barcode from file on server using GET requests with parameters in route and query string.
[**BarcodeRecognizeBodyPost**](RecognizeAPI.md#BarcodeRecognizeBodyPost) | **Post** /barcode/recognize-body | Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.
[**BarcodeRecognizeFormPost**](RecognizeAPI.md#BarcodeRecognizeFormPost) | **Post** /barcode/recognize-form | Recognize barcode from file in request body using POST requests with parameters in multipart form.

## BarcodeRecognizeBarcodeTypeGet

> BarcodeResponseList BarcodeRecognizeBarcodeTypeGet(ctx, barcodeType, url, optional)
Recognize barcode from file on server using GET requests with parameters in route and query string.

### BarcodeRecognizeBarcodeTypeGet Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**DecodeBarcodeType**](.md) | Type of barcode to recognize |
 **url** | **string** | Url to barcode image |
 **optional** | ***RecognizeAPIBarcodeRecognizeBarcodeTypeGetOpts** | optional parameters | nil if no parameters

### BarcodeRecognizeBarcodeTypeGet Optional Parameters

Optional parameters are passed through a pointer to a RecognizeAPIBarcodeRecognizeBarcodeTypeGetOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**RecognitionMode** | [**optional.Interface of RecognitionMode**](.md) | Recognition mode |
**ImageKind** | [**optional.Interface of RecognitionImageKind**](.md) | Image kind |

### BarcodeRecognizeBarcodeTypeGet Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

## BarcodeRecognizeFormPost

> BarcodeResponseList BarcodeRecognizeFormPost(ctx, barcodeType, file, optional)
Recognize barcode from file in request body using POST requests with parameters in multipart form.

### BarcodeRecognizeFormPost Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **barcodeType** | [**DecodeBarcodeType**](DecodeBarcodeType.md) |  |
 **file** | ***os.File*****os.File** |  |
 **optional** | ***RecognizeAPIBarcodeRecognizeFormPostOpts** | optional parameters | nil if no parameters

### BarcodeRecognizeFormPost Optional Parameters

Optional parameters are passed through a pointer to a RecognizeAPIBarcodeRecognizeFormPostOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
**RecognitionMode** | [**optional.Interface of RecognitionMode**](RecognitionMode.md) |  |
**ImageKind** | [**optional.Interface of RecognitionImageKind**](RecognitionImageKind.md) |  |

### BarcodeRecognizeFormPost Return type

[**BarcodeResponseList**](BarcodeResponseList.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
