package barcode

// EncodeData -
type EncodeData struct {
	DataType EncodeDataType `json:"dataType"`
	// String represents data to encode
	Data string `json:"data"`
}
