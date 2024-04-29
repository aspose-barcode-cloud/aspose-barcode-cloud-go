package barcode

// Code16KParams - Code16K parameters.
type Code16KParams struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Size of the left quiet zone in xDimension. Default value: 10, meaning if xDimension = 2px than left quiet zone will be 20px.
	QuietZoneLeftCoef int32 `json:"QuietZoneLeftCoef,omitempty"`
	// Size of the right quiet zone in xDimension. Default value: 1, meaning if xDimension = 2px than right quiet zone will be 2px.
	QuietZoneRightCoef int32 `json:"QuietZoneRightCoef,omitempty"`
}
