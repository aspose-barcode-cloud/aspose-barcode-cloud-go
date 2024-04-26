package barcode

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

// FolderApiService -
type FolderApiService service

// FolderApiCopyFolderOpts - Optional Parameters for FolderApiCopyFolder
type FolderApiCopyFolderOpts struct {
	SrcStorageName  optional.String
	DestStorageName optional.String
}

/*
* CopyFolder -  Copy folder
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param srcPath Source folder path e.g. &#39;/src&#39;
* @param destPath Destination folder path e.g. &#39;/dst&#39;
* @param optional nil or *FolderApiCopyFolderOpts - Optional Parameters:
  - @param "SrcStorageName" (optional.String) -  Source storage name
  - @param "DestStorageName" (optional.String) -  Destination storage name
*/
func (a *FolderApiService) CopyFolder(ctx context.Context, srcPath string, destPath string, optionals *FolderApiCopyFolderOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/folder/copy/{srcPath}"
	requestPath = strings.Replace(requestPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("destPath", parameterToString(destPath, ""))
	if optionals != nil && optionals.SrcStorageName.IsSet() {
		queryParams.Add("srcStorageName", parameterToString(optionals.SrcStorageName.Value(), ""))
	}
	if optionals != nil && optionals.DestStorageName.IsSet() {
		queryParams.Add("destStorageName", parameterToString(optionals.DestStorageName.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	httpResponse, err := a.client.callAPI(r)
	if err != nil || httpResponse == nil {
		return httpResponse, err
	}

	responseBody, err := io.ReadAll(io.Reader(httpResponse.Body))
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}

// FolderApiCreateFolderOpts - Optional Parameters for FolderApiCreateFolder
type FolderApiCreateFolderOpts struct {
	StorageName optional.String
}

/*
* CreateFolder -  Create the folder
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param path Folder path to create e.g. &#39;folder_1/folder_2/&#39;
* @param optional nil or *FolderApiCreateFolderOpts - Optional Parameters:
  - @param "StorageName" (optional.String) -  Storage name
*/
func (a *FolderApiService) CreateFolder(ctx context.Context, path string, optionals *FolderApiCreateFolderOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/folder/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.StorageName.IsSet() {
		queryParams.Add("storageName", parameterToString(optionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	httpResponse, err := a.client.callAPI(r)
	if err != nil || httpResponse == nil {
		return httpResponse, err
	}

	responseBody, err := io.ReadAll(io.Reader(httpResponse.Body))
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}

// FolderApiDeleteFolderOpts - Optional Parameters for FolderApiDeleteFolder
type FolderApiDeleteFolderOpts struct {
	StorageName optional.String
	Recursive   optional.Bool
}

/*
* DeleteFolder -  Delete folder
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param path Folder path e.g. &#39;/folder&#39;
* @param optional nil or *FolderApiDeleteFolderOpts - Optional Parameters:
  - @param "StorageName" (optional.String) -  Storage name
  - @param "Recursive" (optional.Bool) -  Enable to delete folders, subfolders and files
*/
func (a *FolderApiService) DeleteFolder(ctx context.Context, path string, optionals *FolderApiDeleteFolderOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/folder/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.StorageName.IsSet() {
		queryParams.Add("storageName", parameterToString(optionals.StorageName.Value(), ""))
	}
	if optionals != nil && optionals.Recursive.IsSet() {
		queryParams.Add("recursive", parameterToString(optionals.Recursive.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	httpResponse, err := a.client.callAPI(r)
	if err != nil || httpResponse == nil {
		return httpResponse, err
	}

	responseBody, err := io.ReadAll(io.Reader(httpResponse.Body))
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}

// FolderApiGetFilesListOpts - Optional Parameters for FolderApiGetFilesList
type FolderApiGetFilesListOpts struct {
	StorageName optional.String
}

/*
* GetFilesList -  Get all files and folders within a folder
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param path Folder path e.g. &#39;/folder&#39;
* @param optional nil or *FolderApiGetFilesListOpts - Optional Parameters:
  - @param "StorageName" (optional.String) -  Storage name

* @return FilesList
*/
func (a *FolderApiService) GetFilesList(ctx context.Context, path string, optionals *FolderApiGetFilesListOpts) (FilesList, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue FilesList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/folder/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.StorageName.IsSet() {
		queryParams.Add("storageName", parameterToString(optionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := a.client.callAPI(r)
	if err != nil || httpResponse == nil {
		return returnValue, httpResponse, err
	}

	responseBody, err := io.ReadAll(io.Reader(httpResponse.Body))
	httpResponse.Body.Close()
	if err != nil {
		return returnValue, httpResponse, err
	}

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		if httpResponse.StatusCode == 200 {
			var v FilesList
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, err
}

// FolderApiMoveFolderOpts - Optional Parameters for FolderApiMoveFolder
type FolderApiMoveFolderOpts struct {
	SrcStorageName  optional.String
	DestStorageName optional.String
}

/*
* MoveFolder -  Move folder
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param srcPath Folder path to move e.g. &#39;/folder&#39;
* @param destPath Destination folder path to move to e.g &#39;/dst&#39;
* @param optional nil or *FolderApiMoveFolderOpts - Optional Parameters:
  - @param "SrcStorageName" (optional.String) -  Source storage name
  - @param "DestStorageName" (optional.String) -  Destination storage name
*/
func (a *FolderApiService) MoveFolder(ctx context.Context, srcPath string, destPath string, optionals *FolderApiMoveFolderOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/folder/move/{srcPath}"
	requestPath = strings.Replace(requestPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("destPath", parameterToString(destPath, ""))
	if optionals != nil && optionals.SrcStorageName.IsSet() {
		queryParams.Add("srcStorageName", parameterToString(optionals.SrcStorageName.Value(), ""))
	}
	if optionals != nil && optionals.DestStorageName.IsSet() {
		queryParams.Add("destStorageName", parameterToString(optionals.DestStorageName.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	httpResponse, err := a.client.callAPI(r)
	if err != nil || httpResponse == nil {
		return httpResponse, err
	}

	responseBody, err := io.ReadAll(io.Reader(httpResponse.Body))
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}
