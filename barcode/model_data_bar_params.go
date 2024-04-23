package barcode

// DataBarParams - Databar parameters.
type DataBarParams struct {
	// Height/Width ratio of 2D BarCode module. Used for DataBar stacked.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Columns count.
	Columns int32 `json:"Columns,omitempty"`
	// Rows count.
	Rows int32 `json:"Rows,omitempty"`
	// Enables flag of 2D composite component with DataBar barcode
	Is2DCompositeComponent bool `json:"Is2DCompositeComponent,omitempty"`
	// If this flag is set, it allows only GS1 encoding standard for Databar barcode types
	IsAllowOnlyGS1Encoding bool `json:"IsAllowOnlyGS1Encoding,omitempty"`
}
