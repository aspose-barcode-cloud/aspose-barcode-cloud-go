package barcode

// QrEncodeMode :
type QrEncodeMode string

// List of QREncodeMode
const (
	QrEncodeModeAuto             QrEncodeMode = "Auto"
	QrEncodeModeBytes            QrEncodeMode = "Bytes"
	QrEncodeModeUtf8BOM          QrEncodeMode = "Utf8BOM"
	QrEncodeModeUtf16BEBOM       QrEncodeMode = "Utf16BEBOM"
	QrEncodeModeECIEncoding      QrEncodeMode = "ECIEncoding"
	QrEncodeModeExtendedCodetext QrEncodeMode = "ExtendedCodetext"
)
