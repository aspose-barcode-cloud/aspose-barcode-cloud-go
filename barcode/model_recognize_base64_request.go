package barcode

// RecognizeBase64Request - Barcode recognition request.
type RecognizeBase64Request struct {
	// Array of barcode decode types to find.
	BarcodeTypes []DecodeBarcodeType `json:"barcodeTypes"`
	// Barcode image bytes encoded as base-64.
	FileBase64           string               `json:"fileBase64"`
	RecognitionMode      RecognitionMode      `json:"recognitionMode,omitempty"`
	RecognitionImageKind RecognitionImageKind `json:"recognitionImageKind,omitempty"`
}
