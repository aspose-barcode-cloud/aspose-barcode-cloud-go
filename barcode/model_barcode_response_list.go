package barcode

// BarcodeResponseList - Represents information about a barcode list.
type BarcodeResponseList struct {
	// List of barcodes that are present in the image.
	Barcodes []BarcodeResponse `json:"barcodes"`
}
