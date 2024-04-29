package barcode

import (
	"time"
)

// FileVersion -
type FileVersion struct {
	// File or folder name.
	Name string `json:"Name,omitempty"`
	// True if it is a folder.
	IsFolder bool `json:"IsFolder"`
	// File or folder last modified DateTime.
	ModifiedDate time.Time `json:"ModifiedDate,omitempty"`
	// File or folder size.
	Size int64 `json:"Size"`
	// File or folder path.
	Path string `json:"Path,omitempty"`
	// File Version ID.
	VersionId string `json:"VersionId,omitempty"`
	// Specifies whether the file is (true) or is not (false) the latest version of an file.
	IsLatest bool `json:"IsLatest"`
}
