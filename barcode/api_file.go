/*
 * MIT License

 * Copyright (c) 2021 Aspose Pty Ltd

 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:

 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.

 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package barcode

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Linger please
var (
	_ context.Context
)

//FileApiService -
type FileApiService service

//FileApiCopyFileOpts - Optional Parameters for FileApiCopyFile
type FileApiCopyFileOpts struct {
	SrcStorageName  optional.String
	DestStorageName optional.String
	VersionId       optional.String
}

/*
 * CopyFile -  Copy file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param srcPath Source file path e.g. &#39;/folder/file.ext&#39;
 * @param destPath Destination file path
 * @param optional nil or *FileApiCopyFileOpts - Optional Parameters:
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name
     * @param "VersionId" (optional.String) -  File version ID to copy


*/
func (a *FileApiService) CopyFile(ctx context.Context, srcPath string, destPath string, optionals *FileApiCopyFileOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/file/copy/{srcPath}"
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
	if optionals != nil && optionals.VersionId.IsSet() {
		queryParams.Add("versionId", parameterToString(optionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine the Accept header
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

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error: httpResponse.Status,
			text:  string(responseBody),
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}

//FileApiDeleteFileOpts - Optional Parameters for FileApiDeleteFile
type FileApiDeleteFileOpts struct {
	StorageName optional.String
	VersionId   optional.String
}

/*
 * DeleteFile -  Delete file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path File path e.g. &#39;/folder/file.ext&#39;
 * @param optional nil or *FileApiDeleteFileOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID to delete


*/
func (a *FileApiService) DeleteFile(ctx context.Context, path string, optionals *FileApiDeleteFileOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/file/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.StorageName.IsSet() {
		queryParams.Add("storageName", parameterToString(optionals.StorageName.Value(), ""))
	}
	if optionals != nil && optionals.VersionId.IsSet() {
		queryParams.Add("versionId", parameterToString(optionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine the Accept header
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

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error: httpResponse.Status,
			text:  string(responseBody),
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}

//FileApiDownloadFileOpts - Optional Parameters for FileApiDownloadFile
type FileApiDownloadFileOpts struct {
	StorageName optional.String
	VersionId   optional.String
}

/*
 * DownloadFile -  Download file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path File path e.g. &#39;/folder/file.ext&#39;
 * @param optional nil or *FileApiDownloadFileOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID to download

 * @return []byte
*/
func (a *FileApiService) DownloadFile(ctx context.Context, path string, optionals *FileApiDownloadFileOpts) ([]byte, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/file/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.StorageName.IsSet() {
		queryParams.Add("storageName", parameterToString(optionals.StorageName.Value(), ""))
	}
	if optionals != nil && optionals.VersionId.IsSet() {
		queryParams.Add("versionId", parameterToString(optionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine the Accept header
	acceptChoices := []string{"multipart/form-data"}

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

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return returnValue, httpResponse, err
	}

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return responseBody, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error: httpResponse.Status,
			text:  string(responseBody),
		}

		if httpResponse.StatusCode == 200 {
			var v *os.File
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

//FileApiMoveFileOpts - Optional Parameters for FileApiMoveFile
type FileApiMoveFileOpts struct {
	SrcStorageName  optional.String
	DestStorageName optional.String
	VersionId       optional.String
}

/*
 * MoveFile -  Move file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param srcPath Source file path e.g. &#39;/src.ext&#39;
 * @param destPath Destination file path e.g. &#39;/dest.ext&#39;
 * @param optional nil or *FileApiMoveFileOpts - Optional Parameters:
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name
     * @param "VersionId" (optional.String) -  File version ID to move


*/
func (a *FileApiService) MoveFile(ctx context.Context, srcPath string, destPath string, optionals *FileApiMoveFileOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/file/move/{srcPath}"
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
	if optionals != nil && optionals.VersionId.IsSet() {
		queryParams.Add("versionId", parameterToString(optionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine the Accept header
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

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error: httpResponse.Status,
			text:  string(responseBody),
		}

		return httpResponse, newErr
	}

	return httpResponse, err
}

//FileApiUploadFileOpts - Optional Parameters for FileApiUploadFile
type FileApiUploadFileOpts struct {
	StorageName optional.String
}

/*
 * UploadFile -  Upload file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext  If the content is multipart and path does not contains the file name it tries to get them from filename parameter  from Content-Disposition header.
 * @param file File to upload
 * @param optional nil or *FileApiUploadFileOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name

 * @return FilesUploadResult
*/
func (a *FileApiService) UploadFile(ctx context.Context, path string, file *os.File, optionals *FileApiUploadFileOpts) (FilesUploadResult, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue FilesUploadResult
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/file/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.StorageName.IsSet() {
		queryParams.Add("storageName", parameterToString(optionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"multipart/form-data"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine the Accept header
	acceptChoices := []string{"application/json"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	requestFile := file
	if requestFile != nil {
		var err error
		postBody, err = ioutil.ReadAll(requestFile)
		if err != nil {
			return returnValue, nil, err
		}
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := a.client.callAPI(r)
	if err != nil || httpResponse == nil {
		return returnValue, httpResponse, err
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
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
			error: httpResponse.Status,
			text:  string(responseBody),
		}

		if httpResponse.StatusCode == 200 {
			var v FilesUploadResult
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
