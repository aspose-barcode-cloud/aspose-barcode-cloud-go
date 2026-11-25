package barcode

// BatchGenerateRequest - Batch barcode generation request.
type BatchGenerateRequest struct {
	// Format of the returned archive. Default value: Zip.
	ArchiveFormat BatchArchiveFormat `json:"archiveFormat,omitempty"`
	// Barcodes to generate. A batch must contain between one and three items.
	Items []BatchGenerateItem `json:"items"`
}
