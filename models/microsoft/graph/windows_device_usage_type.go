package graph
import (
    "strings"
    "errors"
)
type WindowsDeviceUsageType int

const (
    SINGLEUSER_WINDOWSDEVICEUSAGETYPE WindowsDeviceUsageType = iota
    SHARED_WINDOWSDEVICEUSAGETYPE
)

func (i WindowsDeviceUsageType) String() string {
    return []string{"SINGLEUSER", "SHARED"}[i]
}
func ParseWindowsDeviceUsageType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SINGLEUSER":
            return SINGLEUSER_WINDOWSDEVICEUSAGETYPE, nil
        case "SHARED":
            return SHARED_WINDOWSDEVICEUSAGETYPE, nil
    }
    return 0, errors.New("Unknown WindowsDeviceUsageType value: " + v)
}
func SerializeWindowsDeviceUsageType(values []WindowsDeviceUsageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
