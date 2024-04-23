package barcode

// ApiErrorResponse -
type ApiErrorResponse struct {
	RequestId string   `json:"RequestId,omitempty"`
	Error     ApiError `json:"Error,omitempty"`
}
