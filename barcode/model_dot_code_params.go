package barcode

// DotCodeParams - DotCode parameters.
type DotCodeParams struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Identifies columns count. Sum of the number of rows plus the number of columns of a DotCode symbol must be odd. Number of columns must be at least 5.
	Columns int32 `json:"Columns,omitempty"`
	// Identifies DotCode encode mode. Default value: Auto.
	EncodeMode DotCodeEncodeMode `json:"EncodeMode,omitempty"`
	// Identifies ECI encoding. Used when DotCodeEncodeMode is Auto. Default value: ISO-8859-1.
	ECIEncoding EciEncodings `json:"ECIEncoding,omitempty"`
	// Indicates whether code is used for instruct reader to interpret the following data as instructions for initialization or reprogramming of the bar code reader. Default value is false.
	IsReaderInitialization bool `json:"IsReaderInitialization,omitempty"`
	// Identifies rows count. Sum of the number of rows plus the number of columns of a DotCode symbol must be odd. Number of rows must be at least 5.
	Rows int32 `json:"Rows,omitempty"`
}
