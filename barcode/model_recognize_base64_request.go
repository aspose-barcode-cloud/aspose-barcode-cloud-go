package barcode

// RecognizeBase64Request - Barcode recognize request
type RecognizeBase64Request struct {
	// Array of decode types to find on barcode
	BarcodeTypes    []DecodeBarcodeType  `json:"barcodeTypes"`
	RecognitionMode RecognitionMode      `json:"recognitionMode,omitempty"`
	ImageKind       RecognitionImageKind `json:"imageKind,omitempty"`
	// Barcode image bytes encoded as base-64.
	FileBase64 string `json:"fileBase64"`
}
