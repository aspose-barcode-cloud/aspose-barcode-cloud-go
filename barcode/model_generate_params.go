package barcode

// GenerateParams - Barcode generation parameters.
type GenerateParams struct {
	// Barcode type.
	BarcodeType EncodeBarcodeType `json:"barcodeType"`
	// Data to encode into a barcode.
	EncodeData EncodeData `json:"encodeData"`
	// Optional barcode image parameters.
	BarcodeImageParams BarcodeImageParams `json:"barcodeImageParams,omitempty"`
	// Optional QR barcode generation parameters.
	QrParams QrParams `json:"qrParams,omitempty"`
	// Optional Code128 barcode generation parameters.
	Code128Params Code128Params `json:"code128Params,omitempty"`
	// Optional PDF417 barcode generation parameters.
	Pdf417Params Pdf417Params `json:"pdf417Params,omitempty"`
}
