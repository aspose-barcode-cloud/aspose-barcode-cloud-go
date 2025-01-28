package barcode

// BarcodeResponseList - Represents information about barcode list.
type BarcodeResponseList struct {
	// List of barcodes which are present in image.
	Barcodes []BarcodeResponse `json:"barcodes"`
}
