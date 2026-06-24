package barcode

// ApiErrorResponse - ApiError Response
type ApiErrorResponse struct {
	// Gets or sets request Id.
	RequestId string `json:"requestId"`
	// Gets or sets error.
	Error ApiError `json:"error"`
}
