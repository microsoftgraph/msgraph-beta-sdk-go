// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Windows 10 Device Mode type.
type Windows10DeviceModeType int

const (
    // Standard Configuration
    STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE Windows10DeviceModeType = iota
    // S Mode Configuration
    SMODECONFIGURATION_WINDOWS10DEVICEMODETYPE
)

func (i Windows10DeviceModeType) String() string {
    return []string{"standardConfiguration", "sModeConfiguration"}[i]
}
func ParseWindows10DeviceModeType(v string) (any, error) {
    result := STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE
    switch v {
        case "standardConfiguration":
            result = STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE
        case "sModeConfiguration":
            result = SMODECONFIGURATION_WINDOWS10DEVICEMODETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindows10DeviceModeType(values []Windows10DeviceModeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Windows10DeviceModeType) isMultiValue() bool {
    return false
}
