package barcode

// CodablockParams - Codablock parameters.
type CodablockParams struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Columns count.
	Columns int32 `json:"Columns,omitempty"`
	// Rows count.
	Rows int32 `json:"Rows,omitempty"`
}
