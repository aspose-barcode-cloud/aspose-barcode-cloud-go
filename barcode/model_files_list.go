package barcode

// FilesList - Files list
type FilesList struct {
	// Files and folders contained by folder StorageFile.
	Value []StorageFile `json:"Value,omitempty"`
}
