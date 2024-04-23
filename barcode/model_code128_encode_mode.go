package barcode

// Code128EncodeMode :
type Code128EncodeMode string

// List of Code128EncodeMode
const (
	Code128EncodeModeAuto   Code128EncodeMode = "Auto"
	Code128EncodeModeCodeA  Code128EncodeMode = "CodeA"
	Code128EncodeModeCodeB  Code128EncodeMode = "CodeB"
	Code128EncodeModeCodeAB Code128EncodeMode = "CodeAB"
	Code128EncodeModeCodeC  Code128EncodeMode = "CodeC"
	Code128EncodeModeCodeAC Code128EncodeMode = "CodeAC"
	Code128EncodeModeCodeBC Code128EncodeMode = "CodeBC"
)
