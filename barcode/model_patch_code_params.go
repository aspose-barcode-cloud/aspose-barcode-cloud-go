package barcode

// PatchCodeParams - PatchCode parameters.
type PatchCodeParams struct {
	// Specifies codetext for an extra QR barcode, when PatchCode is generated in page mode.
	ExtraBarcodeText string `json:"ExtraBarcodeText,omitempty"`
	// PatchCode format. Choose PatchOnly to generate single PatchCode. Use page format to generate Patch page with PatchCodes as borders. Default value: PatchFormat.PatchOnly
	PatchFormat PatchFormat `json:"PatchFormat,omitempty"`
}
