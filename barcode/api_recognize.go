package barcode

import (
	"context"
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

// RecognizeAPIRecognizeOpts - Optional Parameters for RecognizeAPIRecognize
type RecognizeAPIRecognizeOpts struct {
	RecognitionMode      optional.Interface
	RecognitionImageKind optional.Interface
}

/*
* Recognize -  Recognize barcode from file on server using GET requests with parameters in route and query string.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType Type of barcode to recognize
* @param fileUrl Url to barcode image
* @param optional nil or *RecognizeAPIRecognizeOpts - Optional Parameters:
  - @param "RecognitionMode" (optional.Interface of RecognitionMode) -  Recognition mode
  - @param "RecognitionImageKind" (optional.Interface of RecognitionImageKind) -  Image kind for recognition

* @return BarcodeResponseList
*/
func (a *RecognizeAPIService) Recognize(ctx context.Context, barcodeType DecodeBarcodeType, fileUrl string, optionals *RecognizeAPIRecognizeOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Get")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/recognize"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("barcodeType", parameterToString(barcodeType, ""))
	queryParams.Add("fileUrl", parameterToString(fileUrl, ""))
	if optionals != nil && optionals.RecognitionMode.IsSet() {
		queryParams.Add("recognitionMode", parameterToString(optionals.RecognitionMode.Value(), ""))
	}
	if optionals != nil && optionals.RecognitionImageKind.IsSet() {
		queryParams.Add("recognitionImageKind", parameterToString(optionals.RecognitionImageKind.Value(), ""))
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
* RecognizeBase64 -  Recognize barcode from file in request body using POST requests with parameters in body in json or xml format.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param recognizeBase64Request Barcode recognition request

* @return BarcodeResponseList
 */
func (a *RecognizeAPIService) RecognizeBase64(ctx context.Context, recognizeBase64Request RecognizeBase64Request) (BarcodeResponseList, *http.Response, error) {
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

// RecognizeAPIRecognizeMultipartOpts - Optional Parameters for RecognizeAPIRecognizeMultipart
type RecognizeAPIRecognizeMultipartOpts struct {
	RecognitionMode      optional.Interface
	RecognitionImageKind optional.Interface
}

/*
* RecognizeMultipart -  Recognize barcode from file in request body using POST requests with parameters in multipart form.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType
* @param file Barcode image file
* @param optional nil or *RecognizeAPIRecognizeMultipartOpts - Optional Parameters:
  - @param "RecognitionMode" (optional.Interface of RecognitionMode) -
  - @param "RecognitionImageKind" (optional.Interface of RecognitionImageKind) -

* @return BarcodeResponseList
*/
func (a *RecognizeAPIService) RecognizeMultipart(ctx context.Context, barcodeType DecodeBarcodeType, file *os.File, optionals *RecognizeAPIRecognizeMultipartOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Post")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/recognize-multipart"

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
	if optionals != nil && optionals.RecognitionImageKind.IsSet() {
		formParams.Add("recognitionImageKind", parameterToString(optionals.RecognitionImageKind.Value(), ""))
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
