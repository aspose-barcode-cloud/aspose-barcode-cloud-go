package barcode

// PatchFormat :
type PatchFormat string

// List of PatchFormat
const (
	PatchFormatPatchOnly           PatchFormat = "PatchOnly"
	PatchFormatA4                  PatchFormat = "A4"
	PatchFormatA4_LANDSCAPE        PatchFormat = "A4_LANDSCAPE"
	PatchFormatUS_Letter           PatchFormat = "US_Letter"
	PatchFormatUS_Letter_LANDSCAPE PatchFormat = "US_Letter_LANDSCAPE"
)
