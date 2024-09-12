package barcode

// ApiErrorResponse - ApiError Response
type ApiErrorResponse struct {
	// Gets or sets request Id.
	RequestId NullableString `json:"requestId,omitempty"`
	Error     ApiError       `json:"error,omitempty"`
}
