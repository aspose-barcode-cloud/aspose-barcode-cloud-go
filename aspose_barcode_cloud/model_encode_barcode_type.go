/*
 * Aspose.Barcode Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose_barcode_cloud

// EncodeBarcodeType : See EncodeTypes
type EncodeBarcodeType string

// List of EncodeBarcodeType
const (
	CODABAR_EncodeBarcodeType                          EncodeBarcodeType = "Codabar"
	CODE11_EncodeBarcodeType                           EncodeBarcodeType = "Code11"
	CODE39_STANDARD_EncodeBarcodeType                  EncodeBarcodeType = "Code39Standard"
	CODE39_EXTENDED_EncodeBarcodeType                  EncodeBarcodeType = "Code39Extended"
	CODE93_STANDARD_EncodeBarcodeType                  EncodeBarcodeType = "Code93Standard"
	CODE93_EXTENDED_EncodeBarcodeType                  EncodeBarcodeType = "Code93Extended"
	CODE128_EncodeBarcodeType                          EncodeBarcodeType = "Code128"
	GS1_CODE128_EncodeBarcodeType                      EncodeBarcodeType = "GS1Code128"
	EAN8_EncodeBarcodeType                             EncodeBarcodeType = "EAN8"
	EAN13_EncodeBarcodeType                            EncodeBarcodeType = "EAN13"
	EAN14_EncodeBarcodeType                            EncodeBarcodeType = "EAN14"
	SCC14_EncodeBarcodeType                            EncodeBarcodeType = "SCC14"
	SSCC18_EncodeBarcodeType                           EncodeBarcodeType = "SSCC18"
	UPCA_EncodeBarcodeType                             EncodeBarcodeType = "UPCA"
	UPCE_EncodeBarcodeType                             EncodeBarcodeType = "UPCE"
	ISBN_EncodeBarcodeType                             EncodeBarcodeType = "ISBN"
	ISSN_EncodeBarcodeType                             EncodeBarcodeType = "ISSN"
	ISMN_EncodeBarcodeType                             EncodeBarcodeType = "ISMN"
	STANDARD2OF5_EncodeBarcodeType                     EncodeBarcodeType = "Standard2of5"
	INTERLEAVED2OF5_EncodeBarcodeType                  EncodeBarcodeType = "Interleaved2of5"
	MATRIX2OF5_EncodeBarcodeType                       EncodeBarcodeType = "Matrix2of5"
	ITALIAN_POST25_EncodeBarcodeType                   EncodeBarcodeType = "ItalianPost25"
	IATA2OF5_EncodeBarcodeType                         EncodeBarcodeType = "IATA2of5"
	ITF14_EncodeBarcodeType                            EncodeBarcodeType = "ITF14"
	ITF6_EncodeBarcodeType                             EncodeBarcodeType = "ITF6"
	MSI_EncodeBarcodeType                              EncodeBarcodeType = "MSI"
	VIN_EncodeBarcodeType                              EncodeBarcodeType = "VIN"
	DEUTSCHE_POST_IDENTCODE_EncodeBarcodeType          EncodeBarcodeType = "DeutschePostIdentcode"
	DEUTSCHE_POST_LEITCODE_EncodeBarcodeType           EncodeBarcodeType = "DeutschePostLeitcode"
	OPC_EncodeBarcodeType                              EncodeBarcodeType = "OPC"
	PZN_EncodeBarcodeType                              EncodeBarcodeType = "PZN"
	CODE16_K_EncodeBarcodeType                         EncodeBarcodeType = "Code16K"
	PHARMACODE_EncodeBarcodeType                       EncodeBarcodeType = "Pharmacode"
	DATA_MATRIX_EncodeBarcodeType                      EncodeBarcodeType = "DataMatrix"
	QR_EncodeBarcodeType                               EncodeBarcodeType = "QR"
	AZTEC_EncodeBarcodeType                            EncodeBarcodeType = "Aztec"
	PDF417_EncodeBarcodeType                           EncodeBarcodeType = "Pdf417"
	MACRO_PDF417_EncodeBarcodeType                     EncodeBarcodeType = "MacroPdf417"
	AUSTRALIA_POST_EncodeBarcodeType                   EncodeBarcodeType = "AustraliaPost"
	POSTNET_EncodeBarcodeType                          EncodeBarcodeType = "Postnet"
	PLANET_EncodeBarcodeType                           EncodeBarcodeType = "Planet"
	ONE_CODE_EncodeBarcodeType                         EncodeBarcodeType = "OneCode"
	RM4_SCC_EncodeBarcodeType                          EncodeBarcodeType = "RM4SCC"
	DATABAR_OMNI_DIRECTIONAL_EncodeBarcodeType         EncodeBarcodeType = "DatabarOmniDirectional"
	DATABAR_TRUNCATED_EncodeBarcodeType                EncodeBarcodeType = "DatabarTruncated"
	DATABAR_LIMITED_EncodeBarcodeType                  EncodeBarcodeType = "DatabarLimited"
	DATABAR_EXPANDED_EncodeBarcodeType                 EncodeBarcodeType = "DatabarExpanded"
	SINGAPORE_POST_EncodeBarcodeType                   EncodeBarcodeType = "SingaporePost"
	GS1_DATA_MATRIX_EncodeBarcodeType                  EncodeBarcodeType = "GS1DataMatrix"
	AUSTRALIAN_POSTE_PARCEL_EncodeBarcodeType          EncodeBarcodeType = "AustralianPosteParcel"
	SWISS_POST_PARCEL_EncodeBarcodeType                EncodeBarcodeType = "SwissPostParcel"
	PATCH_CODE_EncodeBarcodeType                       EncodeBarcodeType = "PatchCode"
	DATABAR_EXPANDED_STACKED_EncodeBarcodeType         EncodeBarcodeType = "DatabarExpandedStacked"
	DATABAR_STACKED_EncodeBarcodeType                  EncodeBarcodeType = "DatabarStacked"
	DATABAR_STACKED_OMNI_DIRECTIONAL_EncodeBarcodeType EncodeBarcodeType = "DatabarStackedOmniDirectional"
	MICRO_PDF417_EncodeBarcodeType                     EncodeBarcodeType = "MicroPdf417"
	GS1_QR_EncodeBarcodeType                           EncodeBarcodeType = "GS1QR"
	MAXI_CODE_EncodeBarcodeType                        EncodeBarcodeType = "MaxiCode"
	CODE32_EncodeBarcodeType                           EncodeBarcodeType = "Code32"
	DATA_LOGIC2OF5_EncodeBarcodeType                   EncodeBarcodeType = "DataLogic2of5"
	DOT_CODE_EncodeBarcodeType                         EncodeBarcodeType = "DotCode"
	DUTCH_KIX_EncodeBarcodeType                        EncodeBarcodeType = "DutchKIX"
	UPCA_GS1_CODE128_COUPON_EncodeBarcodeType          EncodeBarcodeType = "UpcaGs1Code128Coupon"
	UPCA_GS1_DATABAR_COUPON_EncodeBarcodeType          EncodeBarcodeType = "UpcaGs1DatabarCoupon"
	CODABLOCK_F_EncodeBarcodeType                      EncodeBarcodeType = "CodablockF"
	GS1_CODABLOCK_F_EncodeBarcodeType                  EncodeBarcodeType = "GS1CodablockF"
)