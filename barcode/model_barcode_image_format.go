package barcode

// BarcodeImageFormat : Specifies the file format of the image.
type BarcodeImageFormat string

// List of BarcodeImageFormat
const (
	BarcodeImageFormatPng  BarcodeImageFormat = "Png"
	BarcodeImageFormatJpeg BarcodeImageFormat = "Jpeg"
	BarcodeImageFormatSvg  BarcodeImageFormat = "Svg"
	BarcodeImageFormatTiff BarcodeImageFormat = "Tiff"
	BarcodeImageFormatGif  BarcodeImageFormat = "Gif"
)
