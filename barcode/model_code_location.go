package barcode

type CodeLocation string

// List of CodeLocation
const (
	CodeLocationBelow CodeLocation = "Below"
	CodeLocationAbove CodeLocation = "Above"
	CodeLocationNone  CodeLocation = "None"
)
