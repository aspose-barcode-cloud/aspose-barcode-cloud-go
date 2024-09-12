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

// GenerateAPIService -
type GenerateAPIService service

// GenerateAPIGenerateOpts - Optional Parameters for GenerateAPIGenerate
type GenerateAPIGenerateOpts struct {
	DataType        optional.Interface
	ImageFormat     optional.Interface
	TextLocation    optional.Interface
	ForegroundColor optional.String
	BackgroundColor optional.String
	Units           optional.Interface
	Resolution      optional.Float32
	ImageHeight     optional.Float32
	ImageWidth      optional.Float32
	RotationAngle   optional.Int32
}

/*
* Generate -  Generate barcode using GET request with parameters in route and query string.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType Type of barcode to generate.
* @param data String represents data to encode
* @param optional nil or *GenerateAPIGenerateOpts - Optional Parameters:
  - @param "DataType" (optional.Interface of EncodeDataType) -  Type of data to encode.  Default value: StringData.
  - @param "ImageFormat" (optional.Interface of BarcodeImageFormat) -  Barcode output image format.  Default value: png
  - @param "TextLocation" (optional.Interface of CodeLocation) -  Specify the displaying Text Location, set to CodeLocation.None to hide CodeText.  Default value: CodeLocation.Below.
  - @param "ForegroundColor" (optional.String) -  Specify the displaying bars and content Color.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: Black.
  - @param "BackgroundColor" (optional.String) -  Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: White.
  - @param "Units" (optional.Interface of GraphicsUnit) -  Common Units for all measuring in query. Default units: pixel.
  - @param "Resolution" (optional.Float32) -  Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.  Decimal separator is dot.
  - @param "ImageHeight" (optional.Float32) -  Height of the barcode image in given units. Default units: pixel.  Decimal separator is dot.
  - @param "ImageWidth" (optional.Float32) -  Width of the barcode image in given units. Default units: pixel.  Decimal separator is dot.
  - @param "RotationAngle" (optional.Int32) -  BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0.

* @return []byte
*/
func (a *GenerateAPIService) Generate(ctx context.Context, barcodeType EncodeBarcodeType, data string, optionals *GenerateAPIGenerateOpts) ([]byte, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Get")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/generate/{barcodeType}"
	requestPath = strings.Replace(requestPath, "{"+"barcodeType"+"}", fmt.Sprintf("%v", barcodeType), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.DataType.IsSet() {
		queryParams.Add("dataType", parameterToString(optionals.DataType.Value(), ""))
	}
	queryParams.Add("data", parameterToString(data, ""))
	if optionals != nil && optionals.ImageFormat.IsSet() {
		queryParams.Add("imageFormat", parameterToString(optionals.ImageFormat.Value(), ""))
	}
	if optionals != nil && optionals.TextLocation.IsSet() {
		queryParams.Add("textLocation", parameterToString(optionals.TextLocation.Value(), ""))
	}
	if optionals != nil && optionals.ForegroundColor.IsSet() {
		queryParams.Add("foregroundColor", parameterToString(optionals.ForegroundColor.Value(), ""))
	}
	if optionals != nil && optionals.BackgroundColor.IsSet() {
		queryParams.Add("backgroundColor", parameterToString(optionals.BackgroundColor.Value(), ""))
	}
	if optionals != nil && optionals.Units.IsSet() {
		queryParams.Add("units", parameterToString(optionals.Units.Value(), ""))
	}
	if optionals != nil && optionals.Resolution.IsSet() {
		queryParams.Add("resolution", parameterToString(optionals.Resolution.Value(), ""))
	}
	if optionals != nil && optionals.ImageHeight.IsSet() {
		queryParams.Add("imageHeight", parameterToString(optionals.ImageHeight.Value(), ""))
	}
	if optionals != nil && optionals.ImageWidth.IsSet() {
		queryParams.Add("imageWidth", parameterToString(optionals.ImageWidth.Value(), ""))
	}
	if optionals != nil && optionals.RotationAngle.IsSet() {
		queryParams.Add("rotationAngle", parameterToString(optionals.RotationAngle.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"image/png", "image/bmp", "image/gif", "image/jpeg", "image/svg+xml", "image/tiff", "application/json", "application/xml"}

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
			return responseBody, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
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
* GenerateBody -  Generate barcode using POST request with parameters in body in json or xml format.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param generateParams Parameters of generation

* @return []byte
 */
func (a *GenerateAPIService) GenerateBody(ctx context.Context, generateParams GenerateParams) ([]byte, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Post")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/generate-body"

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
	acceptChoices := []string{"image/png", "image/bmp", "image/gif", "image/jpeg", "image/svg+xml", "image/tiff", "application/json", "application/xml"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	// body params
	postBody = &generateParams
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
			return responseBody, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
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

// GenerateAPIGenerateMultipartOpts - Optional Parameters for GenerateAPIGenerateMultipart
type GenerateAPIGenerateMultipartOpts struct {
	DataType        optional.Interface
	ImageFormat     optional.Interface
	TextLocation    optional.Interface
	ForegroundColor optional.String
	BackgroundColor optional.String
	Units           optional.Interface
	Resolution      optional.Float32
	ImageHeight     optional.Float32
	ImageWidth      optional.Float32
	RotationAngle   optional.Int32
}

/*
* GenerateMultipart -  Generate barcode using POST request with parameters in multipart form.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType
* @param data String represents data to encode
* @param optional nil or *GenerateAPIGenerateMultipartOpts - Optional Parameters:
  - @param "DataType" (optional.Interface of EncodeDataType) -
  - @param "ImageFormat" (optional.Interface of BarcodeImageFormat) -
  - @param "TextLocation" (optional.Interface of CodeLocation) -
  - @param "ForegroundColor" (optional.String) -  Specify the displaying bars and content Color.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: Black.
  - @param "BackgroundColor" (optional.String) -  Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: White.
  - @param "Units" (optional.Interface of GraphicsUnit) -
  - @param "Resolution" (optional.Float32) -  Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.  Decimal separator is dot.
  - @param "ImageHeight" (optional.Float32) -  Height of the barcode image in given units. Default units: pixel.  Decimal separator is dot.
  - @param "ImageWidth" (optional.Float32) -  Width of the barcode image in given units. Default units: pixel.  Decimal separator is dot.
  - @param "RotationAngle" (optional.Int32) -  BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0.

* @return []byte
*/
func (a *GenerateAPIService) GenerateMultipart(ctx context.Context, barcodeType EncodeBarcodeType, data string, optionals *GenerateAPIGenerateMultipartOpts) ([]byte, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Post")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/generate-multipart"

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
	acceptChoices := []string{"image/png", "image/bmp", "image/gif", "image/jpeg", "image/svg+xml", "image/tiff", "application/json", "application/xml"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	formParams.Add("barcodeType", parameterToString(barcodeType, ""))
	if optionals != nil && optionals.DataType.IsSet() {
		formParams.Add("dataType", parameterToString(optionals.DataType.Value(), ""))
	}
	formParams.Add("data", parameterToString(data, ""))
	if optionals != nil && optionals.ImageFormat.IsSet() {
		formParams.Add("imageFormat", parameterToString(optionals.ImageFormat.Value(), ""))
	}
	if optionals != nil && optionals.TextLocation.IsSet() {
		formParams.Add("textLocation", parameterToString(optionals.TextLocation.Value(), ""))
	}
	if optionals != nil && optionals.ForegroundColor.IsSet() {
		formParams.Add("foregroundColor", parameterToString(optionals.ForegroundColor.Value(), ""))
	}
	if optionals != nil && optionals.BackgroundColor.IsSet() {
		formParams.Add("backgroundColor", parameterToString(optionals.BackgroundColor.Value(), ""))
	}
	if optionals != nil && optionals.Units.IsSet() {
		formParams.Add("units", parameterToString(optionals.Units.Value(), ""))
	}
	if optionals != nil && optionals.Resolution.IsSet() {
		formParams.Add("resolution", parameterToString(optionals.Resolution.Value(), ""))
	}
	if optionals != nil && optionals.ImageHeight.IsSet() {
		formParams.Add("imageHeight", parameterToString(optionals.ImageHeight.Value(), ""))
	}
	if optionals != nil && optionals.ImageWidth.IsSet() {
		formParams.Add("imageWidth", parameterToString(optionals.ImageWidth.Value(), ""))
	}
	if optionals != nil && optionals.RotationAngle.IsSet() {
		formParams.Add("rotationAngle", parameterToString(optionals.RotationAngle.Value(), ""))
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
			return responseBody, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericAPIError{
			error:      httpResponse.Status,
			text:       string(responseBody),
			StatusCode: httpResponse.StatusCode,
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
