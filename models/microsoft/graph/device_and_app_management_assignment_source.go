package graph
import (
    "strings"
    "errors"
)
// 
type DeviceAndAppManagementAssignmentSource int

const (
    DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE DeviceAndAppManagementAssignmentSource = iota
    POLICYSETS_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
)

func (i DeviceAndAppManagementAssignmentSource) String() string {
    return []string{"DIRECT", "POLICYSETS"}[i]
}
func ParseDeviceAndAppManagementAssignmentSource(v string) (interface{}, error) {
    result := DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
    switch strings.ToUpper(v) {
        case "DIRECT":
            result = DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
        case "POLICYSETS":
            result = POLICYSETS_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
        default:
            return 0, errors.New("Unknown DeviceAndAppManagementAssignmentSource value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAndAppManagementAssignmentSource(values []DeviceAndAppManagementAssignmentSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
