package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE, nil
        case "INCLUDE":
            return INCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE, nil
        case "EXCLUDE":
            return EXCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE, nil
    }
    return 0, errors.New("Unknown DeviceAndAppManagementAssignmentFilterType value: " + v)
}
func SerializeDeviceAndAppManagementAssignmentFilterType(values []DeviceAndAppManagementAssignmentFilterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
