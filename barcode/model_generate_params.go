package barcode

// GenerateParams - Barcode generation parameters
type GenerateParams struct {
	BarcodeType        EncodeBarcodeType  `json:"barcodeType"`
	EncodeData         EncodeData         `json:"encodeData"`
	BarcodeImageParams BarcodeImageParams `json:"barcodeImageParams,omitempty"`
}
