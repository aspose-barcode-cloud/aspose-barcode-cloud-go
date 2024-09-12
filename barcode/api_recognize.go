package barcode

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

// RecognizeAPIService -
type RecognizeAPIService service

// RecognizeAPIBarcodeRecognizeBarcodeTypeGetOpts - Optional Parameters for RecognizeAPIBarcodeRecognizeBarcodeTypeGet
type RecognizeAPIBarcodeRecognizeBarcodeTypeGetOpts struct {
	RecognitionMode optional.Interface
	ImageKind       optional.Interface
}

/*
* BarcodeRecognizeBarcodeTypeGet -  Recognize barcode from file on server using GET requests with parameters in route and query string.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType Type of barcode to recognize
* @param url Url to barcode image
* @param optional nil or *RecognizeAPIBarcodeRecognizeBarcodeTypeGetOpts - Optional Parameters:
  - @param "RecognitionMode" (optional.Interface of RecognitionMode) -  Recognition mode
  - @param "ImageKind" (optional.Interface of RecognitionImageKind) -  Image kind

* @return BarcodeResponseList
*/
func (a *RecognizeAPIService) BarcodeRecognizeBarcodeTypeGet(ctx context.Context, barcodeType DecodeBarcodeType, url string, optionals *RecognizeAPIBarcodeRecognizeBarcodeTypeGetOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Get")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/recognize/{barcodeType}"
	requestPath = strings.Replace(requestPath, "{"+"barcodeType"+"}", fmt.Sprintf("%v", barcodeType), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("url", parameterToString(url, ""))
	if optionals != nil && optionals.RecognitionMode.IsSet() {
		queryParams.Add("recognitionMode", parameterToString(optionals.RecognitionMode.Value(), ""))
	}
	if optionals != nil && optionals.ImageKind.IsSet() {
		queryParams.Add("imageKind", parameterToString(optionals.ImageKind.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json", "application/xml"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileFieldName, fileBytes)
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
			var v BarcodeResponseList
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		if httpResponse.StatusCode == 401 {
			var v ApiErrorResponse
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		if httpResponse.StatusCode == 400 {
			var v ApiErrorResponse
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
* BarcodeRecognizeBodyPost -  Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param recognizeBase64Request Barcode recognition request

* @return BarcodeResponseList
 */
func (a *RecognizeAPIService) BarcodeRecognizeBodyPost(ctx context.Context, recognizeBase64Request RecognizeBase64Request) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Post")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/recognize-body"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json", "application/xml"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json", "application/xml"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	// body params
	postBody = &recognizeBase64Request
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileFieldName, fileBytes)
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
			var v BarcodeResponseList
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		if httpResponse.StatusCode == 401 {
			var v ApiErrorResponse
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		if httpResponse.StatusCode == 400 {
			var v ApiErrorResponse
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

// RecognizeAPIBarcodeRecognizeFormPostOpts - Optional Parameters for RecognizeAPIBarcodeRecognizeFormPost
type RecognizeAPIBarcodeRecognizeFormPostOpts struct {
	RecognitionMode optional.Interface
	ImageKind       optional.Interface
}

/*
* BarcodeRecognizeFormPost -  Recognize barcode from file in request body using POST requests with parameters in multipart form.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType
* @param file
* @param optional nil or *RecognizeAPIBarcodeRecognizeFormPostOpts - Optional Parameters:
  - @param "RecognitionMode" (optional.Interface of RecognitionMode) -
  - @param "ImageKind" (optional.Interface of RecognitionImageKind) -

* @return BarcodeResponseList
*/
func (a *RecognizeAPIService) BarcodeRecognizeFormPost(ctx context.Context, barcodeType DecodeBarcodeType, file *os.File, optionals *RecognizeAPIBarcodeRecognizeFormPostOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Post")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/recognize-form"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypeChoices := []string{"multipart/form-data"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"application/json", "application/xml"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	formParams.Add("barcodeType", parameterToString(barcodeType, ""))
	requestFile := file
	if requestFile != nil {
		fileName = requestFile.Name()
		fileFieldName = "file"
		var err error
		fileBytes, err = io.ReadAll(io.Reader(requestFile))
		if err != nil {
			return returnValue, nil, err
		}
	}
	if optionals != nil && optionals.RecognitionMode.IsSet() {
		formParams.Add("recognitionMode", parameterToString(optionals.RecognitionMode.Value(), ""))
	}
	if optionals != nil && optionals.ImageKind.IsSet() {
		formParams.Add("imageKind", parameterToString(optionals.ImageKind.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, requestPath, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileFieldName, fileBytes)
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
			var v BarcodeResponseList
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		if httpResponse.StatusCode == 401 {
			var v ApiErrorResponse
			err = a.client.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}

		if httpResponse.StatusCode == 400 {
			var v ApiErrorResponse
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
