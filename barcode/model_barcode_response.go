package barcode

// BarcodeResponse - Represents information about barcode.
type BarcodeResponse struct {
	// Barcode data.
	BarcodeValue NullableString `json:"barcodeValue,omitempty"`
	// Type of the barcode.
	Type NullableString `json:"type,omitempty"`
	// Region with barcode.
	Region []RegionPoint `json:"region,omitempty"`
	// Checksum of barcode.
	Checksum NullableString `json:"checksum,omitempty"`
}
