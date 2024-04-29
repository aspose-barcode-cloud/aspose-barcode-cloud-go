package barcode

// DataMatrixParams - DataMatrix parameters.
type DataMatrixParams struct {
	// Height/Width ratio of 2D BarCode module
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// DEPRECATED: This property is obsolete and will be removed in future releases. Unicode symbols detection and encoding will be processed in Auto mode with Extended Channel Interpretation charset designator. Using of own encodings requires manual CodeText encoding into byte[] array.  Sets the encoding of codetext.
	TextEncoding string `json:"TextEncoding,omitempty"`
	// DEPRECATED: Will be replaced with 'DataMatrix.Version' in the next release  Columns count.
	Columns int32 `json:"Columns,omitempty"`
	// Datamatrix ECC type. Default value: DataMatrixEccType.Ecc200.
	DataMatrixEcc DataMatrixEccType `json:"DataMatrixEcc,omitempty"`
	// Encode mode of Datamatrix barcode. Default value: DataMatrixEncodeMode.Auto.
	DataMatrixEncodeMode DataMatrixEncodeMode `json:"DataMatrixEncodeMode,omitempty"`
	// DEPRECATED: Will be replaced with 'DataMatrix.Version' in the next release  Rows count.
	Rows int32 `json:"Rows,omitempty"`
	// Macro Characters 05 and 06 values are used to obtain more compact encoding in special modes. Can be used only with DataMatrixEccType.Ecc200 or DataMatrixEccType.EccAuto. Cannot be used with EncodeTypes.GS1DataMatrix Default value: MacroCharacters.None.
	MacroCharacters MacroCharacter `json:"MacroCharacters,omitempty"`
	// Sets a Datamatrix symbol size. Default value: DataMatrixVersion.Auto.
	Version DataMatrixVersion `json:"Version,omitempty"`
}
