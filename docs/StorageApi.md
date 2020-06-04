# \StorageApi

All URIs are relative to *https://api.aspose.cloud/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiscUsage**](StorageApi.md#GetDiscUsage) | **Get** /barcode/storage/disc | Get disc usage
[**GetFileVersions**](StorageApi.md#GetFileVersions) | **Get** /barcode/storage/version/{path} | Get file versions
[**ObjectExists**](StorageApi.md#ObjectExists) | **Get** /barcode/storage/exist/{path} | Check if file or folder exists
[**StorageExists**](StorageApi.md#StorageExists) | **Get** /barcode/storage/{storageName}/exist | Check if storage exists


# **GetDiscUsage**
> DiscUsage GetDiscUsage(ctx, optional)
Get disc usage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiGetDiscUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiGetDiscUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageName** | **optional.String**| Storage name | 

### Return type

[**DiscUsage**](DiscUsage.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileVersions**
> FileVersions GetFileVersions(ctx, path, optional)
Get file versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| File path e.g. &#39;/file.ext&#39; | 
 **optional** | ***StorageApiGetFileVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiGetFileVersionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 

### Return type

[**FileVersions**](FileVersions.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectExists**
> ObjectExist ObjectExists(ctx, path, optional)
Check if file or folder exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39; | 
 **optional** | ***StorageApiObjectExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiObjectExistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 
 **versionId** | **optional.String**| File version ID | 

### Return type

[**ObjectExist**](ObjectExist.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorageExists**
> StorageExist StorageExists(ctx, storageName)
Check if storage exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storageName** | **string**| Storage name | 

### Return type

[**StorageExist**](StorageExist.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

