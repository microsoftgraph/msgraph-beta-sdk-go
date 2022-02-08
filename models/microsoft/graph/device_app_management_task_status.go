package graph
import (
    "strings"
    "errors"
)
// 
type DeviceAppManagementTaskStatus int

const (
    UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS DeviceAppManagementTaskStatus = iota
    PENDING_DEVICEAPPMANAGEMENTTASKSTATUS
    ACTIVE_DEVICEAPPMANAGEMENTTASKSTATUS
    COMPLETED_DEVICEAPPMANAGEMENTTASKSTATUS
    REJECTED_DEVICEAPPMANAGEMENTTASKSTATUS
)

func (i DeviceAppManagementTaskStatus) String() string {
    return []string{"UNKNOWN", "PENDING", "ACTIVE", "COMPLETED", "REJECTED"}[i]
}
func ParseDeviceAppManagementTaskStatus(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS
        case "PENDING":
            result = PENDING_DEVICEAPPMANAGEMENTTASKSTATUS
        case "ACTIVE":
            result = ACTIVE_DEVICEAPPMANAGEMENTTASKSTATUS
        case "COMPLETED":
            result = COMPLETED_DEVICEAPPMANAGEMENTTASKSTATUS
        case "REJECTED":
            result = REJECTED_DEVICEAPPMANAGEMENTTASKSTATUS
        default:
            return 0, errors.New("Unknown DeviceAppManagementTaskStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAppManagementTaskStatus(values []DeviceAppManagementTaskStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
