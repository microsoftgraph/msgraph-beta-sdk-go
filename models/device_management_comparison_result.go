package models
import (
    "errors"
)
// Provides operations to call the compare method.
type DeviceManagementComparisonResult int

const (
    // Unknown setting comparison
    UNKNOWN_DEVICEMANAGEMENTCOMPARISONRESULT DeviceManagementComparisonResult = iota
    // The setting values are equal
    EQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
    // The setting values are not equal
    NOTEQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
    // The setting is added
    ADDED_DEVICEMANAGEMENTCOMPARISONRESULT
    // The setting is removed
    REMOVED_DEVICEMANAGEMENTCOMPARISONRESULT
)

func (i DeviceManagementComparisonResult) String() string {
    return []string{"unknown", "equal", "notEqual", "added", "removed"}[i]
}
func ParseDeviceManagementComparisonResult(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTCOMPARISONRESULT
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEMANAGEMENTCOMPARISONRESULT
        case "equal":
            result = EQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
        case "notEqual":
            result = NOTEQUAL_DEVICEMANAGEMENTCOMPARISONRESULT
        case "added":
            result = ADDED_DEVICEMANAGEMENTCOMPARISONRESULT
        case "removed":
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
