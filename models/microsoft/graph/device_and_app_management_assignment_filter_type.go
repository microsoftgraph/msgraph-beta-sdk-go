package graph
import (
    "strings"
    "errors"
)
// 
type DeviceAndAppManagementAssignmentFilterType int

const (
    NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE DeviceAndAppManagementAssignmentFilterType = iota
    INCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
    EXCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
)

func (i DeviceAndAppManagementAssignmentFilterType) String() string {
    return []string{"NONE", "INCLUDE", "EXCLUDE"}[i]
}
func ParseDeviceAndAppManagementAssignmentFilterType(v string) (interface{}, error) {
    result := NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
        case "INCLUDE":
            result = INCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
        case "EXCLUDE":
            result = EXCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
        default:
            return 0, errors.New("Unknown DeviceAndAppManagementAssignmentFilterType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAndAppManagementAssignmentFilterType(values []DeviceAndAppManagementAssignmentFilterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
