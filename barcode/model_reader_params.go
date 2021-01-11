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

//ReaderParams - Represents BarcodeReader object.
type ReaderParams struct {
	// The type of barcode to read.
	Type DecodeBarcodeType `json:"Type,omitempty"`
	// Enable checksum validation during recognition for 1D barcodes. Default is treated as Yes for symbologies which must contain checksum, as No where checksum only possible. Checksum never used: Codabar Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN Checksum always used: Rest symbologies
	ChecksumValidation ChecksumValidation `json:"ChecksumValidation,omitempty"`
	// A flag which force engine to detect codetext encoding for Unicode.
	DetectEncoding bool `json:"DetectEncoding,omitempty"`
	// Preset allows to configure recognition quality and speed manually. You can quickly set up Preset by embedded presets: HighPerformance, NormalQuality, HighQuality, MaxBarCodes or you can manually configure separate options. Default value of Preset is NormalQuality.
	Preset PresetType `json:"Preset,omitempty"`
	// Set X for area for recognition.
	RectX int32 `json:"RectX,omitempty"`
	// Set Y for area for recognition.
	RectY int32 `json:"RectY,omitempty"`
	// Set Width of area for recognition.
	RectWidth int32 `json:"RectWidth,omitempty"`
	// Set Height of area for recognition.
	RectHeight int32 `json:"RectHeight,omitempty"`
	// Value indicating whether FNC symbol strip must be done.
	StripFNC bool `json:"StripFNC,omitempty"`
	// Timeout of recognition process.
	Timeout int32 `json:"Timeout,omitempty"`
	// Window size for median smoothing. Typical values are 3 or 4. Default value is 3. AllowMedianSmoothing must be set.
	MedianSmoothingWindowSize int32 `json:"MedianSmoothingWindowSize,omitempty"`
	// Allows engine to enable median smoothing as additional scan. Mode helps to recognize noised barcodes.
	AllowMedianSmoothing bool `json:"AllowMedianSmoothing,omitempty"`
	// Allows engine to recognize color barcodes on color background as additional scan. Extremely slow mode.
	AllowComplexBackground bool `json:"AllowComplexBackground,omitempty"`
	// Allows engine for Datamatrix to recognize dashed industrial Datamatrix barcodes. Slow mode which helps only for dashed barcodes which consist from spots.
	AllowDatamatrixIndustrialBarcodes bool `json:"AllowDatamatrixIndustrialBarcodes,omitempty"`
	// Allows engine to recognize decreased image as additional scan. Size for decreasing is selected by internal engine algorithms. Mode helps to recognize barcodes which are noised and blurred but captured with high resolution.
	AllowDecreasedImage bool `json:"AllowDecreasedImage,omitempty"`
	// Allows engine to use gap between scans to increase recognition speed. Mode can make recognition problems with low height barcodes.
	AllowDetectScanGap bool `json:"AllowDetectScanGap,omitempty"`
	// Allows engine to recognize barcodes which has incorrect checksum or incorrect values. Mode can be used to recognize damaged barcodes with incorrect text.
	AllowIncorrectBarcodes bool `json:"AllowIncorrectBarcodes,omitempty"`
	// Allows engine to recognize inverse color image as additional scan. Mode can be used when barcode is white on black background.
	AllowInvertImage bool `json:"AllowInvertImage,omitempty"`
	// Allows engine for Postal barcodes to recognize slightly noised images. Mode helps to recognize slightly damaged Postal barcodes.
	AllowMicroWhiteSpotsRemoving bool `json:"AllowMicroWhiteSpotsRemoving,omitempty"`
	// Allows engine for 1D barcodes to quickly recognize high quality barcodes which fill almost whole image. Mode helps to quickly recognize generated barcodes from Internet.
	AllowOneDFastBarcodesDetector bool `json:"AllowOneDFastBarcodesDetector,omitempty"`
	// Allows engine for 1D barcodes to recognize barcodes with single wiped/glued bars in pattern.
	AllowOneDWipedBarsRestoration bool `json:"AllowOneDWipedBarsRestoration,omitempty"`
	// Allows engine for QR/MicroQR to recognize damaged MicroQR barcodes.
	AllowQRMicroQrRestoration bool `json:"AllowQRMicroQrRestoration,omitempty"`
	// Allows engine to recognize regular image without any restorations as main scan. Mode to recognize image as is.
	AllowRegularImage bool `json:"AllowRegularImage,omitempty"`
	// Allows engine to recognize barcodes with salt and pepper noise type. Mode can remove small noise with white and black dots.
	AllowSaltAndPepperFiltering bool `json:"AllowSaltAndPepperFiltering,omitempty"`
	// Allows engine to recognize image without small white spots as additional scan. Mode helps to recognize noised image as well as median smoothing filtering.
	AllowWhiteSpotsRemoving bool `json:"AllowWhiteSpotsRemoving,omitempty"`
	// Sets threshold for detected regions that may contain barcodes. Value 0.7 means that bottom 70% of possible regions are filtered out and not processed further. Region likelihood threshold must be between [0.05, 0.9] Use high values for clear images with few barcodes. Use low values for images with many barcodes or for noisy images. Low value may lead to a bigger recognition time.
	RegionLikelihoodThresholdPercent float64 `json:"RegionLikelihoodThresholdPercent,omitempty"`
	// Scan window sizes in pixels. Allowed sizes are 10, 15, 20, 25, 30. Scanning with small window size takes more time and provides more accuracy but may fail in detecting very big barcodes. Combining of several window sizes can improve detection quality.
	ScanWindowSizes []int32 `json:"ScanWindowSizes,omitempty"`
	// Similarity coefficient depends on how homogeneous barcodes are. Use high value for for clear barcodes. Use low values to detect barcodes that ara partly damaged or not lighten evenly. Similarity coefficient must be between [0.5, 0.9]
	Similarity float64 `json:"Similarity,omitempty"`
	// Allows detector to skip search for diagonal barcodes. Setting it to false will increase detection time but allow to find diagonal barcodes that can be missed otherwise. Enabling of diagonal search leads to a bigger detection time.
	SkipDiagonalSearch bool `json:"SkipDiagonalSearch,omitempty"`
	// Allows engine to recognize tiny barcodes on large images. Ignored if AllowIncorrectBarcodes is set to True. Default value: False.
	ReadTinyBarcodes bool `json:"ReadTinyBarcodes,omitempty"`
	// Interpreting Type for the Customer Information of AustralianPost BarCode.Default is CustomerInformationInterpretingType.Other.
	AustralianPostEncodingTable CustomerInformationInterpretingType `json:"AustralianPostEncodingTable,omitempty"`
}
