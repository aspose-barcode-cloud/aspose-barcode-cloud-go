package barcode

// BarcodeImageParams - Optional barcode image parameters.
type BarcodeImageParams struct {
	ImageFormat  BarcodeImageFormat `json:"imageFormat,omitempty"`
	TextLocation CodeLocation       `json:"textLocation,omitempty"`
	// Specify the display color for bars and content. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: Black.
	ForegroundColor string `json:"foregroundColor,omitempty"`
	// Background color of the barcode image. Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value starting with #. For example: AliceBlue or #FF000000. Default value: White.
	BackgroundColor string       `json:"backgroundColor,omitempty"`
	Units           GraphicsUnit `json:"units,omitempty"`
	// Resolution of the barcode image. One value for both dimensions. Default value: 96 dpi. Decimal separator is a dot.
	Resolution float32 `json:"resolution,omitempty"`
	// Height of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
	ImageHeight float32 `json:"imageHeight,omitempty"`
	// Width of the barcode image in the specified units. Default units: pixels. Decimal separator is a dot.
	ImageWidth float32 `json:"imageWidth,omitempty"`
	// Barcode image rotation angle, measured in degrees. For example, RotationAngle = 0 or RotationAngle = 360 means no rotation. If RotationAngle is not equal to 90, 180, 270, or 0, it may increase the difficulty for the scanner to read the image. Default value: 0.
	RotationAngle int32 `json:"rotationAngle,omitempty"`
}
