package barcode

// Pdf417Params - Optional PDF417 barcode generation parameters. Applies to Pdf417, MacroPdf417, MicroPdf417, and GS1MicroPdf417 barcode types.
type Pdf417Params struct {
	Pdf417EncodeMode Pdf417EncodeMode `json:"pdf417EncodeMode,omitempty"`
	Pdf417ErrorLevel Pdf417ErrorLevel `json:"pdf417ErrorLevel,omitempty"`
	// Whether to use truncated PDF417 format (removes right-side stop pattern).
	Pdf417Truncate bool `json:"pdf417Truncate,omitempty"`
	// Number of columns in the PDF417 barcode. Values between 1 and 30. 0 for auto.
	Pdf417Columns int32 `json:"pdf417Columns,omitempty"`
	// Number of rows in the PDF417 barcode. Values between 3 and 90. 0 for automatic.
	Pdf417Rows int32 `json:"pdf417Rows,omitempty"`
	// PDF417 barcode aspect ratio (height/width of the barcode module). Values are defined by the standard: 2 to 5 for MicroPdf417; 3 to 5 for Pdf417 and MacroPdf417.
	Pdf417AspectRatio float32      `json:"pdf417AspectRatio,omitempty"`
	Pdf417ECIEncoding ECIEncodings `json:"pdf417ECIEncoding,omitempty"`
	// Whether the barcode is used for reader initialization (programming).
	Pdf417IsReaderInitialization bool           `json:"pdf417IsReaderInitialization,omitempty"`
	Pdf417MacroCharacters        MacroCharacter `json:"pdf417MacroCharacters,omitempty"`
	// Whether to use linked mode (for MicroPdf417).
	Pdf417IsLinked bool `json:"pdf417IsLinked,omitempty"`
	// Whether to use Code128 emulation for MicroPdf417.
	Pdf417IsCode128Emulation bool `json:"pdf417IsCode128Emulation,omitempty"`
}
