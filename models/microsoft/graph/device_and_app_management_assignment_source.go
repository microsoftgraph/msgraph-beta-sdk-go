package graph
import (
    "strings"
    "errors"
)
type DeviceAndAppManagementAssignmentSource int

const (
    DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE DeviceAndAppManagementAssignmentSource = iota
    POLICYSETS_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
)

func (i DeviceAndAppManagementAssignmentSource) String() string {
    return []string{"DIRECT", "POLICYSETS"}[i]
}
func ParseDeviceAndAppManagementAssignmentSource(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DIRECT":
            return DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE, nil
        case "POLICYSETS":
            return POLICYSETS_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE, nil
    }
    return 0, errors.New("Unknown DeviceAndAppManagementAssignmentSource value: " + v)
}
func SerializeDeviceAndAppManagementAssignmentSource(values []DeviceAndAppManagementAssignmentSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
