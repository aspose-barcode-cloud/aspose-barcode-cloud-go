package barcode

// BarcodeResponse - Represents information about a barcode.
type BarcodeResponse struct {
	// Barcode data.
	BarcodeValue string `json:"barcodeValue,omitempty"`
	// Type of the barcode.
	Type string `json:"type,omitempty"`
	// Region with the barcode.
	Region []RegionPoint `json:"region,omitempty"`
	// Checksum of the barcode.
	Checksum string `json:"checksum,omitempty"`
}
