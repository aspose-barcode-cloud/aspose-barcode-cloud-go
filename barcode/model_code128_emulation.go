package barcode

// Code128Emulation : DEPRECATED. This enum will be removed in future releases Function codewords for Code 128 emulation. Applied for MicroPDF417 only. Ignored for PDF417 and MacroPDF417 barcodes.
type Code128Emulation string

// List of Code128Emulation
const (
	Code128EmulationNone    Code128Emulation = "None"
	Code128EmulationCode903 Code128Emulation = "Code903"
	Code128EmulationCode904 Code128Emulation = "Code904"
	Code128EmulationCode905 Code128Emulation = "Code905"
)
