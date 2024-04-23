package barcode

// QrEncodeType :
type QrEncodeType string

// List of QREncodeType
const (
	QrEncodeTypeAuto         QrEncodeType = "Auto"
	QrEncodeTypeForceQR      QrEncodeType = "ForceQR"
	QrEncodeTypeForceMicroQR QrEncodeType = "ForceMicroQR"
)
