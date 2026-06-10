package barcode

// EncodeDataType : Types of data that can be encoded into a barcode.
type EncodeDataType string

// List of EncodeDataType
const (
	EncodeDataTypeStringData  EncodeDataType = "StringData"
	EncodeDataTypeBase64Bytes EncodeDataType = "Base64Bytes"
	EncodeDataTypeHexBytes    EncodeDataType = "HexBytes"
)
