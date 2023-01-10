package models
import (
    "errors"
)
// Indicates the status of the attempted device scope action
type DeviceScopeActionStatus int

const (
    // Indicates the device scope action failed to trigger.
    FAILED_DEVICESCOPEACTIONSTATUS DeviceScopeActionStatus = iota
    // Indicates the device scope action was successfully triggered.
    SUCCEEDED_DEVICESCOPEACTIONSTATUS
    // Placeholder value for future expansion.
    UNKNOWNFUTUREVALUE_DEVICESCOPEACTIONSTATUS
)

func (i DeviceScopeActionStatus) String() string {
    return []string{"failed", "succeeded", "unknownFutureValue"}[i]
}
func ParseDeviceScopeActionStatus(v string) (interface{}, error) {
    result := FAILED_DEVICESCOPEACTIONSTATUS
    switch v {
        case "failed":
            result = FAILED_DEVICESCOPEACTIONSTATUS
        case "succeeded":
            result = SUCCEEDED_DEVICESCOPEACTIONSTATUS
        case "unknownFutureValue":
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
