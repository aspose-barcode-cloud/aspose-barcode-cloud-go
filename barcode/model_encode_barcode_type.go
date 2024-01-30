/*
 * MIT License

 * Copyright (c) 2024 Aspose Pty Ltd

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

// EncodeBarcodeType : See EncodeTypes
type EncodeBarcodeType string

// List of EncodeBarcodeType
const (
	EncodeBarcodeTypeCodabar                       EncodeBarcodeType = "Codabar"
	EncodeBarcodeTypeCode11                        EncodeBarcodeType = "Code11"
	EncodeBarcodeTypeCode39Standard                EncodeBarcodeType = "Code39Standard"
	EncodeBarcodeTypeCode39Extended                EncodeBarcodeType = "Code39Extended"
	EncodeBarcodeTypeCode93Standard                EncodeBarcodeType = "Code93Standard"
	EncodeBarcodeTypeCode93Extended                EncodeBarcodeType = "Code93Extended"
	EncodeBarcodeTypeCode128                       EncodeBarcodeType = "Code128"
	EncodeBarcodeTypeGS1Code128                    EncodeBarcodeType = "GS1Code128"
	EncodeBarcodeTypeEAN8                          EncodeBarcodeType = "EAN8"
	EncodeBarcodeTypeEAN13                         EncodeBarcodeType = "EAN13"
	EncodeBarcodeTypeEAN14                         EncodeBarcodeType = "EAN14"
	EncodeBarcodeTypeSCC14                         EncodeBarcodeType = "SCC14"
	EncodeBarcodeTypeSSCC18                        EncodeBarcodeType = "SSCC18"
	EncodeBarcodeTypeUPCA                          EncodeBarcodeType = "UPCA"
	EncodeBarcodeTypeUPCE                          EncodeBarcodeType = "UPCE"
	EncodeBarcodeTypeISBN                          EncodeBarcodeType = "ISBN"
	EncodeBarcodeTypeISSN                          EncodeBarcodeType = "ISSN"
	EncodeBarcodeTypeISMN                          EncodeBarcodeType = "ISMN"
	EncodeBarcodeTypeStandard2of5                  EncodeBarcodeType = "Standard2of5"
	EncodeBarcodeTypeInterleaved2of5               EncodeBarcodeType = "Interleaved2of5"
	EncodeBarcodeTypeMatrix2of5                    EncodeBarcodeType = "Matrix2of5"
	EncodeBarcodeTypeItalianPost25                 EncodeBarcodeType = "ItalianPost25"
	EncodeBarcodeTypeIATA2of5                      EncodeBarcodeType = "IATA2of5"
	EncodeBarcodeTypeITF14                         EncodeBarcodeType = "ITF14"
	EncodeBarcodeTypeITF6                          EncodeBarcodeType = "ITF6"
	EncodeBarcodeTypeMSI                           EncodeBarcodeType = "MSI"
	EncodeBarcodeTypeVIN                           EncodeBarcodeType = "VIN"
	EncodeBarcodeTypeDeutschePostIdentcode         EncodeBarcodeType = "DeutschePostIdentcode"
	EncodeBarcodeTypeDeutschePostLeitcode          EncodeBarcodeType = "DeutschePostLeitcode"
	EncodeBarcodeTypeOPC                           EncodeBarcodeType = "OPC"
	EncodeBarcodeTypePZN                           EncodeBarcodeType = "PZN"
	EncodeBarcodeTypeCode16K                       EncodeBarcodeType = "Code16K"
	EncodeBarcodeTypePharmacode                    EncodeBarcodeType = "Pharmacode"
	EncodeBarcodeTypeDataMatrix                    EncodeBarcodeType = "DataMatrix"
	EncodeBarcodeTypeQR                            EncodeBarcodeType = "QR"
	EncodeBarcodeTypeAztec                         EncodeBarcodeType = "Aztec"
	EncodeBarcodeTypePdf417                        EncodeBarcodeType = "Pdf417"
	EncodeBarcodeTypeMacroPdf417                   EncodeBarcodeType = "MacroPdf417"
	EncodeBarcodeTypeAustraliaPost                 EncodeBarcodeType = "AustraliaPost"
	EncodeBarcodeTypePostnet                       EncodeBarcodeType = "Postnet"
	EncodeBarcodeTypePlanet                        EncodeBarcodeType = "Planet"
	EncodeBarcodeTypeOneCode                       EncodeBarcodeType = "OneCode"
	EncodeBarcodeTypeRM4SCC                        EncodeBarcodeType = "RM4SCC"
	EncodeBarcodeTypeDatabarOmniDirectional        EncodeBarcodeType = "DatabarOmniDirectional"
	EncodeBarcodeTypeDatabarTruncated              EncodeBarcodeType = "DatabarTruncated"
	EncodeBarcodeTypeDatabarLimited                EncodeBarcodeType = "DatabarLimited"
	EncodeBarcodeTypeDatabarExpanded               EncodeBarcodeType = "DatabarExpanded"
	EncodeBarcodeTypeSingaporePost                 EncodeBarcodeType = "SingaporePost"
	EncodeBarcodeTypeGS1DataMatrix                 EncodeBarcodeType = "GS1DataMatrix"
	EncodeBarcodeTypeAustralianPosteParcel         EncodeBarcodeType = "AustralianPosteParcel"
	EncodeBarcodeTypeSwissPostParcel               EncodeBarcodeType = "SwissPostParcel"
	EncodeBarcodeTypePatchCode                     EncodeBarcodeType = "PatchCode"
	EncodeBarcodeTypeDatabarExpandedStacked        EncodeBarcodeType = "DatabarExpandedStacked"
	EncodeBarcodeTypeDatabarStacked                EncodeBarcodeType = "DatabarStacked"
	EncodeBarcodeTypeDatabarStackedOmniDirectional EncodeBarcodeType = "DatabarStackedOmniDirectional"
	EncodeBarcodeTypeMicroPdf417                   EncodeBarcodeType = "MicroPdf417"
	EncodeBarcodeTypeGS1QR                         EncodeBarcodeType = "GS1QR"
	EncodeBarcodeTypeMaxiCode                      EncodeBarcodeType = "MaxiCode"
	EncodeBarcodeTypeCode32                        EncodeBarcodeType = "Code32"
	EncodeBarcodeTypeDataLogic2of5                 EncodeBarcodeType = "DataLogic2of5"
	EncodeBarcodeTypeDotCode                       EncodeBarcodeType = "DotCode"
	EncodeBarcodeTypeDutchKIX                      EncodeBarcodeType = "DutchKIX"
	EncodeBarcodeTypeUpcaGs1Code128Coupon          EncodeBarcodeType = "UpcaGs1Code128Coupon"
	EncodeBarcodeTypeUpcaGs1DatabarCoupon          EncodeBarcodeType = "UpcaGs1DatabarCoupon"
	EncodeBarcodeTypeCodablockF                    EncodeBarcodeType = "CodablockF"
	EncodeBarcodeTypeGS1CodablockF                 EncodeBarcodeType = "GS1CodablockF"
	EncodeBarcodeTypeMailmark                      EncodeBarcodeType = "Mailmark"
	EncodeBarcodeTypeGS1DotCode                    EncodeBarcodeType = "GS1DotCode"
	EncodeBarcodeTypeHanXin                        EncodeBarcodeType = "HanXin"
	EncodeBarcodeTypeGS1HanXin                     EncodeBarcodeType = "GS1HanXin"
	EncodeBarcodeTypeGS1Aztec                      EncodeBarcodeType = "GS1Aztec"
	EncodeBarcodeTypeGS1MicroPdf417                EncodeBarcodeType = "GS1MicroPdf417"
)
