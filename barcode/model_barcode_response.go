package barcode

// BarcodeResponse - Represents information about barcode.
type BarcodeResponse struct {
	// Barcode data.
	BarcodeValue string `json:"barcodeValue,omitempty"`
	// Type of the barcode.
	Type string `json:"type,omitempty"`
	// Region with barcode.
	Region []RegionPoint `json:"region,omitempty"`
	// Checksum of barcode.
	Checksum string `json:"checksum,omitempty"`
}
