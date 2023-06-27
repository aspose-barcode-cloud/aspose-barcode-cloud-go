/*
 * MIT License

 * Copyright (c) 2023 Aspose Pty Ltd

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
	"strings"
)

// Linger please
var (
	_ context.Context
)

// StorageApiService -
type StorageApiService service

// StorageApiGetDiscUsageOpts - Optional Parameters for StorageApiGetDiscUsage
type StorageApiGetDiscUsageOpts struct {
	StorageName optional.String
}

/*
* GetDiscUsage -  Get disc usage
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param optional nil or *StorageApiGetDiscUsageOpts - Optional Parameters:
  - @param "StorageName" (optional.String) -  Storage name

* @return DiscUsage
*/
func (a *StorageApiService) GetDiscUsage(ctx context.Context, optionals *StorageApiGetDiscUsageOpts) (DiscUsage, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue DiscUsage
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/disc"

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
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		if httpResponse.StatusCode == 200 {
			var v DiscUsage
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

// StorageApiGetFileVersionsOpts - Optional Parameters for StorageApiGetFileVersions
type StorageApiGetFileVersionsOpts struct {
	StorageName optional.String
}

/*
* GetFileVersions -  Get file versions
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param path File path e.g. &#39;/file.ext&#39;
* @param optional nil or *StorageApiGetFileVersionsOpts - Optional Parameters:
  - @param "StorageName" (optional.String) -  Storage name

* @return FileVersions
*/
func (a *StorageApiService) GetFileVersions(ctx context.Context, path string, optionals *StorageApiGetFileVersionsOpts) (FileVersions, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue FileVersions
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/version/{path}"
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
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		if httpResponse.StatusCode == 200 {
			var v FileVersions
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

// StorageApiObjectExistsOpts - Optional Parameters for StorageApiObjectExists
type StorageApiObjectExistsOpts struct {
	StorageName optional.String
	VersionId   optional.String
}

/*
* ObjectExists -  Check if file or folder exists
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param path File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39;
* @param optional nil or *StorageApiObjectExistsOpts - Optional Parameters:
  - @param "StorageName" (optional.String) -  Storage name
  - @param "VersionId" (optional.String) -  File version ID

* @return ObjectExist
*/
func (a *StorageApiService) ObjectExists(ctx context.Context, path string, optionals *StorageApiObjectExistsOpts) (ObjectExist, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue ObjectExist
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/exist/{path}"
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
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		if httpResponse.StatusCode == 200 {
			var v ObjectExist
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

/*
* StorageExists -  Check if storage exists
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param storageName Storage name

* @return StorageExist
 */
func (a *StorageApiService) StorageExists(ctx context.Context, storageName string) (StorageExist, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue StorageExist
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/storage/{storageName}/exist"
	requestPath = strings.Replace(requestPath, "{"+"storageName"+"}", fmt.Sprintf("%v", storageName), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

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
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
		}

		if httpResponse.StatusCode == 200 {
			var v StorageExist
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
