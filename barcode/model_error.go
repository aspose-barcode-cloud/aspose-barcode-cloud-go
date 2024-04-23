package barcode

// ModelError - Error
type ModelError struct {
	// Code
	Code string `json:"Code,omitempty"`
	// Message
	Message string `json:"Message,omitempty"`
	// Description
	Description string `json:"Description,omitempty"`
	// Inner Error
	InnerError ErrorDetails `json:"InnerError,omitempty"`
}
