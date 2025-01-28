package barcode

// ScanBase64Request - Scan barcode request.
type ScanBase64Request struct {
	// Barcode image bytes encoded as base-64.
	FileBase64 string `json:"fileBase64"`
}
