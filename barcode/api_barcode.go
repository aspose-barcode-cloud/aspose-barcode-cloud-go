/*
 * MIT License

 * Copyright (c) 2022 Aspose Pty Ltd

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

//BarcodeApiService -
type BarcodeApiService service

//BarcodeApiGetBarcodeGenerateOpts - Optional Parameters for BarcodeApiGetBarcodeGenerate
type BarcodeApiGetBarcodeGenerateOpts struct {
	TwoDDisplayText    optional.String
	TextLocation       optional.String
	TextAlignment      optional.String
	TextColor          optional.String
	FontSizeMode       optional.String
	NoWrap             optional.Bool
	Resolution         optional.Float64
	ResolutionX        optional.Float64
	ResolutionY        optional.Float64
	DimensionX         optional.Float64
	TextSpace          optional.Float64
	Units              optional.String
	SizeMode           optional.String
	BarHeight          optional.Float64
	ImageHeight        optional.Float64
	ImageWidth         optional.Float64
	RotationAngle      optional.Float64
	BackColor          optional.String
	BarColor           optional.String
	BorderColor        optional.String
	BorderWidth        optional.Float64
	BorderDashStyle    optional.String
	BorderVisible      optional.Bool
	EnableChecksum     optional.String
	EnableEscape       optional.Bool
	FilledBars         optional.Bool
	AlwaysShowChecksum optional.Bool
	WideNarrowRatio    optional.Float64
	ValidateText       optional.Bool
	SupplementData     optional.String
	SupplementSpace    optional.Float64
	BarWidthReduction  optional.Float64
	Format             optional.String
}

/*
 * GetBarcodeGenerate -  Generate barcode.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param type_ Type of barcode to generate.
 * @param text Text to encode.
 * @param optional nil or *BarcodeApiGetBarcodeGenerateOpts - Optional Parameters:
     * @param "TwoDDisplayText" (optional.String) -  Text that will be displayed instead of codetext in 2D barcodes. Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode
     * @param "TextLocation" (optional.String) -  Specify the displaying Text Location, set to CodeLocation.None to hide CodeText. Default value: CodeLocation.Below.
     * @param "TextAlignment" (optional.String) -  Text alignment.
     * @param "TextColor" (optional.String) -  Specify the displaying CodeText&#39;s Color. Default value: Color.Black.
     * @param "FontSizeMode" (optional.String) -  Specify FontSizeMode. If FontSizeMode is set to Auto, font size will be calculated automatically based on xDimension value. It is recommended to use FontSizeMode.Auto especially in AutoSizeMode.Nearest or AutoSizeMode.Interpolation. Default value: FontSizeMode.Auto.
     * @param "NoWrap" (optional.Bool) -  Specify word wraps (line breaks) within text. Default value: false.
     * @param "Resolution" (optional.Float64) -  Resolution of the BarCode image. One value for both dimensions. Default value: 96 dpi.
     * @param "ResolutionX" (optional.Float64) -  DEPRECATED: Use &#39;Resolution&#39; instead.
     * @param "ResolutionY" (optional.Float64) -  DEPRECATED: Use &#39;Resolution&#39; instead.
     * @param "DimensionX" (optional.Float64) -  The smallest width of the unit of BarCode bars or spaces. Increase this will increase the whole barcode image width. Ignored if AutoSizeMode property is set to AutoSizeMode.Nearest or AutoSizeMode.Interpolation.
     * @param "TextSpace" (optional.Float64) -  Space between the CodeText and the BarCode in Unit value. Default value: 2pt. Ignored for EAN8, EAN13, UPCE, UPCA, ISBN, ISMN, ISSN, UpcaGs1DatabarCoupon.
     * @param "Units" (optional.String) -  Common Units for all measuring in query. Default units: pixel.
     * @param "SizeMode" (optional.String) -  Specifies the different types of automatic sizing modes. Default value: AutoSizeMode.None.
     * @param "BarHeight" (optional.Float64) -  Height of the barcode in given units. Default units: pixel.
     * @param "ImageHeight" (optional.Float64) -  Height of the barcode image in given units. Default units: pixel.
     * @param "ImageWidth" (optional.Float64) -  Width of the barcode image in given units. Default units: pixel.
     * @param "RotationAngle" (optional.Float64) -  BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
     * @param "BackColor" (optional.String) -  Background color of the barcode image. Default value: Color.White.
     * @param "BarColor" (optional.String) -  Bars color. Default value: Color.Black.
     * @param "BorderColor" (optional.String) -  Border color. Default value: Color.Black.
     * @param "BorderWidth" (optional.Float64) -  Border width. Default value: 0. Ignored if Visible is set to false.
     * @param "BorderDashStyle" (optional.String) -  Border dash style. Default value: BorderDashStyle.Solid.
     * @param "BorderVisible" (optional.Bool) -  Border visibility. If false than parameter Width is always ignored (0). Default value: false.
     * @param "EnableChecksum" (optional.String) -  Enable checksum during generation 1D barcodes. Default is treated as Yes for symbology which must contain checksum, as No where checksum only possible. Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN, Codabar Checksum always used: Rest symbology
     * @param "EnableEscape" (optional.Bool) -  Indicates whether explains the character \&quot;\\\&quot; as an escape character in CodeText property. Used for Pdf417, DataMatrix, Code128 only If the EnableEscape is true, \&quot;\\\&quot; will be explained as a special escape character. Otherwise, \&quot;\\\&quot; acts as normal characters. Aspose.BarCode supports input decimal ascii code and mnemonic for ASCII control-code characters. For example, \\013 and \\\\CR stands for CR.
     * @param "FilledBars" (optional.Bool) -  Value indicating whether bars are filled. Only for 1D barcodes. Default value: true.
     * @param "AlwaysShowChecksum" (optional.Bool) -  Always display checksum digit in the human readable text for Code128 and GS1Code128 barcodes.
     * @param "WideNarrowRatio" (optional.Float64) -  Wide bars to Narrow bars ratio. Default value: 3, that is, wide bars are 3 times as wide as narrow bars. Used for ITF, PZN, PharmaCode, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, IATA2of5, VIN, DeutschePost, OPC, Code32, DataLogic2of5, PatchCode, Code39Extended, Code39Standard
     * @param "ValidateText" (optional.Bool) -  Only for 1D barcodes. If codetext is incorrect and value set to true - exception will be thrown. Otherwise codetext will be corrected to match barcode&#39;s specification. Exception always will be thrown for: Databar symbology if codetext is incorrect. Exception always will not be thrown for: AustraliaPost, SingaporePost, Code39Extended, Code93Extended, Code16K, Code128 symbology if codetext is incorrect.
     * @param "SupplementData" (optional.String) -  Supplement parameters. Used for Interleaved2of5, Standard2of5, EAN13, EAN8, UPCA, UPCE, ISBN, ISSN, ISMN.
     * @param "SupplementSpace" (optional.Float64) -  Space between main the BarCode and supplement BarCode.
     * @param "BarWidthReduction" (optional.Float64) -  Bars reduction value that is used to compensate ink spread while printing.
     * @param "Format" (optional.String) -  Result image format.

 * @return []byte
*/
func (a *BarcodeApiService) GetBarcodeGenerate(ctx context.Context, type_ string, text string, optionals *BarcodeApiGetBarcodeGenerateOpts) ([]byte, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/generate"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("Type", parameterToString(type_, ""))
	queryParams.Add("Text", parameterToString(text, ""))
	if optionals != nil && optionals.TwoDDisplayText.IsSet() {
		queryParams.Add("TwoDDisplayText", parameterToString(optionals.TwoDDisplayText.Value(), ""))
	}
	if optionals != nil && optionals.TextLocation.IsSet() {
		queryParams.Add("TextLocation", parameterToString(optionals.TextLocation.Value(), ""))
	}
	if optionals != nil && optionals.TextAlignment.IsSet() {
		queryParams.Add("TextAlignment", parameterToString(optionals.TextAlignment.Value(), ""))
	}
	if optionals != nil && optionals.TextColor.IsSet() {
		queryParams.Add("TextColor", parameterToString(optionals.TextColor.Value(), ""))
	}
	if optionals != nil && optionals.FontSizeMode.IsSet() {
		queryParams.Add("FontSizeMode", parameterToString(optionals.FontSizeMode.Value(), ""))
	}
	if optionals != nil && optionals.NoWrap.IsSet() {
		queryParams.Add("NoWrap", parameterToString(optionals.NoWrap.Value(), ""))
	}
	if optionals != nil && optionals.Resolution.IsSet() {
		queryParams.Add("Resolution", parameterToString(optionals.Resolution.Value(), ""))
	}
	if optionals != nil && optionals.ResolutionX.IsSet() {
		queryParams.Add("ResolutionX", parameterToString(optionals.ResolutionX.Value(), ""))
	}
	if optionals != nil && optionals.ResolutionY.IsSet() {
		queryParams.Add("ResolutionY", parameterToString(optionals.ResolutionY.Value(), ""))
	}
	if optionals != nil && optionals.DimensionX.IsSet() {
		queryParams.Add("DimensionX", parameterToString(optionals.DimensionX.Value(), ""))
	}
	if optionals != nil && optionals.TextSpace.IsSet() {
		queryParams.Add("TextSpace", parameterToString(optionals.TextSpace.Value(), ""))
	}
	if optionals != nil && optionals.Units.IsSet() {
		queryParams.Add("Units", parameterToString(optionals.Units.Value(), ""))
	}
	if optionals != nil && optionals.SizeMode.IsSet() {
		queryParams.Add("SizeMode", parameterToString(optionals.SizeMode.Value(), ""))
	}
	if optionals != nil && optionals.BarHeight.IsSet() {
		queryParams.Add("BarHeight", parameterToString(optionals.BarHeight.Value(), ""))
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
	if optionals != nil && optionals.BackColor.IsSet() {
		queryParams.Add("BackColor", parameterToString(optionals.BackColor.Value(), ""))
	}
	if optionals != nil && optionals.BarColor.IsSet() {
		queryParams.Add("BarColor", parameterToString(optionals.BarColor.Value(), ""))
	}
	if optionals != nil && optionals.BorderColor.IsSet() {
		queryParams.Add("BorderColor", parameterToString(optionals.BorderColor.Value(), ""))
	}
	if optionals != nil && optionals.BorderWidth.IsSet() {
		queryParams.Add("BorderWidth", parameterToString(optionals.BorderWidth.Value(), ""))
	}
	if optionals != nil && optionals.BorderDashStyle.IsSet() {
		queryParams.Add("BorderDashStyle", parameterToString(optionals.BorderDashStyle.Value(), ""))
	}
	if optionals != nil && optionals.BorderVisible.IsSet() {
		queryParams.Add("BorderVisible", parameterToString(optionals.BorderVisible.Value(), ""))
	}
	if optionals != nil && optionals.EnableChecksum.IsSet() {
		queryParams.Add("EnableChecksum", parameterToString(optionals.EnableChecksum.Value(), ""))
	}
	if optionals != nil && optionals.EnableEscape.IsSet() {
		queryParams.Add("EnableEscape", parameterToString(optionals.EnableEscape.Value(), ""))
	}
	if optionals != nil && optionals.FilledBars.IsSet() {
		queryParams.Add("FilledBars", parameterToString(optionals.FilledBars.Value(), ""))
	}
	if optionals != nil && optionals.AlwaysShowChecksum.IsSet() {
		queryParams.Add("AlwaysShowChecksum", parameterToString(optionals.AlwaysShowChecksum.Value(), ""))
	}
	if optionals != nil && optionals.WideNarrowRatio.IsSet() {
		queryParams.Add("WideNarrowRatio", parameterToString(optionals.WideNarrowRatio.Value(), ""))
	}
	if optionals != nil && optionals.ValidateText.IsSet() {
		queryParams.Add("ValidateText", parameterToString(optionals.ValidateText.Value(), ""))
	}
	if optionals != nil && optionals.SupplementData.IsSet() {
		queryParams.Add("SupplementData", parameterToString(optionals.SupplementData.Value(), ""))
	}
	if optionals != nil && optionals.SupplementSpace.IsSet() {
		queryParams.Add("SupplementSpace", parameterToString(optionals.SupplementSpace.Value(), ""))
	}
	if optionals != nil && optionals.BarWidthReduction.IsSet() {
		queryParams.Add("BarWidthReduction", parameterToString(optionals.BarWidthReduction.Value(), ""))
	}
	if optionals != nil && optionals.Format.IsSet() {
		queryParams.Add("format", parameterToString(optionals.Format.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"image/png", "image/bmp", "image/gif", "image/jpeg", "image/svg+xml", "image/tiff"}

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

//BarcodeApiGetBarcodeRecognizeOpts - Optional Parameters for BarcodeApiGetBarcodeRecognize
type BarcodeApiGetBarcodeRecognizeOpts struct {
	Type_                                optional.String
	ChecksumValidation                   optional.String
	DetectEncoding                       optional.Bool
	Preset                               optional.String
	RectX                                optional.Int32
	RectY                                optional.Int32
	RectWidth                            optional.Int32
	RectHeight                           optional.Int32
	StripFNC                             optional.Bool
	Timeout                              optional.Int32
	MedianSmoothingWindowSize            optional.Int32
	AllowMedianSmoothing                 optional.Bool
	AllowComplexBackground               optional.Bool
	AllowDatamatrixIndustrialBarcodes    optional.Bool
	AllowDecreasedImage                  optional.Bool
	AllowDetectScanGap                   optional.Bool
	AllowIncorrectBarcodes               optional.Bool
	AllowInvertImage                     optional.Bool
	AllowMicroWhiteSpotsRemoving         optional.Bool
	AllowOneDFastBarcodesDetector        optional.Bool
	AllowOneDWipedBarsRestoration        optional.Bool
	AllowQRMicroQrRestoration            optional.Bool
	AllowRegularImage                    optional.Bool
	AllowSaltAndPepperFiltering          optional.Bool
	AllowWhiteSpotsRemoving              optional.Bool
	CheckMore1DVariants                  optional.Bool
	FastScanOnly                         optional.Bool
	RegionLikelihoodThresholdPercent     optional.Float64
	ScanWindowSizes                      optional.Interface
	Similarity                           optional.Float64
	SkipDiagonalSearch                   optional.Bool
	ReadTinyBarcodes                     optional.Bool
	AustralianPostEncodingTable          optional.String
	IgnoreEndingFillingPatternsForCTable optional.Bool
	RectangleRegion                      optional.String
	Storage                              optional.String
	Folder                               optional.String
}

/*
 * GetBarcodeRecognize -  Recognize barcode from a file on server.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The image file name.
 * @param optional nil or *BarcodeApiGetBarcodeRecognizeOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  The type of barcode to read.
     * @param "ChecksumValidation" (optional.String) -  Enable checksum validation during recognition for 1D barcodes. Default is treated as Yes for symbologies which must contain checksum, as No where checksum only possible. Checksum never used: Codabar Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN Checksum always used: Rest symbologies
     * @param "DetectEncoding" (optional.Bool) -  A flag which force engine to detect codetext encoding for Unicode.
     * @param "Preset" (optional.String) -  Preset allows to configure recognition quality and speed manually. You can quickly set up Preset by embedded presets: HighPerformance, NormalQuality, HighQuality, MaxBarCodes or you can manually configure separate options. Default value of Preset is NormalQuality.
     * @param "RectX" (optional.Int32) -  Set X for area for recognition.
     * @param "RectY" (optional.Int32) -  Set Y for area for recognition.
     * @param "RectWidth" (optional.Int32) -  Set Width of area for recognition.
     * @param "RectHeight" (optional.Int32) -  Set Height of area for recognition.
     * @param "StripFNC" (optional.Bool) -  Value indicating whether FNC symbol strip must be done.
     * @param "Timeout" (optional.Int32) -  Timeout of recognition process.
     * @param "MedianSmoothingWindowSize" (optional.Int32) -  Window size for median smoothing. Typical values are 3 or 4. Default value is 3. AllowMedianSmoothing must be set.
     * @param "AllowMedianSmoothing" (optional.Bool) -  Allows engine to enable median smoothing as additional scan. Mode helps to recognize noised barcodes.
     * @param "AllowComplexBackground" (optional.Bool) -  Allows engine to recognize color barcodes on color background as additional scan. Extremely slow mode.
     * @param "AllowDatamatrixIndustrialBarcodes" (optional.Bool) -  Allows engine for Datamatrix to recognize dashed industrial Datamatrix barcodes. Slow mode which helps only for dashed barcodes which consist from spots.
     * @param "AllowDecreasedImage" (optional.Bool) -  Allows engine to recognize decreased image as additional scan. Size for decreasing is selected by internal engine algorithms. Mode helps to recognize barcodes which are noised and blurred but captured with high resolution.
     * @param "AllowDetectScanGap" (optional.Bool) -  Allows engine to use gap between scans to increase recognition speed. Mode can make recognition problems with low height barcodes.
     * @param "AllowIncorrectBarcodes" (optional.Bool) -  Allows engine to recognize barcodes which has incorrect checksum or incorrect values. Mode can be used to recognize damaged barcodes with incorrect text.
     * @param "AllowInvertImage" (optional.Bool) -  Allows engine to recognize inverse color image as additional scan. Mode can be used when barcode is white on black background.
     * @param "AllowMicroWhiteSpotsRemoving" (optional.Bool) -  Allows engine for Postal barcodes to recognize slightly noised images. Mode helps to recognize slightly damaged Postal barcodes.
     * @param "AllowOneDFastBarcodesDetector" (optional.Bool) -  Allows engine for 1D barcodes to quickly recognize high quality barcodes which fill almost whole image. Mode helps to quickly recognize generated barcodes from Internet.
     * @param "AllowOneDWipedBarsRestoration" (optional.Bool) -  Allows engine for 1D barcodes to recognize barcodes with single wiped/glued bars in pattern.
     * @param "AllowQRMicroQrRestoration" (optional.Bool) -  Allows engine for QR/MicroQR to recognize damaged MicroQR barcodes.
     * @param "AllowRegularImage" (optional.Bool) -  Allows engine to recognize regular image without any restorations as main scan. Mode to recognize image as is.
     * @param "AllowSaltAndPepperFiltering" (optional.Bool) -  Allows engine to recognize barcodes with salt and pepper noise type. Mode can remove small noise with white and black dots.
     * @param "AllowWhiteSpotsRemoving" (optional.Bool) -  Allows engine to recognize image without small white spots as additional scan. Mode helps to recognize noised image as well as median smoothing filtering.
     * @param "CheckMore1DVariants" (optional.Bool) -  Allows engine to recognize 1D barcodes with checksum by checking more recognition variants. Default value: False.
     * @param "FastScanOnly" (optional.Bool) -  Allows engine for 1D barcodes to quickly recognize middle slice of an image and return result without using any time-consuming algorithms. Default value: False.
     * @param "RegionLikelihoodThresholdPercent" (optional.Float64) -  Sets threshold for detected regions that may contain barcodes. Value 0.7 means that bottom 70% of possible regions are filtered out and not processed further. Region likelihood threshold must be between [0.05, 0.9] Use high values for clear images with few barcodes. Use low values for images with many barcodes or for noisy images. Low value may lead to a bigger recognition time.
     * @param "ScanWindowSizes" (optional.Interface of []int32) -  Scan window sizes in pixels. Allowed sizes are 10, 15, 20, 25, 30. Scanning with small window size takes more time and provides more accuracy but may fail in detecting very big barcodes. Combining of several window sizes can improve detection quality.
     * @param "Similarity" (optional.Float64) -  Similarity coefficient depends on how homogeneous barcodes are. Use high value for for clear barcodes. Use low values to detect barcodes that ara partly damaged or not lighten evenly. Similarity coefficient must be between [0.5, 0.9]
     * @param "SkipDiagonalSearch" (optional.Bool) -  Allows detector to skip search for diagonal barcodes. Setting it to false will increase detection time but allow to find diagonal barcodes that can be missed otherwise. Enabling of diagonal search leads to a bigger detection time.
     * @param "ReadTinyBarcodes" (optional.Bool) -  Allows engine to recognize tiny barcodes on large images. Ignored if AllowIncorrectBarcodes is set to True. Default value: False.
     * @param "AustralianPostEncodingTable" (optional.String) -  Interpreting Type for the Customer Information of AustralianPost BarCode.Default is CustomerInformationInterpretingType.Other.
     * @param "IgnoreEndingFillingPatternsForCTable" (optional.Bool) -  The flag which force AustraliaPost decoder to ignore last filling patterns in Customer Information Field during decoding as CTable method. CTable encoding method does not have any gaps in encoding table and sequnce \&quot;333\&quot; of filling paterns is decoded as letter \&quot;z\&quot;.
     * @param "RectangleRegion" (optional.String) -
     * @param "Storage" (optional.String) -  The image storage.
     * @param "Folder" (optional.String) -  The image folder.

 * @return BarcodeResponseList
*/
func (a *BarcodeApiService) GetBarcodeRecognize(ctx context.Context, name string, optionals *BarcodeApiGetBarcodeRecognizeOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/{name}/recognize"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Type_.IsSet() {
		queryParams.Add("Type", parameterToString(optionals.Type_.Value(), ""))
	}
	if optionals != nil && optionals.ChecksumValidation.IsSet() {
		queryParams.Add("ChecksumValidation", parameterToString(optionals.ChecksumValidation.Value(), ""))
	}
	if optionals != nil && optionals.DetectEncoding.IsSet() {
		queryParams.Add("DetectEncoding", parameterToString(optionals.DetectEncoding.Value(), ""))
	}
	if optionals != nil && optionals.Preset.IsSet() {
		queryParams.Add("Preset", parameterToString(optionals.Preset.Value(), ""))
	}
	if optionals != nil && optionals.RectX.IsSet() {
		queryParams.Add("RectX", parameterToString(optionals.RectX.Value(), ""))
	}
	if optionals != nil && optionals.RectY.IsSet() {
		queryParams.Add("RectY", parameterToString(optionals.RectY.Value(), ""))
	}
	if optionals != nil && optionals.RectWidth.IsSet() {
		queryParams.Add("RectWidth", parameterToString(optionals.RectWidth.Value(), ""))
	}
	if optionals != nil && optionals.RectHeight.IsSet() {
		queryParams.Add("RectHeight", parameterToString(optionals.RectHeight.Value(), ""))
	}
	if optionals != nil && optionals.StripFNC.IsSet() {
		queryParams.Add("StripFNC", parameterToString(optionals.StripFNC.Value(), ""))
	}
	if optionals != nil && optionals.Timeout.IsSet() {
		queryParams.Add("Timeout", parameterToString(optionals.Timeout.Value(), ""))
	}
	if optionals != nil && optionals.MedianSmoothingWindowSize.IsSet() {
		queryParams.Add("MedianSmoothingWindowSize", parameterToString(optionals.MedianSmoothingWindowSize.Value(), ""))
	}
	if optionals != nil && optionals.AllowMedianSmoothing.IsSet() {
		queryParams.Add("AllowMedianSmoothing", parameterToString(optionals.AllowMedianSmoothing.Value(), ""))
	}
	if optionals != nil && optionals.AllowComplexBackground.IsSet() {
		queryParams.Add("AllowComplexBackground", parameterToString(optionals.AllowComplexBackground.Value(), ""))
	}
	if optionals != nil && optionals.AllowDatamatrixIndustrialBarcodes.IsSet() {
		queryParams.Add("AllowDatamatrixIndustrialBarcodes", parameterToString(optionals.AllowDatamatrixIndustrialBarcodes.Value(), ""))
	}
	if optionals != nil && optionals.AllowDecreasedImage.IsSet() {
		queryParams.Add("AllowDecreasedImage", parameterToString(optionals.AllowDecreasedImage.Value(), ""))
	}
	if optionals != nil && optionals.AllowDetectScanGap.IsSet() {
		queryParams.Add("AllowDetectScanGap", parameterToString(optionals.AllowDetectScanGap.Value(), ""))
	}
	if optionals != nil && optionals.AllowIncorrectBarcodes.IsSet() {
		queryParams.Add("AllowIncorrectBarcodes", parameterToString(optionals.AllowIncorrectBarcodes.Value(), ""))
	}
	if optionals != nil && optionals.AllowInvertImage.IsSet() {
		queryParams.Add("AllowInvertImage", parameterToString(optionals.AllowInvertImage.Value(), ""))
	}
	if optionals != nil && optionals.AllowMicroWhiteSpotsRemoving.IsSet() {
		queryParams.Add("AllowMicroWhiteSpotsRemoving", parameterToString(optionals.AllowMicroWhiteSpotsRemoving.Value(), ""))
	}
	if optionals != nil && optionals.AllowOneDFastBarcodesDetector.IsSet() {
		queryParams.Add("AllowOneDFastBarcodesDetector", parameterToString(optionals.AllowOneDFastBarcodesDetector.Value(), ""))
	}
	if optionals != nil && optionals.AllowOneDWipedBarsRestoration.IsSet() {
		queryParams.Add("AllowOneDWipedBarsRestoration", parameterToString(optionals.AllowOneDWipedBarsRestoration.Value(), ""))
	}
	if optionals != nil && optionals.AllowQRMicroQrRestoration.IsSet() {
		queryParams.Add("AllowQRMicroQrRestoration", parameterToString(optionals.AllowQRMicroQrRestoration.Value(), ""))
	}
	if optionals != nil && optionals.AllowRegularImage.IsSet() {
		queryParams.Add("AllowRegularImage", parameterToString(optionals.AllowRegularImage.Value(), ""))
	}
	if optionals != nil && optionals.AllowSaltAndPepperFiltering.IsSet() {
		queryParams.Add("AllowSaltAndPepperFiltering", parameterToString(optionals.AllowSaltAndPepperFiltering.Value(), ""))
	}
	if optionals != nil && optionals.AllowWhiteSpotsRemoving.IsSet() {
		queryParams.Add("AllowWhiteSpotsRemoving", parameterToString(optionals.AllowWhiteSpotsRemoving.Value(), ""))
	}
	if optionals != nil && optionals.CheckMore1DVariants.IsSet() {
		queryParams.Add("CheckMore1DVariants", parameterToString(optionals.CheckMore1DVariants.Value(), ""))
	}
	if optionals != nil && optionals.FastScanOnly.IsSet() {
		queryParams.Add("FastScanOnly", parameterToString(optionals.FastScanOnly.Value(), ""))
	}
	if optionals != nil && optionals.RegionLikelihoodThresholdPercent.IsSet() {
		queryParams.Add("RegionLikelihoodThresholdPercent", parameterToString(optionals.RegionLikelihoodThresholdPercent.Value(), ""))
	}
	if optionals != nil && optionals.ScanWindowSizes.IsSet() {
		queryParams.Add("ScanWindowSizes", parameterToString(optionals.ScanWindowSizes.Value(), "multi"))
	}
	if optionals != nil && optionals.Similarity.IsSet() {
		queryParams.Add("Similarity", parameterToString(optionals.Similarity.Value(), ""))
	}
	if optionals != nil && optionals.SkipDiagonalSearch.IsSet() {
		queryParams.Add("SkipDiagonalSearch", parameterToString(optionals.SkipDiagonalSearch.Value(), ""))
	}
	if optionals != nil && optionals.ReadTinyBarcodes.IsSet() {
		queryParams.Add("ReadTinyBarcodes", parameterToString(optionals.ReadTinyBarcodes.Value(), ""))
	}
	if optionals != nil && optionals.AustralianPostEncodingTable.IsSet() {
		queryParams.Add("AustralianPostEncodingTable", parameterToString(optionals.AustralianPostEncodingTable.Value(), ""))
	}
	if optionals != nil && optionals.IgnoreEndingFillingPatternsForCTable.IsSet() {
		queryParams.Add("IgnoreEndingFillingPatternsForCTable", parameterToString(optionals.IgnoreEndingFillingPatternsForCTable.Value(), ""))
	}
	if optionals != nil && optionals.RectangleRegion.IsSet() {
		queryParams.Add("RectangleRegion", parameterToString(optionals.RectangleRegion.Value(), ""))
	}
	if optionals != nil && optionals.Storage.IsSet() {
		queryParams.Add("storage", parameterToString(optionals.Storage.Value(), ""))
	}
	if optionals != nil && optionals.Folder.IsSet() {
		queryParams.Add("folder", parameterToString(optionals.Folder.Value(), ""))
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
			error: httpResponse.Status,
			text:  string(responseBody),
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

		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, err
}

//BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts - Optional Parameters for BarcodeApiPostBarcodeRecognizeFromUrlOrContent
type BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts struct {
	Type_                                optional.String
	ChecksumValidation                   optional.String
	DetectEncoding                       optional.Bool
	Preset                               optional.String
	RectX                                optional.Int32
	RectY                                optional.Int32
	RectWidth                            optional.Int32
	RectHeight                           optional.Int32
	StripFNC                             optional.Bool
	Timeout                              optional.Int32
	MedianSmoothingWindowSize            optional.Int32
	AllowMedianSmoothing                 optional.Bool
	AllowComplexBackground               optional.Bool
	AllowDatamatrixIndustrialBarcodes    optional.Bool
	AllowDecreasedImage                  optional.Bool
	AllowDetectScanGap                   optional.Bool
	AllowIncorrectBarcodes               optional.Bool
	AllowInvertImage                     optional.Bool
	AllowMicroWhiteSpotsRemoving         optional.Bool
	AllowOneDFastBarcodesDetector        optional.Bool
	AllowOneDWipedBarsRestoration        optional.Bool
	AllowQRMicroQrRestoration            optional.Bool
	AllowRegularImage                    optional.Bool
	AllowSaltAndPepperFiltering          optional.Bool
	AllowWhiteSpotsRemoving              optional.Bool
	CheckMore1DVariants                  optional.Bool
	FastScanOnly                         optional.Bool
	RegionLikelihoodThresholdPercent     optional.Float64
	ScanWindowSizes                      optional.Interface
	Similarity                           optional.Float64
	SkipDiagonalSearch                   optional.Bool
	ReadTinyBarcodes                     optional.Bool
	AustralianPostEncodingTable          optional.String
	IgnoreEndingFillingPatternsForCTable optional.Bool
	RectangleRegion                      optional.String
	Url                                  optional.String
	Image                                optional.Interface
}

/*
 * PostBarcodeRecognizeFromUrlOrContent -  Recognize barcode from an url or from request body. Request body can contain raw data bytes of the image or encoded with base64.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  The type of barcode to read.
     * @param "ChecksumValidation" (optional.String) -  Enable checksum validation during recognition for 1D barcodes. Default is treated as Yes for symbologies which must contain checksum, as No where checksum only possible. Checksum never used: Codabar Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN Checksum always used: Rest symbologies
     * @param "DetectEncoding" (optional.Bool) -  A flag which force engine to detect codetext encoding for Unicode.
     * @param "Preset" (optional.String) -  Preset allows to configure recognition quality and speed manually. You can quickly set up Preset by embedded presets: HighPerformance, NormalQuality, HighQuality, MaxBarCodes or you can manually configure separate options. Default value of Preset is NormalQuality.
     * @param "RectX" (optional.Int32) -  Set X for area for recognition.
     * @param "RectY" (optional.Int32) -  Set Y for area for recognition.
     * @param "RectWidth" (optional.Int32) -  Set Width of area for recognition.
     * @param "RectHeight" (optional.Int32) -  Set Height of area for recognition.
     * @param "StripFNC" (optional.Bool) -  Value indicating whether FNC symbol strip must be done.
     * @param "Timeout" (optional.Int32) -  Timeout of recognition process.
     * @param "MedianSmoothingWindowSize" (optional.Int32) -  Window size for median smoothing. Typical values are 3 or 4. Default value is 3. AllowMedianSmoothing must be set.
     * @param "AllowMedianSmoothing" (optional.Bool) -  Allows engine to enable median smoothing as additional scan. Mode helps to recognize noised barcodes.
     * @param "AllowComplexBackground" (optional.Bool) -  Allows engine to recognize color barcodes on color background as additional scan. Extremely slow mode.
     * @param "AllowDatamatrixIndustrialBarcodes" (optional.Bool) -  Allows engine for Datamatrix to recognize dashed industrial Datamatrix barcodes. Slow mode which helps only for dashed barcodes which consist from spots.
     * @param "AllowDecreasedImage" (optional.Bool) -  Allows engine to recognize decreased image as additional scan. Size for decreasing is selected by internal engine algorithms. Mode helps to recognize barcodes which are noised and blurred but captured with high resolution.
     * @param "AllowDetectScanGap" (optional.Bool) -  Allows engine to use gap between scans to increase recognition speed. Mode can make recognition problems with low height barcodes.
     * @param "AllowIncorrectBarcodes" (optional.Bool) -  Allows engine to recognize barcodes which has incorrect checksum or incorrect values. Mode can be used to recognize damaged barcodes with incorrect text.
     * @param "AllowInvertImage" (optional.Bool) -  Allows engine to recognize inverse color image as additional scan. Mode can be used when barcode is white on black background.
     * @param "AllowMicroWhiteSpotsRemoving" (optional.Bool) -  Allows engine for Postal barcodes to recognize slightly noised images. Mode helps to recognize slightly damaged Postal barcodes.
     * @param "AllowOneDFastBarcodesDetector" (optional.Bool) -  Allows engine for 1D barcodes to quickly recognize high quality barcodes which fill almost whole image. Mode helps to quickly recognize generated barcodes from Internet.
     * @param "AllowOneDWipedBarsRestoration" (optional.Bool) -  Allows engine for 1D barcodes to recognize barcodes with single wiped/glued bars in pattern.
     * @param "AllowQRMicroQrRestoration" (optional.Bool) -  Allows engine for QR/MicroQR to recognize damaged MicroQR barcodes.
     * @param "AllowRegularImage" (optional.Bool) -  Allows engine to recognize regular image without any restorations as main scan. Mode to recognize image as is.
     * @param "AllowSaltAndPepperFiltering" (optional.Bool) -  Allows engine to recognize barcodes with salt and pepper noise type. Mode can remove small noise with white and black dots.
     * @param "AllowWhiteSpotsRemoving" (optional.Bool) -  Allows engine to recognize image without small white spots as additional scan. Mode helps to recognize noised image as well as median smoothing filtering.
     * @param "CheckMore1DVariants" (optional.Bool) -  Allows engine to recognize 1D barcodes with checksum by checking more recognition variants. Default value: False.
     * @param "FastScanOnly" (optional.Bool) -  Allows engine for 1D barcodes to quickly recognize middle slice of an image and return result without using any time-consuming algorithms. Default value: False.
     * @param "RegionLikelihoodThresholdPercent" (optional.Float64) -  Sets threshold for detected regions that may contain barcodes. Value 0.7 means that bottom 70% of possible regions are filtered out and not processed further. Region likelihood threshold must be between [0.05, 0.9] Use high values for clear images with few barcodes. Use low values for images with many barcodes or for noisy images. Low value may lead to a bigger recognition time.
     * @param "ScanWindowSizes" (optional.Interface of []int32) -  Scan window sizes in pixels. Allowed sizes are 10, 15, 20, 25, 30. Scanning with small window size takes more time and provides more accuracy but may fail in detecting very big barcodes. Combining of several window sizes can improve detection quality.
     * @param "Similarity" (optional.Float64) -  Similarity coefficient depends on how homogeneous barcodes are. Use high value for for clear barcodes. Use low values to detect barcodes that ara partly damaged or not lighten evenly. Similarity coefficient must be between [0.5, 0.9]
     * @param "SkipDiagonalSearch" (optional.Bool) -  Allows detector to skip search for diagonal barcodes. Setting it to false will increase detection time but allow to find diagonal barcodes that can be missed otherwise. Enabling of diagonal search leads to a bigger detection time.
     * @param "ReadTinyBarcodes" (optional.Bool) -  Allows engine to recognize tiny barcodes on large images. Ignored if AllowIncorrectBarcodes is set to True. Default value: False.
     * @param "AustralianPostEncodingTable" (optional.String) -  Interpreting Type for the Customer Information of AustralianPost BarCode.Default is CustomerInformationInterpretingType.Other.
     * @param "IgnoreEndingFillingPatternsForCTable" (optional.Bool) -  The flag which force AustraliaPost decoder to ignore last filling patterns in Customer Information Field during decoding as CTable method. CTable encoding method does not have any gaps in encoding table and sequnce \&quot;333\&quot; of filling paterns is decoded as letter \&quot;z\&quot;.
     * @param "RectangleRegion" (optional.String) -
     * @param "Url" (optional.String) -  The image file url.
     * @param "Image" (optional.Interface of *os.File) -  Image data

 * @return BarcodeResponseList
*/
func (a *BarcodeApiService) PostBarcodeRecognizeFromUrlOrContent(ctx context.Context, optionals *BarcodeApiPostBarcodeRecognizeFromUrlOrContentOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/recognize"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Type_.IsSet() {
		queryParams.Add("Type", parameterToString(optionals.Type_.Value(), ""))
	}
	if optionals != nil && optionals.ChecksumValidation.IsSet() {
		queryParams.Add("ChecksumValidation", parameterToString(optionals.ChecksumValidation.Value(), ""))
	}
	if optionals != nil && optionals.DetectEncoding.IsSet() {
		queryParams.Add("DetectEncoding", parameterToString(optionals.DetectEncoding.Value(), ""))
	}
	if optionals != nil && optionals.Preset.IsSet() {
		queryParams.Add("Preset", parameterToString(optionals.Preset.Value(), ""))
	}
	if optionals != nil && optionals.RectX.IsSet() {
		queryParams.Add("RectX", parameterToString(optionals.RectX.Value(), ""))
	}
	if optionals != nil && optionals.RectY.IsSet() {
		queryParams.Add("RectY", parameterToString(optionals.RectY.Value(), ""))
	}
	if optionals != nil && optionals.RectWidth.IsSet() {
		queryParams.Add("RectWidth", parameterToString(optionals.RectWidth.Value(), ""))
	}
	if optionals != nil && optionals.RectHeight.IsSet() {
		queryParams.Add("RectHeight", parameterToString(optionals.RectHeight.Value(), ""))
	}
	if optionals != nil && optionals.StripFNC.IsSet() {
		queryParams.Add("StripFNC", parameterToString(optionals.StripFNC.Value(), ""))
	}
	if optionals != nil && optionals.Timeout.IsSet() {
		queryParams.Add("Timeout", parameterToString(optionals.Timeout.Value(), ""))
	}
	if optionals != nil && optionals.MedianSmoothingWindowSize.IsSet() {
		queryParams.Add("MedianSmoothingWindowSize", parameterToString(optionals.MedianSmoothingWindowSize.Value(), ""))
	}
	if optionals != nil && optionals.AllowMedianSmoothing.IsSet() {
		queryParams.Add("AllowMedianSmoothing", parameterToString(optionals.AllowMedianSmoothing.Value(), ""))
	}
	if optionals != nil && optionals.AllowComplexBackground.IsSet() {
		queryParams.Add("AllowComplexBackground", parameterToString(optionals.AllowComplexBackground.Value(), ""))
	}
	if optionals != nil && optionals.AllowDatamatrixIndustrialBarcodes.IsSet() {
		queryParams.Add("AllowDatamatrixIndustrialBarcodes", parameterToString(optionals.AllowDatamatrixIndustrialBarcodes.Value(), ""))
	}
	if optionals != nil && optionals.AllowDecreasedImage.IsSet() {
		queryParams.Add("AllowDecreasedImage", parameterToString(optionals.AllowDecreasedImage.Value(), ""))
	}
	if optionals != nil && optionals.AllowDetectScanGap.IsSet() {
		queryParams.Add("AllowDetectScanGap", parameterToString(optionals.AllowDetectScanGap.Value(), ""))
	}
	if optionals != nil && optionals.AllowIncorrectBarcodes.IsSet() {
		queryParams.Add("AllowIncorrectBarcodes", parameterToString(optionals.AllowIncorrectBarcodes.Value(), ""))
	}
	if optionals != nil && optionals.AllowInvertImage.IsSet() {
		queryParams.Add("AllowInvertImage", parameterToString(optionals.AllowInvertImage.Value(), ""))
	}
	if optionals != nil && optionals.AllowMicroWhiteSpotsRemoving.IsSet() {
		queryParams.Add("AllowMicroWhiteSpotsRemoving", parameterToString(optionals.AllowMicroWhiteSpotsRemoving.Value(), ""))
	}
	if optionals != nil && optionals.AllowOneDFastBarcodesDetector.IsSet() {
		queryParams.Add("AllowOneDFastBarcodesDetector", parameterToString(optionals.AllowOneDFastBarcodesDetector.Value(), ""))
	}
	if optionals != nil && optionals.AllowOneDWipedBarsRestoration.IsSet() {
		queryParams.Add("AllowOneDWipedBarsRestoration", parameterToString(optionals.AllowOneDWipedBarsRestoration.Value(), ""))
	}
	if optionals != nil && optionals.AllowQRMicroQrRestoration.IsSet() {
		queryParams.Add("AllowQRMicroQrRestoration", parameterToString(optionals.AllowQRMicroQrRestoration.Value(), ""))
	}
	if optionals != nil && optionals.AllowRegularImage.IsSet() {
		queryParams.Add("AllowRegularImage", parameterToString(optionals.AllowRegularImage.Value(), ""))
	}
	if optionals != nil && optionals.AllowSaltAndPepperFiltering.IsSet() {
		queryParams.Add("AllowSaltAndPepperFiltering", parameterToString(optionals.AllowSaltAndPepperFiltering.Value(), ""))
	}
	if optionals != nil && optionals.AllowWhiteSpotsRemoving.IsSet() {
		queryParams.Add("AllowWhiteSpotsRemoving", parameterToString(optionals.AllowWhiteSpotsRemoving.Value(), ""))
	}
	if optionals != nil && optionals.CheckMore1DVariants.IsSet() {
		queryParams.Add("CheckMore1DVariants", parameterToString(optionals.CheckMore1DVariants.Value(), ""))
	}
	if optionals != nil && optionals.FastScanOnly.IsSet() {
		queryParams.Add("FastScanOnly", parameterToString(optionals.FastScanOnly.Value(), ""))
	}
	if optionals != nil && optionals.RegionLikelihoodThresholdPercent.IsSet() {
		queryParams.Add("RegionLikelihoodThresholdPercent", parameterToString(optionals.RegionLikelihoodThresholdPercent.Value(), ""))
	}
	if optionals != nil && optionals.ScanWindowSizes.IsSet() {
		queryParams.Add("ScanWindowSizes", parameterToString(optionals.ScanWindowSizes.Value(), "multi"))
	}
	if optionals != nil && optionals.Similarity.IsSet() {
		queryParams.Add("Similarity", parameterToString(optionals.Similarity.Value(), ""))
	}
	if optionals != nil && optionals.SkipDiagonalSearch.IsSet() {
		queryParams.Add("SkipDiagonalSearch", parameterToString(optionals.SkipDiagonalSearch.Value(), ""))
	}
	if optionals != nil && optionals.ReadTinyBarcodes.IsSet() {
		queryParams.Add("ReadTinyBarcodes", parameterToString(optionals.ReadTinyBarcodes.Value(), ""))
	}
	if optionals != nil && optionals.AustralianPostEncodingTable.IsSet() {
		queryParams.Add("AustralianPostEncodingTable", parameterToString(optionals.AustralianPostEncodingTable.Value(), ""))
	}
	if optionals != nil && optionals.IgnoreEndingFillingPatternsForCTable.IsSet() {
		queryParams.Add("IgnoreEndingFillingPatternsForCTable", parameterToString(optionals.IgnoreEndingFillingPatternsForCTable.Value(), ""))
	}
	if optionals != nil && optionals.RectangleRegion.IsSet() {
		queryParams.Add("RectangleRegion", parameterToString(optionals.RectangleRegion.Value(), ""))
	}
	if optionals != nil && optionals.Url.IsSet() {
		queryParams.Add("url", parameterToString(optionals.Url.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"multipart/form-data", "application/x-www-form-urlencoded", "application/octet-stream"}

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

	if optionals != nil && optionals.Image.IsSet() {
		if requestFile, fileOk := optionals.Image.Value().(*os.File); fileOk {
			fileName = requestFile.Name()
			var err error
			fileBytes, err = ioutil.ReadAll(requestFile)
			if err != nil {
				return returnValue, nil, err
			}
		} else if requestBytes, bytesOK := optionals.Image.Value().([]byte); bytesOK {
			fileName = "Image"
			fileBytes = requestBytes
		} else {
			return returnValue, nil, reportError("image should be *os.File or []byte")
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
			var v BarcodeResponseList
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

//BarcodeApiPostGenerateMultipleOpts - Optional Parameters for BarcodeApiPostGenerateMultiple
type BarcodeApiPostGenerateMultipleOpts struct {
	Format optional.String
}

/*
 * PostGenerateMultiple -  Generate multiple barcodes and return in response stream
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param generatorParamsList List of barcodes
 * @param optional nil or *BarcodeApiPostGenerateMultipleOpts - Optional Parameters:
     * @param "Format" (optional.String) -  Format to return stream in

 * @return []byte
*/
func (a *BarcodeApiService) PostGenerateMultiple(ctx context.Context, generatorParamsList GeneratorParamsList, optionals *BarcodeApiPostGenerateMultipleOpts) ([]byte, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []byte
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/generateMultiple"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Format.IsSet() {
		queryParams.Add("format", parameterToString(optionals.Format.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json", "application/xml"}

	// set Content-Type header
	httpContentType := selectHeaderContentType(contentTypeChoices)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// to determine Accept header
	acceptChoices := []string{"image/png", "image/bmp", "image/gif", "image/jpeg", "image/svg+xml", "image/tiff"}

	// set Accept header
	httpHeaderAccept := selectHeaderAccept(acceptChoices)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}
	// body params
	postBody = &generatorParamsList
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

//BarcodeApiPutBarcodeGenerateFileOpts - Optional Parameters for BarcodeApiPutBarcodeGenerateFile
type BarcodeApiPutBarcodeGenerateFileOpts struct {
	TwoDDisplayText    optional.String
	TextLocation       optional.String
	TextAlignment      optional.String
	TextColor          optional.String
	FontSizeMode       optional.String
	NoWrap             optional.Bool
	Resolution         optional.Float64
	ResolutionX        optional.Float64
	ResolutionY        optional.Float64
	DimensionX         optional.Float64
	TextSpace          optional.Float64
	Units              optional.String
	SizeMode           optional.String
	BarHeight          optional.Float64
	ImageHeight        optional.Float64
	ImageWidth         optional.Float64
	RotationAngle      optional.Float64
	BackColor          optional.String
	BarColor           optional.String
	BorderColor        optional.String
	BorderWidth        optional.Float64
	BorderDashStyle    optional.String
	BorderVisible      optional.Bool
	EnableChecksum     optional.String
	EnableEscape       optional.Bool
	FilledBars         optional.Bool
	AlwaysShowChecksum optional.Bool
	WideNarrowRatio    optional.Float64
	ValidateText       optional.Bool
	SupplementData     optional.String
	SupplementSpace    optional.Float64
	BarWidthReduction  optional.Float64
	Storage            optional.String
	Folder             optional.String
	Format             optional.String
}

/*
 * PutBarcodeGenerateFile -  Generate barcode and save on server (from query params or from file with json or xml content)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The image file name.
 * @param type_ Type of barcode to generate.
 * @param text Text to encode.
 * @param optional nil or *BarcodeApiPutBarcodeGenerateFileOpts - Optional Parameters:
     * @param "TwoDDisplayText" (optional.String) -  Text that will be displayed instead of codetext in 2D barcodes. Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode
     * @param "TextLocation" (optional.String) -  Specify the displaying Text Location, set to CodeLocation.None to hide CodeText. Default value: CodeLocation.Below.
     * @param "TextAlignment" (optional.String) -  Text alignment.
     * @param "TextColor" (optional.String) -  Specify the displaying CodeText&#39;s Color. Default value: Color.Black.
     * @param "FontSizeMode" (optional.String) -  Specify FontSizeMode. If FontSizeMode is set to Auto, font size will be calculated automatically based on xDimension value. It is recommended to use FontSizeMode.Auto especially in AutoSizeMode.Nearest or AutoSizeMode.Interpolation. Default value: FontSizeMode.Auto.
     * @param "NoWrap" (optional.Bool) -  Specify word wraps (line breaks) within text. Default value: false.
     * @param "Resolution" (optional.Float64) -  Resolution of the BarCode image. One value for both dimensions. Default value: 96 dpi.
     * @param "ResolutionX" (optional.Float64) -  DEPRECATED: Use &#39;Resolution&#39; instead.
     * @param "ResolutionY" (optional.Float64) -  DEPRECATED: Use &#39;Resolution&#39; instead.
     * @param "DimensionX" (optional.Float64) -  The smallest width of the unit of BarCode bars or spaces. Increase this will increase the whole barcode image width. Ignored if AutoSizeMode property is set to AutoSizeMode.Nearest or AutoSizeMode.Interpolation.
     * @param "TextSpace" (optional.Float64) -  Space between the CodeText and the BarCode in Unit value. Default value: 2pt. Ignored for EAN8, EAN13, UPCE, UPCA, ISBN, ISMN, ISSN, UpcaGs1DatabarCoupon.
     * @param "Units" (optional.String) -  Common Units for all measuring in query. Default units: pixel.
     * @param "SizeMode" (optional.String) -  Specifies the different types of automatic sizing modes. Default value: AutoSizeMode.None.
     * @param "BarHeight" (optional.Float64) -  Height of the barcode in given units. Default units: pixel.
     * @param "ImageHeight" (optional.Float64) -  Height of the barcode image in given units. Default units: pixel.
     * @param "ImageWidth" (optional.Float64) -  Width of the barcode image in given units. Default units: pixel.
     * @param "RotationAngle" (optional.Float64) -  BarCode image rotation angle, measured in degree, e.g. RotationAngle &#x3D; 0 or RotationAngle &#x3D; 360 means no rotation. If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
     * @param "BackColor" (optional.String) -  Background color of the barcode image. Default value: Color.White.
     * @param "BarColor" (optional.String) -  Bars color. Default value: Color.Black.
     * @param "BorderColor" (optional.String) -  Border color. Default value: Color.Black.
     * @param "BorderWidth" (optional.Float64) -  Border width. Default value: 0. Ignored if Visible is set to false.
     * @param "BorderDashStyle" (optional.String) -  Border dash style. Default value: BorderDashStyle.Solid.
     * @param "BorderVisible" (optional.Bool) -  Border visibility. If false than parameter Width is always ignored (0). Default value: false.
     * @param "EnableChecksum" (optional.String) -  Enable checksum during generation 1D barcodes. Default is treated as Yes for symbology which must contain checksum, as No where checksum only possible. Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN, Codabar Checksum always used: Rest symbology
     * @param "EnableEscape" (optional.Bool) -  Indicates whether explains the character \&quot;\\\&quot; as an escape character in CodeText property. Used for Pdf417, DataMatrix, Code128 only If the EnableEscape is true, \&quot;\\\&quot; will be explained as a special escape character. Otherwise, \&quot;\\\&quot; acts as normal characters. Aspose.BarCode supports input decimal ascii code and mnemonic for ASCII control-code characters. For example, \\013 and \\\\CR stands for CR.
     * @param "FilledBars" (optional.Bool) -  Value indicating whether bars are filled. Only for 1D barcodes. Default value: true.
     * @param "AlwaysShowChecksum" (optional.Bool) -  Always display checksum digit in the human readable text for Code128 and GS1Code128 barcodes.
     * @param "WideNarrowRatio" (optional.Float64) -  Wide bars to Narrow bars ratio. Default value: 3, that is, wide bars are 3 times as wide as narrow bars. Used for ITF, PZN, PharmaCode, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, IATA2of5, VIN, DeutschePost, OPC, Code32, DataLogic2of5, PatchCode, Code39Extended, Code39Standard
     * @param "ValidateText" (optional.Bool) -  Only for 1D barcodes. If codetext is incorrect and value set to true - exception will be thrown. Otherwise codetext will be corrected to match barcode&#39;s specification. Exception always will be thrown for: Databar symbology if codetext is incorrect. Exception always will not be thrown for: AustraliaPost, SingaporePost, Code39Extended, Code93Extended, Code16K, Code128 symbology if codetext is incorrect.
     * @param "SupplementData" (optional.String) -  Supplement parameters. Used for Interleaved2of5, Standard2of5, EAN13, EAN8, UPCA, UPCE, ISBN, ISSN, ISMN.
     * @param "SupplementSpace" (optional.Float64) -  Space between main the BarCode and supplement BarCode.
     * @param "BarWidthReduction" (optional.Float64) -  Bars reduction value that is used to compensate ink spread while printing.
     * @param "Storage" (optional.String) -  Image&#39;s storage.
     * @param "Folder" (optional.String) -  Image&#39;s folder.
     * @param "Format" (optional.String) -  The image format.

 * @return ResultImageInfo
*/
func (a *BarcodeApiService) PutBarcodeGenerateFile(ctx context.Context, name string, type_ string, text string, optionals *BarcodeApiPutBarcodeGenerateFileOpts) (ResultImageInfo, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue ResultImageInfo
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/{name}/generate"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("Type", parameterToString(type_, ""))
	queryParams.Add("Text", parameterToString(text, ""))
	if optionals != nil && optionals.TwoDDisplayText.IsSet() {
		queryParams.Add("TwoDDisplayText", parameterToString(optionals.TwoDDisplayText.Value(), ""))
	}
	if optionals != nil && optionals.TextLocation.IsSet() {
		queryParams.Add("TextLocation", parameterToString(optionals.TextLocation.Value(), ""))
	}
	if optionals != nil && optionals.TextAlignment.IsSet() {
		queryParams.Add("TextAlignment", parameterToString(optionals.TextAlignment.Value(), ""))
	}
	if optionals != nil && optionals.TextColor.IsSet() {
		queryParams.Add("TextColor", parameterToString(optionals.TextColor.Value(), ""))
	}
	if optionals != nil && optionals.FontSizeMode.IsSet() {
		queryParams.Add("FontSizeMode", parameterToString(optionals.FontSizeMode.Value(), ""))
	}
	if optionals != nil && optionals.NoWrap.IsSet() {
		queryParams.Add("NoWrap", parameterToString(optionals.NoWrap.Value(), ""))
	}
	if optionals != nil && optionals.Resolution.IsSet() {
		queryParams.Add("Resolution", parameterToString(optionals.Resolution.Value(), ""))
	}
	if optionals != nil && optionals.ResolutionX.IsSet() {
		queryParams.Add("ResolutionX", parameterToString(optionals.ResolutionX.Value(), ""))
	}
	if optionals != nil && optionals.ResolutionY.IsSet() {
		queryParams.Add("ResolutionY", parameterToString(optionals.ResolutionY.Value(), ""))
	}
	if optionals != nil && optionals.DimensionX.IsSet() {
		queryParams.Add("DimensionX", parameterToString(optionals.DimensionX.Value(), ""))
	}
	if optionals != nil && optionals.TextSpace.IsSet() {
		queryParams.Add("TextSpace", parameterToString(optionals.TextSpace.Value(), ""))
	}
	if optionals != nil && optionals.Units.IsSet() {
		queryParams.Add("Units", parameterToString(optionals.Units.Value(), ""))
	}
	if optionals != nil && optionals.SizeMode.IsSet() {
		queryParams.Add("SizeMode", parameterToString(optionals.SizeMode.Value(), ""))
	}
	if optionals != nil && optionals.BarHeight.IsSet() {
		queryParams.Add("BarHeight", parameterToString(optionals.BarHeight.Value(), ""))
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
	if optionals != nil && optionals.BackColor.IsSet() {
		queryParams.Add("BackColor", parameterToString(optionals.BackColor.Value(), ""))
	}
	if optionals != nil && optionals.BarColor.IsSet() {
		queryParams.Add("BarColor", parameterToString(optionals.BarColor.Value(), ""))
	}
	if optionals != nil && optionals.BorderColor.IsSet() {
		queryParams.Add("BorderColor", parameterToString(optionals.BorderColor.Value(), ""))
	}
	if optionals != nil && optionals.BorderWidth.IsSet() {
		queryParams.Add("BorderWidth", parameterToString(optionals.BorderWidth.Value(), ""))
	}
	if optionals != nil && optionals.BorderDashStyle.IsSet() {
		queryParams.Add("BorderDashStyle", parameterToString(optionals.BorderDashStyle.Value(), ""))
	}
	if optionals != nil && optionals.BorderVisible.IsSet() {
		queryParams.Add("BorderVisible", parameterToString(optionals.BorderVisible.Value(), ""))
	}
	if optionals != nil && optionals.EnableChecksum.IsSet() {
		queryParams.Add("EnableChecksum", parameterToString(optionals.EnableChecksum.Value(), ""))
	}
	if optionals != nil && optionals.EnableEscape.IsSet() {
		queryParams.Add("EnableEscape", parameterToString(optionals.EnableEscape.Value(), ""))
	}
	if optionals != nil && optionals.FilledBars.IsSet() {
		queryParams.Add("FilledBars", parameterToString(optionals.FilledBars.Value(), ""))
	}
	if optionals != nil && optionals.AlwaysShowChecksum.IsSet() {
		queryParams.Add("AlwaysShowChecksum", parameterToString(optionals.AlwaysShowChecksum.Value(), ""))
	}
	if optionals != nil && optionals.WideNarrowRatio.IsSet() {
		queryParams.Add("WideNarrowRatio", parameterToString(optionals.WideNarrowRatio.Value(), ""))
	}
	if optionals != nil && optionals.ValidateText.IsSet() {
		queryParams.Add("ValidateText", parameterToString(optionals.ValidateText.Value(), ""))
	}
	if optionals != nil && optionals.SupplementData.IsSet() {
		queryParams.Add("SupplementData", parameterToString(optionals.SupplementData.Value(), ""))
	}
	if optionals != nil && optionals.SupplementSpace.IsSet() {
		queryParams.Add("SupplementSpace", parameterToString(optionals.SupplementSpace.Value(), ""))
	}
	if optionals != nil && optionals.BarWidthReduction.IsSet() {
		queryParams.Add("BarWidthReduction", parameterToString(optionals.BarWidthReduction.Value(), ""))
	}
	if optionals != nil && optionals.Storage.IsSet() {
		queryParams.Add("storage", parameterToString(optionals.Storage.Value(), ""))
	}
	if optionals != nil && optionals.Folder.IsSet() {
		queryParams.Add("folder", parameterToString(optionals.Folder.Value(), ""))
	}
	if optionals != nil && optionals.Format.IsSet() {
		queryParams.Add("format", parameterToString(optionals.Format.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"multipart/form-data", "application/x-www-form-urlencoded", "application/json", "application/xml"}

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
			error: httpResponse.Status,
			text:  string(responseBody),
		}

		if httpResponse.StatusCode == 200 {
			var v ResultImageInfo
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

//BarcodeApiPutBarcodeRecognizeFromBodyOpts - Optional Parameters for BarcodeApiPutBarcodeRecognizeFromBody
type BarcodeApiPutBarcodeRecognizeFromBodyOpts struct {
	Type_   optional.String
	Storage optional.String
	Folder  optional.String
}

/*
 * PutBarcodeRecognizeFromBody -  Recognition of a barcode from file on server with parameters in body.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The image file name.
 * @param readerParams BarcodeReader object with parameters.
 * @param optional nil or *BarcodeApiPutBarcodeRecognizeFromBodyOpts - Optional Parameters:
     * @param "Type_" (optional.String) -
     * @param "Storage" (optional.String) -  The storage name
     * @param "Folder" (optional.String) -  The image folder.

 * @return BarcodeResponseList
*/
func (a *BarcodeApiService) PutBarcodeRecognizeFromBody(ctx context.Context, name string, readerParams ReaderParams, optionals *BarcodeApiPutBarcodeRecognizeFromBodyOpts) (BarcodeResponseList, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue BarcodeResponseList
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/{name}/recognize"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Type_.IsSet() {
		queryParams.Add("type", parameterToString(optionals.Type_.Value(), ""))
	}
	if optionals != nil && optionals.Storage.IsSet() {
		queryParams.Add("storage", parameterToString(optionals.Storage.Value(), ""))
	}
	if optionals != nil && optionals.Folder.IsSet() {
		queryParams.Add("folder", parameterToString(optionals.Folder.Value(), ""))
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
	// body params
	postBody = &readerParams
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
			var v BarcodeResponseList
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

//BarcodeApiPutGenerateMultipleOpts - Optional Parameters for BarcodeApiPutGenerateMultiple
type BarcodeApiPutGenerateMultipleOpts struct {
	Format  optional.String
	Folder  optional.String
	Storage optional.String
}

/*
 * PutGenerateMultiple -  Generate image with multiple barcodes and put new file on server
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name New filename
 * @param generatorParamsList List of barcodes
 * @param optional nil or *BarcodeApiPutGenerateMultipleOpts - Optional Parameters:
     * @param "Format" (optional.String) -  Format of file
     * @param "Folder" (optional.String) -  Folder to place file to
     * @param "Storage" (optional.String) -  The storage name

 * @return ResultImageInfo
*/
func (a *BarcodeApiService) PutGenerateMultiple(ctx context.Context, name string, generatorParamsList GeneratorParamsList, optionals *BarcodeApiPutGenerateMultipleOpts) (ResultImageInfo, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue ResultImageInfo
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/barcode/{name}/generateMultiple"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Format.IsSet() {
		queryParams.Add("format", parameterToString(optionals.Format.Value(), ""))
	}
	if optionals != nil && optionals.Folder.IsSet() {
		queryParams.Add("folder", parameterToString(optionals.Folder.Value(), ""))
	}
	if optionals != nil && optionals.Storage.IsSet() {
		queryParams.Add("storage", parameterToString(optionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypeChoices := []string{"application/json", "application/xml"}

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
	// body params
	postBody = &generatorParamsList
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
			var v ResultImageInfo
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
