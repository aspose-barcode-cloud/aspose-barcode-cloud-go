package barcode

// ApiErrorResponse - ApiError Response
type ApiErrorResponse struct {
	// Gets or sets request Id.
	RequestId string   `json:"requestId,omitempty"`
	Error     ApiError `json:"error,omitempty"`
}
