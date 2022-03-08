package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the compare method.
type DeviceManagementComparisonResult int

const (
    UNKNOWN_DEVICEMANAGEMENTCOMPARISONRESULT DeviceManagementComparisonResult = iota
    EQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
    NOTEQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
    ADDED_DEVICEMANAGEMENTCOMPARISONRESULT
    REMOVED_DEVICEMANAGEMENTCOMPARISONRESULT
)

func (i DeviceManagementComparisonResult) String() string {
    return []string{"UNKNOWN", "EQUAL", "NOTEQUAL", "ADDED", "REMOVED"}[i]
}
func ParseDeviceManagementComparisonResult(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTCOMPARISONRESULT
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTCOMPARISONRESULT
        case "EQUAL":
            result = EQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
        case "NOTEQUAL":
            result = NOTEQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
        case "ADDED":
            result = ADDED_DEVICEMANAGEMENTCOMPARISONRESULT
        case "REMOVED":
            result = REMOVED_DEVICEMANAGEMENTCOMPARISONRESULT
        default:
            return 0, errors.New("Unknown DeviceManagementComparisonResult value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementComparisonResult(values []DeviceManagementComparisonResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
