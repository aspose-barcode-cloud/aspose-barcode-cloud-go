/*
 * Aspose.Barcode Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose_barcode_cloud

import (
	"time"
)

// File or folder information
type StorageFile struct {
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
}
