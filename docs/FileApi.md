# barcode\FileApi

All URIs are relative to *<https://api.aspose.cloud/v3.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**CopyFile**](FileApi.md#CopyFile) | **Put** /barcode/storage/file/copy/{srcPath} | Copy file
[**DeleteFile**](FileApi.md#DeleteFile) | **Delete** /barcode/storage/file/{path} | Delete file
[**DownloadFile**](FileApi.md#DownloadFile) | **Get** /barcode/storage/file/{path} | Download file
[**MoveFile**](FileApi.md#MoveFile) | **Put** /barcode/storage/file/move/{srcPath} | Move file
[**UploadFile**](FileApi.md#UploadFile) | **Put** /barcode/storage/file/{path} | Upload file

## CopyFile

> CopyFile(ctx, srcPath, destPath, optional)
Copy file

### CopyFile Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **srcPath** | **string** | Source file path e.g. &#39;/folder/file.ext&#39; |
 **destPath** | **string** | Destination file path |
 **optional** | ***FileApiCopyFileOpts** | optional parameters | nil if no parameters

### CopyFile Optional Parameters

Optional parameters are passed through a pointer to a FileApiCopyFileOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----


 **SrcStorageName** | **optional.String** | Source storage name |
 **DestStorageName** | **optional.String** | Destination storage name |
 **VersionId** | **optional.String** | File version ID to copy |

### CopyFile Return type

 (empty response body)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## DeleteFile

> DeleteFile(ctx, path, optional)
Delete file

### DeleteFile Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **path** | **string** | File path e.g. &#39;/folder/file.ext&#39; |
 **optional** | ***FileApiDeleteFileOpts** | optional parameters | nil if no parameters

### DeleteFile Optional Parameters

Optional parameters are passed through a pointer to a FileApiDeleteFileOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----

 **StorageName** | **optional.String** | Storage name |
 **VersionId** | **optional.String** | File version ID to delete |

### DeleteFile Return type

 (empty response body)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## DownloadFile

> *os.File DownloadFile(ctx, path, optional)
Download file

### DownloadFile Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **path** | **string** | File path e.g. &#39;/folder/file.ext&#39; |
 **optional** | ***FileApiDownloadFileOpts** | optional parameters | nil if no parameters

### DownloadFile Optional Parameters

Optional parameters are passed through a pointer to a FileApiDownloadFileOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----

 **StorageName** | **optional.String** | Storage name |
 **VersionId** | **optional.String** | File version ID to download |

### DownloadFile Return type

**byte[]**

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## MoveFile

> MoveFile(ctx, srcPath, destPath, optional)
Move file

### MoveFile Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **srcPath** | **string** | Source file path e.g. &#39;/src.ext&#39; |
 **destPath** | **string** | Destination file path e.g. &#39;/dest.ext&#39; |
 **optional** | ***FileApiMoveFileOpts** | optional parameters | nil if no parameters

### MoveFile Optional Parameters

Optional parameters are passed through a pointer to a FileApiMoveFileOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----


 **SrcStorageName** | **optional.String** | Source storage name |
 **DestStorageName** | **optional.String** | Destination storage name |
 **VersionId** | **optional.String** | File version ID to move |

### MoveFile Return type

 (empty response body)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## UploadFile

> FilesUploadResult UploadFile(ctx, path, file, optional)
Upload file

### UploadFile Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **path** | **string** | Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.  |
 **file** | ***os.File** | File to upload |
 **optional** | ***FileApiUploadFileOpts** | optional parameters | nil if no parameters

### UploadFile Optional Parameters

Optional parameters are passed through a pointer to a FileApiUploadFileOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----


 **StorageName** | **optional.String** | Storage name |

### UploadFile Return type

[**FilesUploadResult**](FilesUploadResult.md)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
