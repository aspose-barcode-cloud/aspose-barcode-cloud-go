package barcode

import (
	"time"
)

// ApiError -
type ApiError struct {
	Code        string    `json:"Code,omitempty"`
	Message     string    `json:"Message,omitempty"`
	Description string    `json:"Description,omitempty"`
	DateTime    time.Time `json:"DateTime,omitempty"`
	InnerError  *ApiError `json:"InnerError,omitempty"`
}
