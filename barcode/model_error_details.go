package barcode

import (
	"time"
)

// ErrorDetails - The error details
type ErrorDetails struct {
	// The request id
	RequestId string `json:"RequestId,omitempty"`
	// Date
	Date time.Time `json:"Date"`
}
