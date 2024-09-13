package barcode

// EncodeData - Data to encode in barcode
type EncodeData struct {
	DataType EncodeDataType `json:"dataType,omitempty"`
	// String represents data to encode
	Data string `json:"data"`
}
