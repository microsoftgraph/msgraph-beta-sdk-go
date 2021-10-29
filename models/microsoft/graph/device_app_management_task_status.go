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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS, nil
        case "PENDING":
            return PENDING_DEVICEAPPMANAGEMENTTASKSTATUS, nil
        case "ACTIVE":
            return ACTIVE_DEVICEAPPMANAGEMENTTASKSTATUS, nil
        case "COMPLETED":
            return COMPLETED_DEVICEAPPMANAGEMENTTASKSTATUS, nil
        case "REJECTED":
            return REJECTED_DEVICEAPPMANAGEMENTTASKSTATUS, nil
    }
    return 0, errors.New("Unknown DeviceAppManagementTaskStatus value: " + v)
}
func SerializeDeviceAppManagementTaskStatus(values []DeviceAppManagementTaskStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
