package barcode

// QrParams - QR parameters.
type QrParams struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// DEPRECATED: This property is obsolete and will be removed in future releases. Unicode symbols detection and encoding will be processed in Auto mode with Extended Channel Interpretation charset designator. Using of own encodings requires manual CodeText encoding into byte[] array.  Sets the encoding of codetext.
	TextEncoding string `json:"TextEncoding,omitempty"`
	// QR / MicroQR selector mode. Select ForceQR for standard QR symbols, Auto for MicroQR.
	EncodeType QrEncodeType `json:"EncodeType,omitempty"`
	// Extended Channel Interpretation Identifiers. It is used to tell the barcode reader details about the used references for encoding the data in the symbol. Current implementation consists all well known charset encodings.
	ECIEncoding EciEncodings `json:"ECIEncoding,omitempty"`
	// QR symbology type of BarCode's encoding mode. Default value: QREncodeMode.Auto.
	EncodeMode QrEncodeMode `json:"EncodeMode,omitempty"`
	// Level of Reed-Solomon error correction for QR barcode. From low to high: LevelL, LevelM, LevelQ, LevelH. see QRErrorLevel.
	ErrorLevel QrErrorLevel `json:"ErrorLevel,omitempty"`
	// Version of QR Code. From Version1 to Version40 for QR code and from M1 to M4 for MicroQr. Default value is QRVersion.Auto.
	Version QrVersion `json:"Version,omitempty"`
	// QR structured append parameters.
	StructuredAppend StructuredAppend `json:"StructuredAppend,omitempty"`
}
