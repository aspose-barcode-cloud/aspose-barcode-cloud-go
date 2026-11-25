package barcode

// BatchArchiveFormat : Archive format for batch barcode generation.
type BatchArchiveFormat string

// List of BatchArchiveFormat
const (
	BatchArchiveFormatZip     BatchArchiveFormat = "Zip"
	BatchArchiveFormatTar     BatchArchiveFormat = "Tar"
	BatchArchiveFormatTarGzip BatchArchiveFormat = "TarGzip"
)
