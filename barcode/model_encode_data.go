package barcode

// EncodeData - Data to encode in a barcode.
type EncodeData struct {
	// Type of data to encode. Default value: StringData.
	DataType EncodeDataType `json:"dataType,omitempty"`
	// String that represents the data to encode.
	Data string `json:"data"`
}
