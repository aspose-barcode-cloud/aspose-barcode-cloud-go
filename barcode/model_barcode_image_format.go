package barcode

// BarcodeImageFormat : Specifies the file format of the image.
type BarcodeImageFormat string

// List of BarcodeImageFormat
const (
	BarcodeImageFormatGif  BarcodeImageFormat = "Gif"
	BarcodeImageFormatJpeg BarcodeImageFormat = "Jpeg"
	BarcodeImageFormatPng  BarcodeImageFormat = "Png"
	BarcodeImageFormatTiff BarcodeImageFormat = "Tiff"
	BarcodeImageFormatSvg  BarcodeImageFormat = "Svg"
)
