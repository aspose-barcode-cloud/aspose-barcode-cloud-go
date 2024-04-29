package barcode

// HanXinParams - HanXin params.
type HanXinParams struct {
	// Encoding mode for XanXin barcodes. Default value: HanXinEncodeMode.Auto.
	EncodeMode HanXinEncodeMode `json:"EncodeMode,omitempty"`
	// Allowed Han Xin error correction levels from L1 to L4. Default value: HanXinErrorLevel.L1.
	ErrorLevel HanXinErrorLevel `json:"ErrorLevel,omitempty"`
	// Allowed Han Xin versions, Auto and Version01 - Version84. Default value: HanXinVersion.Auto.
	Version HanXinVersion `json:"Version,omitempty"`
	// Extended Channel Interpretation Identifiers. It is used to tell the barcode reader details about the used references for encoding the data in the symbol. Current implementation consists all well known charset encodings. Default value: ECIEncodings.ISO_8859_1
	ECIEncoding EciEncodings `json:"ECIEncoding,omitempty"`
}
