package barcode

// DataMatrixEncodeMode : DataMatrix encoder's encoding mode, default to Auto
type DataMatrixEncodeMode string

// List of DataMatrixEncodeMode
const (
	DataMatrixEncodeModeAuto             DataMatrixEncodeMode = "Auto"
	DataMatrixEncodeModeASCII            DataMatrixEncodeMode = "ASCII"
	DataMatrixEncodeModeFull             DataMatrixEncodeMode = "Full"
	DataMatrixEncodeModeCustom           DataMatrixEncodeMode = "Custom"
	DataMatrixEncodeModeC40              DataMatrixEncodeMode = "C40"
	DataMatrixEncodeModeText             DataMatrixEncodeMode = "Text"
	DataMatrixEncodeModeEDIFACT          DataMatrixEncodeMode = "EDIFACT"
	DataMatrixEncodeModeANSIX12          DataMatrixEncodeMode = "ANSIX12"
	DataMatrixEncodeModeExtendedCodetext DataMatrixEncodeMode = "ExtendedCodetext"
)
