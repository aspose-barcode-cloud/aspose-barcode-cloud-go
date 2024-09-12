package barcode

// RegionPoint - Wrapper around Drawing.Point for proper specification.
type RegionPoint struct {
	// X-coordinate
	X int32 `json:"x,omitempty"`
	// Y-coordinate
	Y int32 `json:"y,omitempty"`
}
