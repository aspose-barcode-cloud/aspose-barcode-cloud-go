package barcode

// DiscUsage - Class for disc space information.
type DiscUsage struct {
	// Application used disc space.
	UsedSize int64 `json:"UsedSize"`
	// Total disc space.
	TotalSize int64 `json:"TotalSize"`
}
