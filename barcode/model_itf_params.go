package barcode

// ItfParams - ITF parameters.
type ItfParams struct {
	// ITF border (bearer bar) thickness in Unit value. Default value: 12pt.
	BorderThickness float64 `json:"BorderThickness,omitempty"`
	// Border type of ITF barcode. Default value: ITF14BorderType.Bar.
	BorderType Itf14BorderType `json:"BorderType,omitempty"`
	// Size of the quiet zones in xDimension. Default value: 10, meaning if xDimension = 2px than quiet zones will be 20px.
	QuietZoneCoef int32 `json:"QuietZoneCoef,omitempty"`
}
