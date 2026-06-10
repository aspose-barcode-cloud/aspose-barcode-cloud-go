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

// GenerateAPIService -
type GenerateAPIService service

// GenerateAPIGenerateOpts - Optional Parameters for GenerateAPIGenerate
type GenerateAPIGenerateOpts struct {
	DataType                     optional.Interface
	ImageFormat                  optional.Interface
	TextLocation                 optional.Interface
	ForegroundColor              optional.String
	BackgroundColor              optional.String
	Units                        optional.Interface
	Resolution                   optional.Float32
	ImageHeight                  optional.Float32
	ImageWidth                   optional.Float32
	RotationAngle                optional.Int32
	QrEncodeMode                 optional.Interface
	QrErrorLevel                 optional.Interface
	QrVersion                    optional.Interface
	QrECIEncoding                optional.Interface
	QrAspectRatio                optional.Float32
	MicroQRVersion               optional.Interface
	RectMicroQrVersion           optional.Interface
	Code128EncodeMode            optional.Interface
	Pdf417EncodeMode             optional.Interface
	Pdf417ErrorLevel             optional.Interface
	Pdf417Truncate               optional.Bool
	Pdf417Columns                optional.Int32
	Pdf417Rows                   optional.Int32
	Pdf417AspectRatio            optional.Float32
	Pdf417ECIEncoding            optional.Interface
	Pdf417IsReaderInitialization optional.Bool
	Pdf417MacroCharacters        optional.Interface
	Pdf417IsLinked               optional.Bool
	Pdf417IsCode128Emulation     optional.Bool
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
	if optionals != nil && optionals.QrEncodeMode.IsSet() {
		queryParams.Add("qrEncodeMode", parameterToString(optionals.QrEncodeMode.Value(), ""))
	}
	if optionals != nil && optionals.QrErrorLevel.IsSet() {
		queryParams.Add("qrErrorLevel", parameterToString(optionals.QrErrorLevel.Value(), ""))
	}
	if optionals != nil && optionals.QrVersion.IsSet() {
		queryParams.Add("qrVersion", parameterToString(optionals.QrVersion.Value(), ""))
	}
	if optionals != nil && optionals.QrECIEncoding.IsSet() {
		queryParams.Add("qrECIEncoding", parameterToString(optionals.QrECIEncoding.Value(), ""))
	}
	if optionals != nil && optionals.QrAspectRatio.IsSet() {
		queryParams.Add("qrAspectRatio", parameterToString(optionals.QrAspectRatio.Value(), ""))
	}
	if optionals != nil && optionals.MicroQRVersion.IsSet() {
		queryParams.Add("microQRVersion", parameterToString(optionals.MicroQRVersion.Value(), ""))
	}
	if optionals != nil && optionals.RectMicroQrVersion.IsSet() {
		queryParams.Add("rectMicroQrVersion", parameterToString(optionals.RectMicroQrVersion.Value(), ""))
	}
	if optionals != nil && optionals.Code128EncodeMode.IsSet() {
		queryParams.Add("code128EncodeMode", parameterToString(optionals.Code128EncodeMode.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417EncodeMode.IsSet() {
		queryParams.Add("pdf417EncodeMode", parameterToString(optionals.Pdf417EncodeMode.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417ErrorLevel.IsSet() {
		queryParams.Add("pdf417ErrorLevel", parameterToString(optionals.Pdf417ErrorLevel.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417Truncate.IsSet() {
		queryParams.Add("pdf417Truncate", parameterToString(optionals.Pdf417Truncate.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417Columns.IsSet() {
		queryParams.Add("pdf417Columns", parameterToString(optionals.Pdf417Columns.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417Rows.IsSet() {
		queryParams.Add("pdf417Rows", parameterToString(optionals.Pdf417Rows.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417AspectRatio.IsSet() {
		queryParams.Add("pdf417AspectRatio", parameterToString(optionals.Pdf417AspectRatio.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417ECIEncoding.IsSet() {
		queryParams.Add("pdf417ECIEncoding", parameterToString(optionals.Pdf417ECIEncoding.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417IsReaderInitialization.IsSet() {
		queryParams.Add("pdf417IsReaderInitialization", parameterToString(optionals.Pdf417IsReaderInitialization.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417MacroCharacters.IsSet() {
		queryParams.Add("pdf417MacroCharacters", parameterToString(optionals.Pdf417MacroCharacters.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417IsLinked.IsSet() {
		queryParams.Add("pdf417IsLinked", parameterToString(optionals.Pdf417IsLinked.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417IsCode128Emulation.IsSet() {
		queryParams.Add("pdf417IsCode128Emulation", parameterToString(optionals.Pdf417IsCode128Emulation.Value(), ""))
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
	DataType                     optional.Interface
	ImageFormat                  optional.Interface
	TextLocation                 optional.Interface
	ForegroundColor              optional.String
	BackgroundColor              optional.String
	Units                        optional.Interface
	Resolution                   optional.Float32
	ImageHeight                  optional.Float32
	ImageWidth                   optional.Float32
	RotationAngle                optional.Int32
	QrEncodeMode                 optional.Interface
	QrErrorLevel                 optional.Interface
	QrVersion                    optional.Interface
	QrECIEncoding                optional.Interface
	QrAspectRatio                optional.Float32
	MicroQRVersion               optional.Interface
	RectMicroQrVersion           optional.Interface
	Code128EncodeMode            optional.Interface
	Pdf417EncodeMode             optional.Interface
	Pdf417ErrorLevel             optional.Interface
	Pdf417Truncate               optional.Bool
	Pdf417Columns                optional.Int32
	Pdf417Rows                   optional.Int32
	Pdf417AspectRatio            optional.Float32
	Pdf417ECIEncoding            optional.Interface
	Pdf417IsReaderInitialization optional.Bool
	Pdf417MacroCharacters        optional.Interface
	Pdf417IsLinked               optional.Bool
	Pdf417IsCode128Emulation     optional.Bool
}

/*
* GenerateMultipart -  Generate a barcode using a POST request with parameters in a multipart form.
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
* @param barcodeType
* @param data String that represents the data to encode.
* @param optional nil or *GenerateAPIGenerateMultipartOpts - Optional Parameters:
  - @param "DataType" (optional.Interface of EncodeDataType) -
  - @param "ImageFormat" (optional.Interface of BarcodeImageFormat) -
  - @param "TextLocation" (optional.Interface of CodeLocation) -
  - @param "ForegroundColor" (optional.String) -  Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black.
  - @param "BackgroundColor" (optional.String) -  Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White.
  - @param "Units" (optional.Interface of GraphicsUnit) -
  - @param "Resolution" (optional.Float32) -  Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot.
  - @param "ImageHeight" (optional.Float32) -  Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
  - @param "ImageWidth" (optional.Float32) -  Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
  - @param "RotationAngle" (optional.Int32) -  Barcode image rotation angle, measured in degrees. For example, RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
  - @param "QrEncodeMode" (optional.Interface of QREncodeMode) -
  - @param "QrErrorLevel" (optional.Interface of QRErrorLevel) -
  - @param "QrVersion" (optional.Interface of QRVersion) -
  - @param "QrECIEncoding" (optional.Interface of ECIEncodings) -
  - @param "QrAspectRatio" (optional.Float32) -  QR barcode aspect ratio. Values: 0 to 1.
  - @param "MicroQRVersion" (optional.Interface of MicroQRVersion) -
  - @param "RectMicroQrVersion" (optional.Interface of RectMicroQRVersion) -
  - @param "Code128EncodeMode" (optional.Interface of Code128EncodeMode) -
  - @param "Pdf417EncodeMode" (optional.Interface of Pdf417EncodeMode) -
  - @param "Pdf417ErrorLevel" (optional.Interface of Pdf417ErrorLevel) -
  - @param "Pdf417Truncate" (optional.Bool) -  Whether to use truncated PDF417 format (removes right-side stop pattern).
  - @param "Pdf417Columns" (optional.Int32) -  Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto.
  - @param "Pdf417Rows" (optional.Int32) -  Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic.
  - @param "Pdf417AspectRatio" (optional.Float32) -  PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417.
  - @param "Pdf417ECIEncoding" (optional.Interface of ECIEncodings) -
  - @param "Pdf417IsReaderInitialization" (optional.Bool) -  Whether the barcode is used for reader initialization (programming).
  - @param "Pdf417MacroCharacters" (optional.Interface of MacroCharacter) -
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
	if optionals != nil && optionals.QrEncodeMode.IsSet() {
		formParams.Add("qrEncodeMode", parameterToString(optionals.QrEncodeMode.Value(), ""))
	}
	if optionals != nil && optionals.QrErrorLevel.IsSet() {
		formParams.Add("qrErrorLevel", parameterToString(optionals.QrErrorLevel.Value(), ""))
	}
	if optionals != nil && optionals.QrVersion.IsSet() {
		formParams.Add("qrVersion", parameterToString(optionals.QrVersion.Value(), ""))
	}
	if optionals != nil && optionals.QrECIEncoding.IsSet() {
		formParams.Add("qrECIEncoding", parameterToString(optionals.QrECIEncoding.Value(), ""))
	}
	if optionals != nil && optionals.QrAspectRatio.IsSet() {
		formParams.Add("qrAspectRatio", parameterToString(optionals.QrAspectRatio.Value(), ""))
	}
	if optionals != nil && optionals.MicroQRVersion.IsSet() {
		formParams.Add("microQRVersion", parameterToString(optionals.MicroQRVersion.Value(), ""))
	}
	if optionals != nil && optionals.RectMicroQrVersion.IsSet() {
		formParams.Add("rectMicroQrVersion", parameterToString(optionals.RectMicroQrVersion.Value(), ""))
	}
	if optionals != nil && optionals.Code128EncodeMode.IsSet() {
		formParams.Add("code128EncodeMode", parameterToString(optionals.Code128EncodeMode.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417EncodeMode.IsSet() {
		formParams.Add("pdf417EncodeMode", parameterToString(optionals.Pdf417EncodeMode.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417ErrorLevel.IsSet() {
		formParams.Add("pdf417ErrorLevel", parameterToString(optionals.Pdf417ErrorLevel.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417Truncate.IsSet() {
		formParams.Add("pdf417Truncate", parameterToString(optionals.Pdf417Truncate.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417Columns.IsSet() {
		formParams.Add("pdf417Columns", parameterToString(optionals.Pdf417Columns.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417Rows.IsSet() {
		formParams.Add("pdf417Rows", parameterToString(optionals.Pdf417Rows.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417AspectRatio.IsSet() {
		formParams.Add("pdf417AspectRatio", parameterToString(optionals.Pdf417AspectRatio.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417ECIEncoding.IsSet() {
		formParams.Add("pdf417ECIEncoding", parameterToString(optionals.Pdf417ECIEncoding.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417IsReaderInitialization.IsSet() {
		formParams.Add("pdf417IsReaderInitialization", parameterToString(optionals.Pdf417IsReaderInitialization.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417MacroCharacters.IsSet() {
		formParams.Add("pdf417MacroCharacters", parameterToString(optionals.Pdf417MacroCharacters.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417IsLinked.IsSet() {
		formParams.Add("pdf417IsLinked", parameterToString(optionals.Pdf417IsLinked.Value(), ""))
	}
	if optionals != nil && optionals.Pdf417IsCode128Emulation.IsSet() {
		formParams.Add("pdf417IsCode128Emulation", parameterToString(optionals.Pdf417IsCode128Emulation.Value(), ""))
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
