package barcode

// RegionPoint - Wrapper around Drawing.Point for proper specification.
type RegionPoint struct {
	// X-coordinate
	X int32 `json:"X"`
	// Y-coordinate
	Y int32 `json:"Y"`
}
