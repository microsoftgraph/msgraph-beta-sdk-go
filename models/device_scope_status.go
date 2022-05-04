package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceScopeStatus int

const (
    NONE_DEVICESCOPESTATUS DeviceScopeStatus = iota
    COMPUTING_DEVICESCOPESTATUS
    INSUFFICIENTDATA_DEVICESCOPESTATUS
    COMPLETED_DEVICESCOPESTATUS
    UNKNOWNFUTUREVALUE_DEVICESCOPESTATUS
)

func (i DeviceScopeStatus) String() string {
    return []string{"NONE", "COMPUTING", "INSUFFICIENTDATA", "COMPLETED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeviceScopeStatus(v string) (interface{}, error) {
    result := NONE_DEVICESCOPESTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICESCOPESTATUS
        case "COMPUTING":
            result = COMPUTING_DEVICESCOPESTATUS
        case "INSUFFICIENTDATA":
            result = INSUFFICIENTDATA_DEVICESCOPESTATUS
        case "COMPLETED":
            result = COMPLETED_DEVICESCOPESTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPESTATUS
        default:
            return 0, errors.New("Unknown DeviceScopeStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceScopeStatus(values []DeviceScopeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
