package barcode

// QrParams - Optional QR barcode generation parameters. Applies to QR, GS1QR, MicroQR, and RectMicroQR barcode types.
type QrParams struct {
	// QR barcode encode mode.
	QrEncodeMode QREncodeMode `json:"qrEncodeMode,omitempty"`
	// QR barcode error correction level.
	QrErrorLevel QRErrorLevel `json:"qrErrorLevel,omitempty"`
	// QR barcode version. Automatically selects the smallest version that fits the data.
	QrVersion QRVersion `json:"qrVersion,omitempty"`
	// ECI encoding for QR barcode data.
	QrECIEncoding ECIEncodings `json:"qrECIEncoding,omitempty"`
	// QR barcode aspect ratio. Values: 0 to 1.
	QrAspectRatio float32 `json:"qrAspectRatio,omitempty"`
	// MicroQR barcode version. Used when BarcodeType is MicroQR.
	MicroQRVersion MicroQRVersion `json:"microQRVersion,omitempty"`
	// RectMicroQR barcode version. Used when BarcodeType is RectMicroQR.
	RectMicroQrVersion RectMicroQRVersion `json:"rectMicroQrVersion,omitempty"`
}
