package barcode

// AustralianPostParams - AustralianPost barcode parameters.
type AustralianPostParams struct {
	// Interpreting type for the Customer Information of AustralianPost, default to CustomerInformationInterpretingType.Other
	EncodingTable CustomerInformationInterpretingType `json:"EncodingTable,omitempty"`
	// Short bar's height of AustralianPost barcode.
	ShortBarHeight float64 `json:"ShortBarHeight,omitempty"`
}
