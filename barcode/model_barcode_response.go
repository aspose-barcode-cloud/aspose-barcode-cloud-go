package barcode

// BarcodeResponse - Represents information about barcode.
type BarcodeResponse struct {
	// Barcode data.
	BarcodeValue string `json:"BarcodeValue,omitempty"`
	// Type of the barcode.
	Type string `json:"Type,omitempty"`
	// Region with barcode.
	Region []RegionPoint `json:"Region,omitempty"`
	// Checksum of barcode.
	Checksum string `json:"Checksum,omitempty"`
}
