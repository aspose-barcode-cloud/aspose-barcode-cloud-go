package barcode

// GeneratorParamsList - Represents list of barcode generators
type GeneratorParamsList struct {
	// List of barcode generators
	BarcodeBuilders []GeneratorParams `json:"BarcodeBuilders,omitempty"`
	// Shift step according to X axis
	XStep int32 `json:"XStep"`
	// Shift step according to Y axis
	YStep int32 `json:"YStep"`
}
