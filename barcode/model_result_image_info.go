package barcode

// ResultImageInfo - Created image info.
type ResultImageInfo struct {
	// Result file size.
	FileSize int64 `json:"FileSize"`
	// Result image width.
	ImageWidth int32 `json:"ImageWidth,omitempty"`
	// Result image height.
	ImageHeight int32 `json:"ImageHeight,omitempty"`
}
