# barcode\FolderApi

All URIs are relative to *<https://api.aspose.cloud/v3.0>*

Method | HTTP request | Description
------ | ------------ | -----------
[**CopyFolder**](FolderApi.md#CopyFolder) | **Put** /barcode/storage/folder/copy/{srcPath} | Copy folder
[**CreateFolder**](FolderApi.md#CreateFolder) | **Put** /barcode/storage/folder/{path} | Create the folder
[**DeleteFolder**](FolderApi.md#DeleteFolder) | **Delete** /barcode/storage/folder/{path} | Delete folder
[**GetFilesList**](FolderApi.md#GetFilesList) | **Get** /barcode/storage/folder/{path} | Get all files and folders within a folder
[**MoveFolder**](FolderApi.md#MoveFolder) | **Put** /barcode/storage/folder/move/{srcPath} | Move folder

## CopyFolder

> CopyFolder(ctx, srcPath, destPath, optional)
Copy folder

### CopyFolder Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **srcPath** | **string**| Source folder path e.g. &#39;/src&#39; |
 **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; |
 **optional** | ***FolderApiCopyFolderOpts** | optional parameters | nil if no parameters

### CopyFolder Optional Parameters

Optional parameters are passed through a pointer to a FolderApiCopyFolderOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----


 **srcStorageName** | **optional.String**| Source storage name |
 **destStorageName** | **optional.String**| Destination storage name |

### CopyFolder Return type

 (empty response body)

### CopyFolder Authorization

[JWT](../README.md#JWT)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## CreateFolder

> CreateFolder(ctx, path, optional)
Create the folder

### CreateFolder Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **path** | **string**| Folder path to create e.g. &#39;folder_1/folder_2/&#39; |
 **optional** | ***FolderApiCreateFolderOpts** | optional parameters | nil if no parameters

### CreateFolder Optional Parameters

Optional parameters are passed through a pointer to a FolderApiCreateFolderOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----

 **storageName** | **optional.String**| Storage name |

### CreateFolder Return type

 (empty response body)

### CreateFolder Authorization

[JWT](../README.md#JWT)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## DeleteFolder

> DeleteFolder(ctx, path, optional)
Delete folder

### DeleteFolder Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **path** | **string**| Folder path e.g. &#39;/folder&#39; |
 **optional** | ***FolderApiDeleteFolderOpts** | optional parameters | nil if no parameters

### DeleteFolder Optional Parameters

Optional parameters are passed through a pointer to a FolderApiDeleteFolderOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----

 **storageName** | **optional.String**| Storage name |
 **recursive** | **optional.Bool**| Enable to delete folders, subfolders and files | [default to false]

### DeleteFolder Return type

 (empty response body)

### DeleteFolder Authorization

[JWT](../README.md#JWT)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## GetFilesList

> FilesList GetFilesList(ctx, path, optional)
Get all files and folders within a folder

### GetFilesList Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **path** | **string**| Folder path e.g. &#39;/folder&#39; |
 **optional** | ***FolderApiGetFilesListOpts** | optional parameters | nil if no parameters

### GetFilesList Optional Parameters

Optional parameters are passed through a pointer to a FolderApiGetFilesListOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----

 **storageName** | **optional.String**| Storage name |

### GetFilesList Return type

[**FilesList**](FilesList.md)

### GetFilesList Authorization

[JWT](../README.md#JWT)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

## MoveFolder

> MoveFolder(ctx, srcPath, destPath, optional)
Move folder

### MoveFolder Required Parameters

Name | Type | Description  | Notes
---- | ---- | ------------ | -----
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **srcPath** | **string**| Folder path to move e.g. &#39;/folder&#39; |
 **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; |
 **optional** | ***FolderApiMoveFolderOpts** | optional parameters | nil if no parameters

### MoveFolder Optional Parameters

Optional parameters are passed through a pointer to a FolderApiMoveFolderOpts struct

Name | Type | Description  | Notes
---- | ---- | ------------ | -----


 **srcStorageName** | **optional.String**| Source storage name |
 **destStorageName** | **optional.String**| Destination storage name |

### MoveFolder Return type

 (empty response body)

### MoveFolder Authorization

[JWT](../README.md#JWT)

[[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
