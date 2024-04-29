package barcode

// HanXinEncodeMode :
type HanXinEncodeMode string

// List of HanXinEncodeMode
const (
	HanXinEncodeModeAuto     HanXinEncodeMode = "Auto"
	HanXinEncodeModeBinary   HanXinEncodeMode = "Binary"
	HanXinEncodeModeECI      HanXinEncodeMode = "ECI"
	HanXinEncodeModeUnicode  HanXinEncodeMode = "Unicode"
	HanXinEncodeModeURI      HanXinEncodeMode = "URI"
	HanXinEncodeModeExtended HanXinEncodeMode = "Extended"
)
