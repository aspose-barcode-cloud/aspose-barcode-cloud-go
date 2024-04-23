package barcode

// FontParams - Font.
type FontParams struct {
	// Font family.
	Family string `json:"Family,omitempty"`
	// Font size in units.
	Size float64 `json:"Size,omitempty"`
	// Font style.
	Style FontStyle `json:"Style,omitempty"`
}
