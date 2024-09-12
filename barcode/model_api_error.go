package barcode

// ApiError - Api Error.
type ApiError struct {
	// Gets or sets api error code.
	Code NullableString `json:"code"`
	// Gets or sets error message.
	Message NullableString `json:"message"`
	// Gets or sets error description.
	Description NullableString `json:"description,omitempty"`
	// Gets or sets server datetime.
	DateTime   NullableTime `json:"dateTime,omitempty"`
	InnerError *ApiError    `json:"innerError,omitempty"`
}
