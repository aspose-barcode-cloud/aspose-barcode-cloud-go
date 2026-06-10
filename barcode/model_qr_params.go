package barcode

// QrParams - Optional QR barcode generation parameters. Applies to QR, GS1QR, MicroQR, and RectMicroQR barcode types.
type QrParams struct {
	QrEncodeMode  QREncodeMode `json:"qrEncodeMode,omitempty"`
	QrErrorLevel  QRErrorLevel `json:"qrErrorLevel,omitempty"`
	QrVersion     QRVersion    `json:"qrVersion,omitempty"`
	QrECIEncoding ECIEncodings `json:"qrECIEncoding,omitempty"`
	// QR barcode aspect ratio. Values: 0 to 1.
	QrAspectRatio      float32            `json:"qrAspectRatio,omitempty"`
	MicroQRVersion     MicroQRVersion     `json:"microQRVersion,omitempty"`
	RectMicroQrVersion RectMicroQRVersion `json:"rectMicroQrVersion,omitempty"`
}
