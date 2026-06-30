package barcode

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
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
	DataType           optional.Interface
	BarcodeImageParams optional.Interface
	QrParams           optional.Interface
	Code128Params      optional.Interface
	Pdf417Params       optional.Interface
}

/*
* Generate -  Generate a barcode using a GET request with parameters in the route and query string.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType Type of barcode to generate.
* @param data String that represents the data to encode.
* @param optional nil or *GenerateAPIGenerateOpts - Optional Parameters:
  - @param "DataType" (optional.Interface of EncodeDataType) -  Type of data to encode. Default value: StringData.
  - @param "ImageFormat" (optional.Interface of BarcodeImageFormat) -  Barcode output image format. Default value: png.
  - @param "TextLocation" (optional.Interface of CodeLocation) -  Specify the displayed text location. Set to CodeLocation.None to hide CodeText. Default value depends on BarcodeType: CodeLocation.Below for 1D barcodes and CodeLocation.None for 2D barcodes.
  - @param "ForegroundColor" (optional.String) -  Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black.
  - @param "BackgroundColor" (optional.String) -  Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White.
  - @param "Units" (optional.Interface of GraphicsUnit) -  Common units for all measurements. Default units: pixels.
  - @param "Resolution" (optional.Float32) -  Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot.
  - @param "ImageHeight" (optional.Float32) -  Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
  - @param "ImageWidth" (optional.Float32) -  Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
  - @param "RotationAngle" (optional.Int32) -  Barcode image rotation angle, measured in degrees. For example, RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
  - @param "QrEncodeMode" (optional.Interface of QREncodeMode) -  QR barcode encode mode.
  - @param "QrErrorLevel" (optional.Interface of QRErrorLevel) -  QR barcode error correction level.
  - @param "QrVersion" (optional.Interface of QRVersion) -  QR barcode version. Automatically selects the smallest version that fits the data.
  - @param "QrECIEncoding" (optional.Interface of ECIEncodings) -  ECI encoding for QR barcode data.
  - @param "QrAspectRatio" (optional.Float32) -  QR barcode aspect ratio. Values: 0 to 1.
  - @param "MicroQRVersion" (optional.Interface of MicroQRVersion) -  MicroQR barcode version. Used when BarcodeType is MicroQR.
  - @param "RectMicroQrVersion" (optional.Interface of RectMicroQRVersion) -  RectMicroQR barcode version. Used when BarcodeType is RectMicroQR.
  - @param "Code128EncodeMode" (optional.Interface of Code128EncodeMode) -  Code128 barcode encode mode. Controls which Code 128 subset (A, B, C, or mix) is used.
  - @param "Pdf417EncodeMode" (optional.Interface of Pdf417EncodeMode) -  PDF417 barcode encode mode.
  - @param "Pdf417ErrorLevel" (optional.Interface of Pdf417ErrorLevel) -  PDF417 barcode error correction level.
  - @param "Pdf417Truncate" (optional.Bool) -  Whether to use truncated PDF417 format (removes right-side stop pattern).
  - @param "Pdf417Columns" (optional.Int32) -  Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto.
  - @param "Pdf417Rows" (optional.Int32) -  Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic.
  - @param "Pdf417AspectRatio" (optional.Float32) -  PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417.
  - @param "Pdf417ECIEncoding" (optional.Interface of ECIEncodings) -  ECI encoding for PDF417 barcode data.
  - @param "Pdf417IsReaderInitialization" (optional.Bool) -  Whether the barcode is used for reader initialization (programming).
  - @param "Pdf417MacroCharacters" (optional.Interface of MacroCharacter) -  Macro character to prepend (structured append).
  - @param "Pdf417IsLinked" (optional.Bool) -  Whether to use linked mode (for MicroPdf417).
  - @param "Pdf417IsCode128Emulation" (optional.Bool) -  Whether to use Code128 emulation for MicroPdf417.

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
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if imageFormatValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ImageFormat; !reflect.ValueOf(imageFormatValue).IsZero() {
			queryParams.Add("imageFormat", parameterToString(imageFormatValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if textLocationValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).TextLocation; !reflect.ValueOf(textLocationValue).IsZero() {
			queryParams.Add("textLocation", parameterToString(textLocationValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if foregroundColorValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ForegroundColor; !reflect.ValueOf(foregroundColorValue).IsZero() {
			queryParams.Add("foregroundColor", parameterToString(foregroundColorValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if backgroundColorValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).BackgroundColor; !reflect.ValueOf(backgroundColorValue).IsZero() {
			queryParams.Add("backgroundColor", parameterToString(backgroundColorValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if unitsValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).Units; !reflect.ValueOf(unitsValue).IsZero() {
			queryParams.Add("units", parameterToString(unitsValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if resolutionValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).Resolution; !reflect.ValueOf(resolutionValue).IsZero() {
			queryParams.Add("resolution", parameterToString(resolutionValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if imageHeightValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ImageHeight; !reflect.ValueOf(imageHeightValue).IsZero() {
			queryParams.Add("imageHeight", parameterToString(imageHeightValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if imageWidthValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ImageWidth; !reflect.ValueOf(imageWidthValue).IsZero() {
			queryParams.Add("imageWidth", parameterToString(imageWidthValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if rotationAngleValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).RotationAngle; !reflect.ValueOf(rotationAngleValue).IsZero() {
			queryParams.Add("rotationAngle", parameterToString(rotationAngleValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrEncodeModeValue := optionals.QrParams.Value().(QrParams).QrEncodeMode; !reflect.ValueOf(qrEncodeModeValue).IsZero() {
			queryParams.Add("qrEncodeMode", parameterToString(qrEncodeModeValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrErrorLevelValue := optionals.QrParams.Value().(QrParams).QrErrorLevel; !reflect.ValueOf(qrErrorLevelValue).IsZero() {
			queryParams.Add("qrErrorLevel", parameterToString(qrErrorLevelValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrVersionValue := optionals.QrParams.Value().(QrParams).QrVersion; !reflect.ValueOf(qrVersionValue).IsZero() {
			queryParams.Add("qrVersion", parameterToString(qrVersionValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrECIEncodingValue := optionals.QrParams.Value().(QrParams).QrECIEncoding; !reflect.ValueOf(qrECIEncodingValue).IsZero() {
			queryParams.Add("qrECIEncoding", parameterToString(qrECIEncodingValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrAspectRatioValue := optionals.QrParams.Value().(QrParams).QrAspectRatio; !reflect.ValueOf(qrAspectRatioValue).IsZero() {
			queryParams.Add("qrAspectRatio", parameterToString(qrAspectRatioValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if microQRVersionValue := optionals.QrParams.Value().(QrParams).MicroQRVersion; !reflect.ValueOf(microQRVersionValue).IsZero() {
			queryParams.Add("microQRVersion", parameterToString(microQRVersionValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if rectMicroQrVersionValue := optionals.QrParams.Value().(QrParams).RectMicroQrVersion; !reflect.ValueOf(rectMicroQrVersionValue).IsZero() {
			queryParams.Add("rectMicroQrVersion", parameterToString(rectMicroQrVersionValue, ""))
		}
	}
	if optionals != nil && optionals.Code128Params.IsSet() {
		if code128EncodeModeValue := optionals.Code128Params.Value().(Code128Params).Code128EncodeMode; !reflect.ValueOf(code128EncodeModeValue).IsZero() {
			queryParams.Add("code128EncodeMode", parameterToString(code128EncodeModeValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417EncodeModeValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417EncodeMode; !reflect.ValueOf(pdf417EncodeModeValue).IsZero() {
			queryParams.Add("pdf417EncodeMode", parameterToString(pdf417EncodeModeValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417ErrorLevelValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417ErrorLevel; !reflect.ValueOf(pdf417ErrorLevelValue).IsZero() {
			queryParams.Add("pdf417ErrorLevel", parameterToString(pdf417ErrorLevelValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417TruncateValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417Truncate; !reflect.ValueOf(pdf417TruncateValue).IsZero() {
			queryParams.Add("pdf417Truncate", parameterToString(pdf417TruncateValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417ColumnsValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417Columns; !reflect.ValueOf(pdf417ColumnsValue).IsZero() {
			queryParams.Add("pdf417Columns", parameterToString(pdf417ColumnsValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417RowsValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417Rows; !reflect.ValueOf(pdf417RowsValue).IsZero() {
			queryParams.Add("pdf417Rows", parameterToString(pdf417RowsValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417AspectRatioValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417AspectRatio; !reflect.ValueOf(pdf417AspectRatioValue).IsZero() {
			queryParams.Add("pdf417AspectRatio", parameterToString(pdf417AspectRatioValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417ECIEncodingValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417ECIEncoding; !reflect.ValueOf(pdf417ECIEncodingValue).IsZero() {
			queryParams.Add("pdf417ECIEncoding", parameterToString(pdf417ECIEncodingValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417IsReaderInitializationValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417IsReaderInitialization; !reflect.ValueOf(pdf417IsReaderInitializationValue).IsZero() {
			queryParams.Add("pdf417IsReaderInitialization", parameterToString(pdf417IsReaderInitializationValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417MacroCharactersValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417MacroCharacters; !reflect.ValueOf(pdf417MacroCharactersValue).IsZero() {
			queryParams.Add("pdf417MacroCharacters", parameterToString(pdf417MacroCharactersValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417IsLinkedValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417IsLinked; !reflect.ValueOf(pdf417IsLinkedValue).IsZero() {
			queryParams.Add("pdf417IsLinked", parameterToString(pdf417IsLinkedValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417IsCode128EmulationValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417IsCode128Emulation; !reflect.ValueOf(pdf417IsCode128EmulationValue).IsZero() {
			queryParams.Add("pdf417IsCode128Emulation", parameterToString(pdf417IsCode128EmulationValue, ""))
		}
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
		if httpResponse.StatusCode >= 400 && httpResponse.StatusCode < 500 {
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
* GenerateBody -  Generate a barcode using a POST request with parameters in the request body in JSON or XML format.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param generateParams Generation parameters.

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
		if httpResponse.StatusCode >= 400 && httpResponse.StatusCode < 500 {
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
	DataType           optional.Interface
	BarcodeImageParams optional.Interface
	QrParams           optional.Interface
	Code128Params      optional.Interface
	Pdf417Params       optional.Interface
}

/*
* GenerateMultipart -  Generate a barcode using a POST request with parameters in a multipart form.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType See https://reference.aspose.com/barcode/net/aspose.barcode.generation/encodetypes/
* @param data String that represents the data to encode.
* @param optional nil or *GenerateAPIGenerateMultipartOpts - Optional Parameters:
  - @param "DataType" (optional.Interface of EncodeDataType) -  Type of data to encode. Default value: StringData.
  - @param "ImageFormat" (optional.Interface of BarcodeImageFormat) -  Barcode output image format. Default value: png.
  - @param "TextLocation" (optional.Interface of CodeLocation) -  Specify the displayed text location. Set to CodeLocation.None to hide CodeText. Default value depends on BarcodeType: CodeLocation.Below for 1D barcodes and CodeLocation.None for 2D barcodes.
  - @param "ForegroundColor" (optional.String) -  Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black.
  - @param "BackgroundColor" (optional.String) -  Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White.
  - @param "Units" (optional.Interface of GraphicsUnit) -  Common units for all measurements. Default units: pixels.
  - @param "Resolution" (optional.Float32) -  Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot.
  - @param "ImageHeight" (optional.Float32) -  Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
  - @param "ImageWidth" (optional.Float32) -  Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
  - @param "RotationAngle" (optional.Int32) -  Barcode image rotation angle, measured in degrees. For example, RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
  - @param "QrEncodeMode" (optional.Interface of QREncodeMode) -  QR barcode encode mode.
  - @param "QrErrorLevel" (optional.Interface of QRErrorLevel) -  QR barcode error correction level.
  - @param "QrVersion" (optional.Interface of QRVersion) -  QR barcode version. Automatically selects the smallest version that fits the data.
  - @param "QrECIEncoding" (optional.Interface of ECIEncodings) -  ECI encoding for QR barcode data.
  - @param "QrAspectRatio" (optional.Float32) -  QR barcode aspect ratio. Values: 0 to 1.
  - @param "MicroQRVersion" (optional.Interface of MicroQRVersion) -  MicroQR barcode version. Used when BarcodeType is MicroQR.
  - @param "RectMicroQrVersion" (optional.Interface of RectMicroQRVersion) -  RectMicroQR barcode version. Used when BarcodeType is RectMicroQR.
  - @param "Code128EncodeMode" (optional.Interface of Code128EncodeMode) -  Code128 barcode encode mode. Controls which Code 128 subset (A, B, C, or mix) is used.
  - @param "Pdf417EncodeMode" (optional.Interface of Pdf417EncodeMode) -  PDF417 barcode encode mode.
  - @param "Pdf417ErrorLevel" (optional.Interface of Pdf417ErrorLevel) -  PDF417 barcode error correction level.
  - @param "Pdf417Truncate" (optional.Bool) -  Whether to use truncated PDF417 format (removes right-side stop pattern).
  - @param "Pdf417Columns" (optional.Int32) -  Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto.
  - @param "Pdf417Rows" (optional.Int32) -  Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic.
  - @param "Pdf417AspectRatio" (optional.Float32) -  PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417.
  - @param "Pdf417ECIEncoding" (optional.Interface of ECIEncodings) -  ECI encoding for PDF417 barcode data.
  - @param "Pdf417IsReaderInitialization" (optional.Bool) -  Whether the barcode is used for reader initialization (programming).
  - @param "Pdf417MacroCharacters" (optional.Interface of MacroCharacter) -  Macro character to prepend (structured append).
  - @param "Pdf417IsLinked" (optional.Bool) -  Whether to use linked mode (for MicroPdf417).
  - @param "Pdf417IsCode128Emulation" (optional.Bool) -  Whether to use Code128 emulation for MicroPdf417.

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
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if imageFormatValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ImageFormat; !reflect.ValueOf(imageFormatValue).IsZero() {
			formParams.Add("imageFormat", parameterToString(imageFormatValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if textLocationValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).TextLocation; !reflect.ValueOf(textLocationValue).IsZero() {
			formParams.Add("textLocation", parameterToString(textLocationValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if foregroundColorValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ForegroundColor; !reflect.ValueOf(foregroundColorValue).IsZero() {
			formParams.Add("foregroundColor", parameterToString(foregroundColorValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if backgroundColorValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).BackgroundColor; !reflect.ValueOf(backgroundColorValue).IsZero() {
			formParams.Add("backgroundColor", parameterToString(backgroundColorValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if unitsValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).Units; !reflect.ValueOf(unitsValue).IsZero() {
			formParams.Add("units", parameterToString(unitsValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if resolutionValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).Resolution; !reflect.ValueOf(resolutionValue).IsZero() {
			formParams.Add("resolution", parameterToString(resolutionValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if imageHeightValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ImageHeight; !reflect.ValueOf(imageHeightValue).IsZero() {
			formParams.Add("imageHeight", parameterToString(imageHeightValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if imageWidthValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).ImageWidth; !reflect.ValueOf(imageWidthValue).IsZero() {
			formParams.Add("imageWidth", parameterToString(imageWidthValue, ""))
		}
	}
	if optionals != nil && optionals.BarcodeImageParams.IsSet() {
		if rotationAngleValue := optionals.BarcodeImageParams.Value().(BarcodeImageParams).RotationAngle; !reflect.ValueOf(rotationAngleValue).IsZero() {
			formParams.Add("rotationAngle", parameterToString(rotationAngleValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrEncodeModeValue := optionals.QrParams.Value().(QrParams).QrEncodeMode; !reflect.ValueOf(qrEncodeModeValue).IsZero() {
			formParams.Add("qrEncodeMode", parameterToString(qrEncodeModeValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrErrorLevelValue := optionals.QrParams.Value().(QrParams).QrErrorLevel; !reflect.ValueOf(qrErrorLevelValue).IsZero() {
			formParams.Add("qrErrorLevel", parameterToString(qrErrorLevelValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrVersionValue := optionals.QrParams.Value().(QrParams).QrVersion; !reflect.ValueOf(qrVersionValue).IsZero() {
			formParams.Add("qrVersion", parameterToString(qrVersionValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrECIEncodingValue := optionals.QrParams.Value().(QrParams).QrECIEncoding; !reflect.ValueOf(qrECIEncodingValue).IsZero() {
			formParams.Add("qrECIEncoding", parameterToString(qrECIEncodingValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if qrAspectRatioValue := optionals.QrParams.Value().(QrParams).QrAspectRatio; !reflect.ValueOf(qrAspectRatioValue).IsZero() {
			formParams.Add("qrAspectRatio", parameterToString(qrAspectRatioValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if microQRVersionValue := optionals.QrParams.Value().(QrParams).MicroQRVersion; !reflect.ValueOf(microQRVersionValue).IsZero() {
			formParams.Add("microQRVersion", parameterToString(microQRVersionValue, ""))
		}
	}
	if optionals != nil && optionals.QrParams.IsSet() {
		if rectMicroQrVersionValue := optionals.QrParams.Value().(QrParams).RectMicroQrVersion; !reflect.ValueOf(rectMicroQrVersionValue).IsZero() {
			formParams.Add("rectMicroQrVersion", parameterToString(rectMicroQrVersionValue, ""))
		}
	}
	if optionals != nil && optionals.Code128Params.IsSet() {
		if code128EncodeModeValue := optionals.Code128Params.Value().(Code128Params).Code128EncodeMode; !reflect.ValueOf(code128EncodeModeValue).IsZero() {
			formParams.Add("code128EncodeMode", parameterToString(code128EncodeModeValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417EncodeModeValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417EncodeMode; !reflect.ValueOf(pdf417EncodeModeValue).IsZero() {
			formParams.Add("pdf417EncodeMode", parameterToString(pdf417EncodeModeValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417ErrorLevelValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417ErrorLevel; !reflect.ValueOf(pdf417ErrorLevelValue).IsZero() {
			formParams.Add("pdf417ErrorLevel", parameterToString(pdf417ErrorLevelValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417TruncateValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417Truncate; !reflect.ValueOf(pdf417TruncateValue).IsZero() {
			formParams.Add("pdf417Truncate", parameterToString(pdf417TruncateValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417ColumnsValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417Columns; !reflect.ValueOf(pdf417ColumnsValue).IsZero() {
			formParams.Add("pdf417Columns", parameterToString(pdf417ColumnsValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417RowsValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417Rows; !reflect.ValueOf(pdf417RowsValue).IsZero() {
			formParams.Add("pdf417Rows", parameterToString(pdf417RowsValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417AspectRatioValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417AspectRatio; !reflect.ValueOf(pdf417AspectRatioValue).IsZero() {
			formParams.Add("pdf417AspectRatio", parameterToString(pdf417AspectRatioValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417ECIEncodingValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417ECIEncoding; !reflect.ValueOf(pdf417ECIEncodingValue).IsZero() {
			formParams.Add("pdf417ECIEncoding", parameterToString(pdf417ECIEncodingValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417IsReaderInitializationValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417IsReaderInitialization; !reflect.ValueOf(pdf417IsReaderInitializationValue).IsZero() {
			formParams.Add("pdf417IsReaderInitialization", parameterToString(pdf417IsReaderInitializationValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417MacroCharactersValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417MacroCharacters; !reflect.ValueOf(pdf417MacroCharactersValue).IsZero() {
			formParams.Add("pdf417MacroCharacters", parameterToString(pdf417MacroCharactersValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417IsLinkedValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417IsLinked; !reflect.ValueOf(pdf417IsLinkedValue).IsZero() {
			formParams.Add("pdf417IsLinked", parameterToString(pdf417IsLinkedValue, ""))
		}
	}
	if optionals != nil && optionals.Pdf417Params.IsSet() {
		if pdf417IsCode128EmulationValue := optionals.Pdf417Params.Value().(Pdf417Params).Pdf417IsCode128Emulation; !reflect.ValueOf(pdf417IsCode128EmulationValue).IsZero() {
			formParams.Add("pdf417IsCode128Emulation", parameterToString(pdf417IsCode128EmulationValue, ""))
		}
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
		if httpResponse.StatusCode >= 400 && httpResponse.StatusCode < 500 {
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
