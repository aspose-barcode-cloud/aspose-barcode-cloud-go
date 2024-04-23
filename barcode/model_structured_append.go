package barcode

// StructuredAppend - QR structured append parameters.
type StructuredAppend struct {
	// The index of the QR structured append mode barcode. Index starts from 0.
	SequenceIndicator int32 `json:"SequenceIndicator,omitempty"`
	// QR structured append mode barcodes quantity. Max value is 16.
	TotalCount int32 `json:"TotalCount,omitempty"`
	// QR structured append mode parity data.
	ParityByte int32 `json:"ParityByte,omitempty"`
}
