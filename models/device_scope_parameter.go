package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceScopeParameter int

const (
    NONE_DEVICESCOPEPARAMETER DeviceScopeParameter = iota
    SCOPETAG_DEVICESCOPEPARAMETER
    UNKNOWNFUTUREVALUE_DEVICESCOPEPARAMETER
)

func (i DeviceScopeParameter) String() string {
    return []string{"NONE", "SCOPETAG", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeviceScopeParameter(v string) (interface{}, error) {
    result := NONE_DEVICESCOPEPARAMETER
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICESCOPEPARAMETER
        case "SCOPETAG":
            result = SCOPETAG_DEVICESCOPEPARAMETER
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPEPARAMETER
        default:
            return 0, errors.New("Unknown DeviceScopeParameter value: " + v)
    }
    return &result, nil
}
func SerializeDeviceScopeParameter(values []DeviceScopeParameter) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
