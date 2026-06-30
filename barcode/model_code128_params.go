package barcode

// Code128Params - Optional Code128 barcode generation parameters.
type Code128Params struct {
	// Code128 barcode encode mode. Controls which Code 128 subset (A, B, C, or mix) is used.
	Code128EncodeMode Code128EncodeMode `json:"code128EncodeMode,omitempty"`
}
