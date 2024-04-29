package barcode

// FilesUploadResult - File upload result
type FilesUploadResult struct {
	// List of uploaded file names
	Uploaded []string `json:"Uploaded,omitempty"`
	// List of errors.
	Errors []ModelError `json:"Errors,omitempty"`
}
