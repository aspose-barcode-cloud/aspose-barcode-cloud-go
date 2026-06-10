package barcode

// EncodeData - Data to encode in a barcode.
type EncodeData struct {
	DataType EncodeDataType `json:"dataType,omitempty"`
	// String that represents the data to encode.
	Data string `json:"data"`
}
