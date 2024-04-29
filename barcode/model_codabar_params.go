package barcode

// CodabarParams - Codabar parameters.
type CodabarParams struct {
	// Checksum algorithm for Codabar barcodes. Default value: CodabarChecksumMode.Mod16. To enable checksum calculation set value EnableChecksum.Yes to property EnableChecksum.
	ChecksumMode CodabarChecksumMode `json:"ChecksumMode,omitempty"`
	// Start symbol (character) of Codabar symbology. Default value: CodabarSymbol.A
	StartSymbol CodabarSymbol `json:"StartSymbol,omitempty"`
	// Stop symbol (character) of Codabar symbology. Default value: CodabarSymbol.A
	StopSymbol CodabarSymbol `json:"StopSymbol,omitempty"`
}
