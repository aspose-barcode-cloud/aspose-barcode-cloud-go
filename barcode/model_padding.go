package barcode

// Padding - Padding around barcode.
type Padding struct {
	// Left padding.
	Left float64 `json:"Left,omitempty"`
	// Right padding.
	Right float64 `json:"Right,omitempty"`
	// Top padding.
	Top float64 `json:"Top,omitempty"`
	// Bottom padding.
	Bottom float64 `json:"Bottom,omitempty"`
}
