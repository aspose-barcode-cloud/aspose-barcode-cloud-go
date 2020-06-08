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

package aspose_barcode_cloud

// DecodeBarcodeType : See DecodeType
type DecodeBarcodeType string

// List of DecodeBarcodeType
const (
	DecodeBarcodeTypeall                           DecodeBarcodeType = "all"
	DecodeBarcodeTypeAustraliaPost                 DecodeBarcodeType = "AustraliaPost"
	DecodeBarcodeTypeAztec                         DecodeBarcodeType = "Aztec"
	DecodeBarcodeTypeISBN                          DecodeBarcodeType = "ISBN"
	DecodeBarcodeTypeCodabar                       DecodeBarcodeType = "Codabar"
	DecodeBarcodeTypeCode11                        DecodeBarcodeType = "Code11"
	DecodeBarcodeTypeCode128                       DecodeBarcodeType = "Code128"
	DecodeBarcodeTypeGS1Code128                    DecodeBarcodeType = "GS1Code128"
	DecodeBarcodeTypeCode39Extended                DecodeBarcodeType = "Code39Extended"
	DecodeBarcodeTypeCode39Standard                DecodeBarcodeType = "Code39Standard"
	DecodeBarcodeTypeCode93Extended                DecodeBarcodeType = "Code93Extended"
	DecodeBarcodeTypeCode93Standard                DecodeBarcodeType = "Code93Standard"
	DecodeBarcodeTypeDataMatrix                    DecodeBarcodeType = "DataMatrix"
	DecodeBarcodeTypeDeutschePostIdentcode         DecodeBarcodeType = "DeutschePostIdentcode"
	DecodeBarcodeTypeDeutschePostLeitcode          DecodeBarcodeType = "DeutschePostLeitcode"
	DecodeBarcodeTypeEAN13                         DecodeBarcodeType = "EAN13"
	DecodeBarcodeTypeEAN14                         DecodeBarcodeType = "EAN14"
	DecodeBarcodeTypeEAN8                          DecodeBarcodeType = "EAN8"
	DecodeBarcodeTypeIATA2of5                      DecodeBarcodeType = "IATA2of5"
	DecodeBarcodeTypeInterleaved2of5               DecodeBarcodeType = "Interleaved2of5"
	DecodeBarcodeTypeISSN                          DecodeBarcodeType = "ISSN"
	DecodeBarcodeTypeISMN                          DecodeBarcodeType = "ISMN"
	DecodeBarcodeTypeItalianPost25                 DecodeBarcodeType = "ItalianPost25"
	DecodeBarcodeTypeITF14                         DecodeBarcodeType = "ITF14"
	DecodeBarcodeTypeITF6                          DecodeBarcodeType = "ITF6"
	DecodeBarcodeTypeMacroPdf417                   DecodeBarcodeType = "MacroPdf417"
	DecodeBarcodeTypeMatrix2of5                    DecodeBarcodeType = "Matrix2of5"
	DecodeBarcodeTypeMSI                           DecodeBarcodeType = "MSI"
	DecodeBarcodeTypeOneCode                       DecodeBarcodeType = "OneCode"
	DecodeBarcodeTypeOPC                           DecodeBarcodeType = "OPC"
	DecodeBarcodeTypePatchCode                     DecodeBarcodeType = "PatchCode"
	DecodeBarcodeTypePdf417                        DecodeBarcodeType = "Pdf417"
	DecodeBarcodeTypeMicroPdf417                   DecodeBarcodeType = "MicroPdf417"
	DecodeBarcodeTypePlanet                        DecodeBarcodeType = "Planet"
	DecodeBarcodeTypePostnet                       DecodeBarcodeType = "Postnet"
	DecodeBarcodeTypePZN                           DecodeBarcodeType = "PZN"
	DecodeBarcodeTypeQR                            DecodeBarcodeType = "QR"
	DecodeBarcodeTypeMicroQR                       DecodeBarcodeType = "MicroQR"
	DecodeBarcodeTypeRM4SCC                        DecodeBarcodeType = "RM4SCC"
	DecodeBarcodeTypeSCC14                         DecodeBarcodeType = "SCC14"
	DecodeBarcodeTypeSSCC18                        DecodeBarcodeType = "SSCC18"
	DecodeBarcodeTypeStandard2of5                  DecodeBarcodeType = "Standard2of5"
	DecodeBarcodeTypeSupplement                    DecodeBarcodeType = "Supplement"
	DecodeBarcodeTypeUPCA                          DecodeBarcodeType = "UPCA"
	DecodeBarcodeTypeUPCE                          DecodeBarcodeType = "UPCE"
	DecodeBarcodeTypeVIN                           DecodeBarcodeType = "VIN"
	DecodeBarcodeTypePharmacode                    DecodeBarcodeType = "Pharmacode"
	DecodeBarcodeTypeGS1DataMatrix                 DecodeBarcodeType = "GS1DataMatrix"
	DecodeBarcodeTypeDatabarOmniDirectional        DecodeBarcodeType = "DatabarOmniDirectional"
	DecodeBarcodeTypeDatabarTruncated              DecodeBarcodeType = "DatabarTruncated"
	DecodeBarcodeTypeDatabarLimited                DecodeBarcodeType = "DatabarLimited"
	DecodeBarcodeTypeDatabarExpanded               DecodeBarcodeType = "DatabarExpanded"
	DecodeBarcodeTypeSwissPostParcel               DecodeBarcodeType = "SwissPostParcel"
	DecodeBarcodeTypeAustralianPosteParcel         DecodeBarcodeType = "AustralianPosteParcel"
	DecodeBarcodeTypeCode16K                       DecodeBarcodeType = "Code16K"
	DecodeBarcodeTypeDatabarStackedOmniDirectional DecodeBarcodeType = "DatabarStackedOmniDirectional"
	DecodeBarcodeTypeDatabarStacked                DecodeBarcodeType = "DatabarStacked"
	DecodeBarcodeTypeDatabarExpandedStacked        DecodeBarcodeType = "DatabarExpandedStacked"
	DecodeBarcodeTypeCompactPdf417                 DecodeBarcodeType = "CompactPdf417"
	DecodeBarcodeTypeGS1QR                         DecodeBarcodeType = "GS1QR"
	DecodeBarcodeTypeMaxiCode                      DecodeBarcodeType = "MaxiCode"
	DecodeBarcodeTypeMicrE13B                      DecodeBarcodeType = "MicrE13B"
	DecodeBarcodeTypeCode32                        DecodeBarcodeType = "Code32"
	DecodeBarcodeTypeDataLogic2of5                 DecodeBarcodeType = "DataLogic2of5"
	DecodeBarcodeTypeDotCode                       DecodeBarcodeType = "DotCode"
	DecodeBarcodeTypeDutchKIX                      DecodeBarcodeType = "DutchKIX"
	DecodeBarcodeTypeCodablockF                    DecodeBarcodeType = "CodablockF"
)
