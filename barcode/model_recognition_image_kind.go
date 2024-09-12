package barcode

// RecognitionImageKind : Kind of image to recognize
type RecognitionImageKind string

// List of RecognitionImageKind
const (
	RecognitionImageKindPhoto           RecognitionImageKind = "Photo"
	RecognitionImageKindScannedDocument RecognitionImageKind = "ScannedDocument"
	RecognitionImageKindClearImage      RecognitionImageKind = "ClearImage"
)
