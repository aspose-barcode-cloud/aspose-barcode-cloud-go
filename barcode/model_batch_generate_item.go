package barcode

// BatchGenerateItem - One barcode to generate in a batch.
type BatchGenerateItem struct {
	// Optional file name inside the archive. When omitted or blank, a sequential name is generated. Supplied safe names are preserved; the image format is independent of the file extension.
	FileName string `json:"fileName,omitempty"`
	// Barcode generation parameters.
	Generation GenerateParams `json:"generation"`
}
