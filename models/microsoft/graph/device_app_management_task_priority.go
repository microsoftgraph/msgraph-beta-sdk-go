package graph
import (
    "strings"
    "errors"
)
// 
type DeviceAppManagementTaskPriority int

const (
    NONE_DEVICEAPPMANAGEMENTTASKPRIORITY DeviceAppManagementTaskPriority = iota
    HIGH_DEVICEAPPMANAGEMENTTASKPRIORITY
    LOW_DEVICEAPPMANAGEMENTTASKPRIORITY
)

func (i DeviceAppManagementTaskPriority) String() string {
    return []string{"NONE", "HIGH", "LOW"}[i]
}
func ParseDeviceAppManagementTaskPriority(v string) (interface{}, error) {
    result := NONE_DEVICEAPPMANAGEMENTTASKPRIORITY
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEAPPMANAGEMENTTASKPRIORITY
        case "HIGH":
            result = HIGH_DEVICEAPPMANAGEMENTTASKPRIORITY
        case "LOW":
            result = LOW_DEVICEAPPMANAGEMENTTASKPRIORITY
        default:
            return 0, errors.New("Unknown DeviceAppManagementTaskPriority value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAppManagementTaskPriority(values []DeviceAppManagementTaskPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
