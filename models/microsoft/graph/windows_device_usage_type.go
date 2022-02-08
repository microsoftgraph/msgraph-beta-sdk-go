package graph
import (
    "strings"
    "errors"
)
// 
type WindowsDeviceUsageType int

const (
    SINGLEUSER_WINDOWSDEVICEUSAGETYPE WindowsDeviceUsageType = iota
    SHARED_WINDOWSDEVICEUSAGETYPE
)

func (i WindowsDeviceUsageType) String() string {
    return []string{"SINGLEUSER", "SHARED"}[i]
}
func ParseWindowsDeviceUsageType(v string) (interface{}, error) {
    result := SINGLEUSER_WINDOWSDEVICEUSAGETYPE
    switch strings.ToUpper(v) {
        case "SINGLEUSER":
            result = SINGLEUSER_WINDOWSDEVICEUSAGETYPE
        case "SHARED":
            result = SHARED_WINDOWSDEVICEUSAGETYPE
        default:
            return 0, errors.New("Unknown WindowsDeviceUsageType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDeviceUsageType(values []WindowsDeviceUsageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
