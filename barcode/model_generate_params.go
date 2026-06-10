package barcode

// GenerateParams - Barcode generation parameters.
type GenerateParams struct {
	BarcodeType        EncodeBarcodeType  `json:"barcodeType"`
	EncodeData         EncodeData         `json:"encodeData"`
	BarcodeImageParams BarcodeImageParams `json:"barcodeImageParams,omitempty"`
	QrParams           QrParams           `json:"qrParams,omitempty"`
	Code128Params      Code128Params      `json:"code128Params,omitempty"`
	Pdf417Params       Pdf417Params       `json:"pdf417Params,omitempty"`
}
