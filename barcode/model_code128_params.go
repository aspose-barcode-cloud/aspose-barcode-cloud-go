package barcode

// Code128Params - Code128 params.
type Code128Params struct {
	// Encoding mode for Code128 barcodes. Code 128 specification Default value: Code128EncodeMode.Auto.
	EncodeMode Code128EncodeMode `json:"EncodeMode,omitempty"`
}
