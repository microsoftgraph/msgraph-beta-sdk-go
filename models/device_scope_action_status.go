package models
import (
    "strings"
    "errors"
)
// Provides operations to call the triggerDeviceScopeAction method.
type DeviceScopeActionStatus int

const (
    FAILED_DEVICESCOPEACTIONSTATUS DeviceScopeActionStatus = iota
    SUCCEEDED_DEVICESCOPEACTIONSTATUS
    UNKNOWNFUTUREVALUE_DEVICESCOPEACTIONSTATUS
)

func (i DeviceScopeActionStatus) String() string {
    return []string{"FAILED", "SUCCEEDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeviceScopeActionStatus(v string) (interface{}, error) {
    result := FAILED_DEVICESCOPEACTIONSTATUS
    switch strings.ToUpper(v) {
        case "FAILED":
            result = FAILED_DEVICESCOPEACTIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_DEVICESCOPEACTIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPEACTIONSTATUS
        default:
            return 0, errors.New("Unknown DeviceScopeActionStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceScopeActionStatus(values []DeviceScopeActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
