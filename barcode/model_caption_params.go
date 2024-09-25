package barcode

// CaptionParams - Caption
type CaptionParams struct {
	// Caption text.
	Text string `json:"Text,omitempty"`
	// Text alignment.
	Alignment TextAlignment `json:"Alignment,omitempty"`
	// Text color.   Default value: black   Use named colors like: red, green, blue   Or HTML colors like: #FF0000, #00FF00, #0000FF
	Color string `json:"Color,omitempty"`
	// Is caption visible.
	Visible bool `json:"Visible,omitempty"`
	// Font.
	Font FontParams `json:"Font,omitempty"`
	// Padding.
	Padding Padding `json:"Padding,omitempty"`
	// Specify word wraps (line breaks) within text. Default value: false.
	NoWrap bool `json:"NoWrap,omitempty"`
}
