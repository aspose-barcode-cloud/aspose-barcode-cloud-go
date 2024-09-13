package barcode

import (
	"time"
)

// ApiError - Api Error.
type ApiError struct {
	// Gets or sets api error code.
	Code string `json:"code"`
	// Gets or sets error message.
	Message string `json:"message"`
	// Gets or sets error description.
	Description string `json:"description,omitempty"`
	// Gets or sets server datetime.
	DateTime   time.Time `json:"dateTime,omitempty"`
	InnerError *ApiError `json:"innerError,omitempty"`
}
