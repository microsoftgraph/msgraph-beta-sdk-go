package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEAPPMANAGEMENTTASKPRIORITY, nil
        case "HIGH":
            return HIGH_DEVICEAPPMANAGEMENTTASKPRIORITY, nil
        case "LOW":
            return LOW_DEVICEAPPMANAGEMENTTASKPRIORITY, nil
    }
    return 0, errors.New("Unknown DeviceAppManagementTaskPriority value: " + v)
}
func SerializeDeviceAppManagementTaskPriority(values []DeviceAppManagementTaskPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
