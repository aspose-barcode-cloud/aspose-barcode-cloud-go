package barcode

// PresetType : See QualitySettings allows to configure recognition quality and speed manually.
type PresetType string

// List of PresetType
const (
	PresetTypeHighPerformance      PresetType = "HighPerformance"
	PresetTypeNormalQuality        PresetType = "NormalQuality"
	PresetTypeHighQualityDetection PresetType = "HighQualityDetection"
	PresetTypeMaxQualityDetection  PresetType = "MaxQualityDetection"
	PresetTypeHighQuality          PresetType = "HighQuality"
	PresetTypeMaxBarCodes          PresetType = "MaxBarCodes"
)
