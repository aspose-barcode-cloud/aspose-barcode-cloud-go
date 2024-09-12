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

// GenerateAPIBarcodeGenerateBarcodeTypeGetOpts - Optional Parameters for GenerateAPIBarcodeGenerateBarcodeTypeGet
type GenerateAPIBarcodeGenerateBarcodeTypeGetOpts struct {
	ImageFormat     optional.Interface
	TwoDDisplayText optional.string
	TextLocation    optional.Interface
	TextAlignment   optional.Interface
	ForegroundColor optional.string
	BackgroundColor optional.string
	Units           optional.Interface
	Resolution      optional.float32
	ImageHeight     optional.float32
	ImageWidth      optional.float32
	RotationAngle   optional.int32
}

/*
* BarcodeGenerateBarcodeTypeGet -  Generate barcode using GET request with parameters in route and query string.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType Type of barcode to generate.
* @param dataType Type of data to encode.
* @param data String represents data to encode
* @param optional nil or *GenerateAPIBarcodeGenerateBarcodeTypeGetOpts - Optional Parameters:
  - @param "ImageFormat" (optional.Interface of AvailableBarCodeImageFormat) -  Barcode output image format.  Default value: png
  - @param "TwoDDisplayText" (optional.string) -  Text that will be displayed instead of codetext in 2D barcodes.  Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode
  - @param "TextLocation" (optional.Interface of CodeLocation) -  Specify the displaying Text Location, set to CodeLocation.None to hide CodeText.  Default value: CodeLocation.Below.
  - @param "TextAlignment" (optional.Interface of TextAlignment) -  Text alignment.  Default value: TextAligment.Left
  - @param "ForegroundColor" (optional.string) -  Specify the displaying bars and content Color.   Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.Black.
  - @param "BackgroundColor" (optional.string) -  Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.White.
  - @param "Units" (optional.Interface of AvailableGraphicsUnit) -  Common Units for all measuring in query. Default units: pixel.
  - @param "Resolution" (optional.float32) -  Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.
  - @param "ImageHeight" (optional.float32) -  Height of the barcode image in given units. Default units: pixel.
  - @param "ImageWidth" (optional.float32) -  Width of the barcode image in given units. Default units: pixel.
  - @param "RotationAngle" (optional.int32) -  BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0.

* @return []byte
*/
func (a *GenerateAPIService) BarcodeGenerateBarcodeTypeGet(ctx context.Context, barcodeType EncodeBarcodeType, dataType EncodeDataType, data string, optionals *GenerateAPIBarcodeGenerateBarcodeTypeGetOpts) ([]byte, *http.Response, error) {
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

	queryParams.Add("DataType", parameterToString(dataType, ""))
	queryParams.Add("Data", parameterToString(data, ""))
	if optionals != nil && optionals.ImageFormat.IsSet() {
		queryParams.Add("ImageFormat", parameterToString(optionals.ImageFormat.Value(), ""))
	}
	if optionals != nil && optionals.TwoDDisplayText.IsSet() {
		queryParams.Add("TwoDDisplayText", parameterToString(optionals.TwoDDisplayText.Value(), ""))
	}
	if optionals != nil && optionals.TextLocation.IsSet() {
		queryParams.Add("TextLocation", parameterToString(optionals.TextLocation.Value(), ""))
	}
	if optionals != nil && optionals.TextAlignment.IsSet() {
		queryParams.Add("TextAlignment", parameterToString(optionals.TextAlignment.Value(), ""))
	}
	if optionals != nil && optionals.ForegroundColor.IsSet() {
		queryParams.Add("ForegroundColor", parameterToString(optionals.ForegroundColor.Value(), ""))
	}
	if optionals != nil && optionals.BackgroundColor.IsSet() {
		queryParams.Add("BackgroundColor", parameterToString(optionals.BackgroundColor.Value(), ""))
	}
	if optionals != nil && optionals.Units.IsSet() {
		queryParams.Add("Units", parameterToString(optionals.Units.Value(), ""))
	}
	if optionals != nil && optionals.Resolution.IsSet() {
		queryParams.Add("Resolution", parameterToString(optionals.Resolution.Value(), ""))
	}
	if optionals != nil && optionals.ImageHeight.IsSet() {
		queryParams.Add("ImageHeight", parameterToString(optionals.ImageHeight.Value(), ""))
	}
	if optionals != nil && optionals.ImageWidth.IsSet() {
		queryParams.Add("ImageWidth", parameterToString(optionals.ImageWidth.Value(), ""))
	}
	if optionals != nil && optionals.RotationAngle.IsSet() {
		queryParams.Add("RotationAngle", parameterToString(optionals.RotationAngle.Value(), ""))
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
* BarcodeGenerateBodyPost -  Generate barcode using POST request with parameters in body in json or xml format.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param generateParams Parameters of generation

* @return []byte
 */
func (a *GenerateAPIService) BarcodeGenerateBodyPost(ctx context.Context, generateParams GenerateParams) ([]byte, *http.Response, error) {
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

// GenerateAPIBarcodeGenerateFormPostOpts - Optional Parameters for GenerateAPIBarcodeGenerateFormPost
type GenerateAPIBarcodeGenerateFormPostOpts struct {
	ImageFormat     optional.Interface
	TwoDDisplayText optional.string
	TextLocation    optional.Interface
	TextAlignment   optional.Interface
	ForegroundColor optional.string
	BackgroundColor optional.string
	Units           optional.Interface
	Resolution      optional.float32
	ImageHeight     optional.float32
	ImageWidth      optional.float32
	RotationAngle   optional.int32
}

/*
* BarcodeGenerateFormPost -  Generate barcode using POST request with parameters in url ecncoded form.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType
* @param dataType
* @param data String represents data to encode
* @param optional nil or *GenerateAPIBarcodeGenerateFormPostOpts - Optional Parameters:
  - @param "ImageFormat" (optional.Interface of AvailableBarCodeImageFormat) -
  - @param "TwoDDisplayText" (optional.string) -  Text that will be displayed instead of codetext in 2D barcodes.  Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode
  - @param "TextLocation" (optional.Interface of CodeLocation) -
  - @param "TextAlignment" (optional.Interface of TextAlignment) -
  - @param "ForegroundColor" (optional.string) -  Specify the displaying bars and content Color.   Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.Black.
  - @param "BackgroundColor" (optional.string) -  Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.   For example: Color.AliceBlue or #FF000000  Default value: Color.White.
  - @param "Units" (optional.Interface of AvailableGraphicsUnit) -
  - @param "Resolution" (optional.float32) -  Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.
  - @param "ImageHeight" (optional.float32) -  Height of the barcode image in given units. Default units: pixel.
  - @param "ImageWidth" (optional.float32) -  Width of the barcode image in given units. Default units: pixel.
  - @param "RotationAngle" (optional.int32) -  BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0.

* @return []byte
*/
func (a *GenerateAPIService) BarcodeGenerateFormPost(ctx context.Context, barcodeType EncodeBarcodeType, dataType EncodeDataType, data string, optionals *GenerateAPIBarcodeGenerateFormPostOpts) ([]byte, *http.Response, error) {
	var (
		httpMethod    = strings.ToUpper("Post")
		postBody      interface{}
		fileName      string
		fileFieldName string
		fileBytes     []byte
		returnValue   []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/generate-form"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypeChoices := []string{"application/x-www-form-urlencoded"}

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
	formParams.Add("DataType", parameterToString(dataType, ""))
	formParams.Add("Data", parameterToString(data, ""))
	if optionals != nil && optionals.ImageFormat.IsSet() {
		formParams.Add("ImageFormat", parameterToString(optionals.ImageFormat.Value(), ""))
	}
	if optionals != nil && optionals.TwoDDisplayText.IsSet() {
		formParams.Add("TwoDDisplayText", parameterToString(optionals.TwoDDisplayText.Value(), ""))
	}
	if optionals != nil && optionals.TextLocation.IsSet() {
		formParams.Add("TextLocation", parameterToString(optionals.TextLocation.Value(), ""))
	}
	if optionals != nil && optionals.TextAlignment.IsSet() {
		formParams.Add("TextAlignment", parameterToString(optionals.TextAlignment.Value(), ""))
	}
	if optionals != nil && optionals.ForegroundColor.IsSet() {
		formParams.Add("ForegroundColor", parameterToString(optionals.ForegroundColor.Value(), ""))
	}
	if optionals != nil && optionals.BackgroundColor.IsSet() {
		formParams.Add("BackgroundColor", parameterToString(optionals.BackgroundColor.Value(), ""))
	}
	if optionals != nil && optionals.Units.IsSet() {
		formParams.Add("Units", parameterToString(optionals.Units.Value(), ""))
	}
	if optionals != nil && optionals.Resolution.IsSet() {
		formParams.Add("Resolution", parameterToString(optionals.Resolution.Value(), ""))
	}
	if optionals != nil && optionals.ImageHeight.IsSet() {
		formParams.Add("ImageHeight", parameterToString(optionals.ImageHeight.Value(), ""))
	}
	if optionals != nil && optionals.ImageWidth.IsSet() {
		formParams.Add("ImageWidth", parameterToString(optionals.ImageWidth.Value(), ""))
	}
	if optionals != nil && optionals.RotationAngle.IsSet() {
		formParams.Add("RotationAngle", parameterToString(optionals.RotationAngle.Value(), ""))
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
