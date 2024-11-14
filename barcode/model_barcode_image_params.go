package barcode

// BarcodeImageParams - Barcode image optional parameters
type BarcodeImageParams struct {
	ImageFormat  BarcodeImageFormat `json:"imageFormat,omitempty"`
	TextLocation CodeLocation       `json:"textLocation,omitempty"`
	// Specify the displaying bars and content Color.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: Black.
	ForegroundColor string `json:"foregroundColor,omitempty"`
	// Background color of the barcode image.  Value: Color name from https://reference.aspose.com/drawing/net/system.drawing/color/ or ARGB value started with #.  For example: AliceBlue or #FF000000  Default value: White.
	BackgroundColor string       `json:"backgroundColor,omitempty"`
	Units           GraphicsUnit `json:"units,omitempty"`
	// Resolution of the BarCode image.  One value for both dimensions.  Default value: 96 dpi.  Decimal separator is dot.
	Resolution float32 `json:"resolution,omitempty"`
	// Height of the barcode image in given units. Default units: pixel.  Decimal separator is dot.
	ImageHeight float32 `json:"imageHeight,omitempty"`
	// Width of the barcode image in given units. Default units: pixel.  Decimal separator is dot.
	ImageWidth float32 `json:"imageWidth,omitempty"`
	// BarCode image rotation angle, measured in degree, e.g. RotationAngle = 0 or RotationAngle = 360 means no rotation.  If RotationAngle NOT equal to 90, 180, 270 or 0, it may increase the difficulty for the scanner to read the image.  Default value: 0.
	RotationAngle int32 `json:"rotationAngle,omitempty"`
}
