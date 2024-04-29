package barcode

// AvailableGraphicsUnit : Subset of GraphicsUnit.
type AvailableGraphicsUnit string

// List of AvailableGraphicsUnit
const (
	AvailableGraphicsUnitPixel      AvailableGraphicsUnit = "Pixel"
	AvailableGraphicsUnitPoint      AvailableGraphicsUnit = "Point"
	AvailableGraphicsUnitInch       AvailableGraphicsUnit = "Inch"
	AvailableGraphicsUnitMillimeter AvailableGraphicsUnit = "Millimeter"
)
