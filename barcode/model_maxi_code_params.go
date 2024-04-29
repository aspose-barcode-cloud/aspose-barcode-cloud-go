package barcode

// MaxiCodeParams - MaxiCode parameters.
type MaxiCodeParams struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Mode for MaxiCode barcodes.
	Mode MaxiCodeMode `json:"Mode,omitempty"`
	// Encoding mode for MaxiCode barcodes.
	EncodeMode MaxiCodeEncodeMode `json:"EncodeMode,omitempty"`
}
