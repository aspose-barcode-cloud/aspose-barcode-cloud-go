package barcode

// QREncodeMode : QR barcode encode mode. Subset of https://reference.aspose.com/barcode/net/aspose.barcode.generation/qrencodemode/ Obsolete members (Bytes, Utf8BOM, Utf16BEBOM, ECIEncoding, ExtendedCodetext) are omitted.
type QREncodeMode string

// List of QREncodeMode
const (
	QREncodeModeAuto     QREncodeMode = "Auto"
	QREncodeModeExtended QREncodeMode = "Extended"
	QREncodeModeBinary   QREncodeMode = "Binary"
	QREncodeModeECI      QREncodeMode = "ECI"
)
