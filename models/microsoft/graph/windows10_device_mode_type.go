package graph
import (
    "strings"
    "errors"
)
type Windows10DeviceModeType int

const (
    STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE Windows10DeviceModeType = iota
    SMODECONFIGURATION_WINDOWS10DEVICEMODETYPE
)

func (i Windows10DeviceModeType) String() string {
    return []string{"STANDARDCONFIGURATION", "SMODECONFIGURATION"}[i]
}
func ParseWindows10DeviceModeType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "STANDARDCONFIGURATION":
            return STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE, nil
        case "SMODECONFIGURATION":
            return SMODECONFIGURATION_WINDOWS10DEVICEMODETYPE, nil
    }
    return 0, errors.New("Unknown Windows10DeviceModeType value: " + v)
}
func SerializeWindows10DeviceModeType(values []Windows10DeviceModeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
