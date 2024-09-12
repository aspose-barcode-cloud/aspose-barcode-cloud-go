package barcode

// AvailableBarCodeImageFormat : Specifies the file format of the image.
type AvailableBarCodeImageFormat string

// List of AvailableBarCodeImageFormat
const (
	AvailableBarCodeImageFormatGif  AvailableBarCodeImageFormat = Gif
	AvailableBarCodeImageFormatJpeg AvailableBarCodeImageFormat = Jpeg
	AvailableBarCodeImageFormatPng  AvailableBarCodeImageFormat = Png
	AvailableBarCodeImageFormatTiff AvailableBarCodeImageFormat = Tiff
	AvailableBarCodeImageFormatSvg  AvailableBarCodeImageFormat = Svg
)
