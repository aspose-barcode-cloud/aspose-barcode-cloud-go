package barcode

// ObjectExist - Object exists
type ObjectExist struct {
	// Indicates that the file or folder exists.
	Exists bool `json:"Exists"`
	// True if it is a folder, false if it is a file.
	IsFolder bool `json:"IsFolder"`
}
