/*
 * MIT License

 * Copyright (c) 2020 Aspose Pty Ltd

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

//GeneratorParams - Represents extended BarcodeGenerator params.
type GeneratorParams struct {
	// Type of barcode to generate.
	TypeOfBarcode EncodeBarcodeType `json:"TypeOfBarcode"`
	// Text to encode.
	Text string `json:"Text"`
	// Text that will be displayed instead of codetext in 2D barcodes. Used for: Aztec, Pdf417, DataMatrix, QR, MaxiCode, DotCode
	TwoDDisplayText string `json:"TwoDDisplayText,omitempty"`
	// Specify the displaying Text Location, set to CodeLocation.None to hide CodeText. Default value: CodeLocation.Below.
	TextLocation CodeLocation `json:"TextLocation,omitempty"`
	// Text alignment.
	TextAlignment TextAlignment `json:"TextAlignment,omitempty"`
	// Specify the displaying CodeText's Color. Default value: Color.Black.
	TextColor string `json:"TextColor,omitempty"`
	// Specify the displaying Text's font. Default value: Arial 5pt regular. Ignored if FontSizeMode is set to FontSizeMode.Auto.
	Font FontParams `json:"Font,omitempty"`
	// Specify FontSizeMode. If FontSizeMode is set to Auto, font size will be calculated automatically based on xDimension value. It is recommended to use FontSizeMode.Auto especially in AutoSizeMode.Nearest or AutoSizeMode.Interpolation. Default value: FontSizeMode.Auto.
	FontSizeMode FontMode `json:"FontSizeMode,omitempty"`
	// Resolution of the BarCode image. One value for both dimensions. Default value: 96 dpi.
	Resolution float64 `json:"Resolution,omitempty"`
	// DEPRECATED: Use 'Resolution' instead.
	ResolutionX float64 `json:"ResolutionX,omitempty"`
	// DEPRECATED: Use 'Resolution' instead.
	ResolutionY float64 `json:"ResolutionY,omitempty"`
	// The smallest width of the unit of BarCode bars or spaces. Increase this will increase the whole barcode image width. Ignored if AutoSizeMode property is set to AutoSizeMode.Nearest or AutoSizeMode.Interpolation.
	DimensionX float64 `json:"DimensionX,omitempty"`
	// Space between the CodeText and the BarCode in Unit value. Default value: 2pt. Ignored for EAN8, EAN13, UPCE, UPCA, ISBN, ISMN, ISSN, UpcaGs1DatabarCoupon.
	TextSpace float64 `json:"TextSpace,omitempty"`
	// Common Units for all measuring in query. Default units: pixel.
	Units AvailableGraphicsUnit `json:"Units,omitempty"`
	// Specifies the different types of automatic sizing modes. Default value: AutoSizeMode.None.
	SizeMode AutoSizeMode `json:"SizeMode,omitempty"`
	// Height of the barcode in given units. Default units: pixel.
	BarHeight float64 `json:"BarHeight,omitempty"`
	// Height of the barcode image in given units. Default units: pixel.
	ImageHeight float64 `json:"ImageHeight,omitempty"`
	// Width of the barcode image in given units. Default units: pixel.
	ImageWidth float64 `json:"ImageWidth,omitempty"`
	// BarCode image rotation angle, measured in degree, e.g. RotationAngle = 0 or RotationAngle = 360 means no rotation. If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
	RotationAngle float64 `json:"RotationAngle,omitempty"`
	// Barcode paddings. Default value: 5pt 5pt 5pt 5pt.
	Padding Padding `json:"Padding,omitempty"`
	// Additional caption above barcode.
	CaptionAbove CaptionParams `json:"CaptionAbove,omitempty"`
	// Additional caption below barcode.
	CaptionBelow CaptionParams `json:"CaptionBelow,omitempty"`
	// Background color of the barcode image. Default value: Color.White.
	BackColor string `json:"BackColor,omitempty"`
	// Bars color. Default value: Color.Black.
	BarColor string `json:"BarColor,omitempty"`
	// Border color. Default value: Color.Black.
	BorderColor string `json:"BorderColor,omitempty"`
	// Border width. Default value: 0. Ignored if Visible is set to false.
	BorderWidth float64 `json:"BorderWidth,omitempty"`
	// Border dash style. Default value: BorderDashStyle.Solid.
	BorderDashStyle BorderDashStyle `json:"BorderDashStyle,omitempty"`
	// Border visibility. If false than parameter Width is always ignored (0). Default value: false.
	BorderVisible bool `json:"BorderVisible,omitempty"`
	// Enable checksum during generation 1D barcodes. Default is treated as Yes for symbology which must contain checksum, as No where checksum only possible. Checksum is possible: Code39 Standard/Extended, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, DeutschePostIdentcode, DeutschePostLeitcode, VIN, Codabar Checksum always used: Rest symbology
	EnableChecksum EnableChecksum `json:"EnableChecksum,omitempty"`
	// Indicates whether explains the character \"\\\" as an escape character in CodeText property. Used for Pdf417, DataMatrix, Code128 only If the EnableEscape is true, \"\\\" will be explained as a special escape character. Otherwise, \"\\\" acts as normal characters. Aspose.BarCode supports input decimal ascii code and mnemonic for ASCII control-code characters. For example, \\013 and \\\\CR stands for CR.
	EnableEscape bool `json:"EnableEscape,omitempty"`
	// Value indicating whether bars are filled. Only for 1D barcodes. Default value: true.
	FilledBars bool `json:"FilledBars,omitempty"`
	// Always display checksum digit in the human readable text for Code128 and GS1Code128 barcodes.
	AlwaysShowChecksum bool `json:"AlwaysShowChecksum,omitempty"`
	// Wide bars to Narrow bars ratio. Default value: 3, that is, wide bars are 3 times as wide as narrow bars. Used for ITF, PZN, PharmaCode, Standard2of5, Interleaved2of5, Matrix2of5, ItalianPost25, IATA2of5, VIN, DeutschePost, OPC, Code32, DataLogic2of5, PatchCode, Code39Extended, Code39Standard
	WideNarrowRatio float64 `json:"WideNarrowRatio,omitempty"`
	// Only for 1D barcodes. If codetext is incorrect and value set to true - exception will be thrown. Otherwise codetext will be corrected to match barcode's specification. Exception always will be thrown for: Databar symbology if codetext is incorrect. Exception always will not be thrown for: AustraliaPost, SingaporePost, Code39Extended, Code93Extended, Code16K, Code128 symbology if codetext is incorrect.
	ValidateText bool `json:"ValidateText,omitempty"`
	// Supplement parameters. Used for Interleaved2of5, Standard2of5, EAN13, EAN8, UPCA, UPCE, ISBN, ISSN, ISMN.
	SupplementData string `json:"SupplementData,omitempty"`
	// Space between main the BarCode and supplement BarCode.
	SupplementSpace float64 `json:"SupplementSpace,omitempty"`
	// Bars reduction value that is used to compensate ink spread while printing.
	BarWidthReduction float64 `json:"BarWidthReduction,omitempty"`
	// AustralianPost params.
	AustralianPost AustralianPostParams `json:"AustralianPost,omitempty"`
	// Aztec params.
	Aztec AztecParams `json:"Aztec,omitempty"`
	// Codabar params.
	Codabar CodabarParams `json:"Codabar,omitempty"`
	// Codablock params.
	Codablock CodablockParams `json:"Codablock,omitempty"`
	// Code16K params.
	Code16K Code16KParams `json:"Code16K,omitempty"`
	// Coupon params.
	Coupon CouponParams `json:"Coupon,omitempty"`
	// DataBar params.
	DataBar DataBarParams `json:"DataBar,omitempty"`
	// DataMatrix params.
	DataMatrix DataMatrixParams `json:"DataMatrix,omitempty"`
	// DotCode params.
	DotCode DotCodeParams `json:"DotCode,omitempty"`
	// ITF params.
	ITF ItfParams `json:"ITF,omitempty"`
	// MaxiCode params.
	MaxiCode MaxiCodeParams `json:"MaxiCode,omitempty"`
	// Pdf417 params.
	Pdf417 Pdf417Params `json:"Pdf417,omitempty"`
	// Postal params.
	Postal PostalParams `json:"Postal,omitempty"`
	// QR params.
	QR QrParams `json:"QR,omitempty"`
	// PatchCode params.
	PatchCode PatchCodeParams `json:"PatchCode,omitempty"`
}
