package barcode

// EncodeDataType : Types of data can be encoded to barcode
type EncodeDataType string

// List of EncodeDataType
const (
	EncodeDataTypeStringData  EncodeDataType = "StringData"
	EncodeDataTypeBase64Bytes EncodeDataType = "Base64Bytes"
	EncodeDataTypeHexBytes    EncodeDataType = "HexBytes"
)
