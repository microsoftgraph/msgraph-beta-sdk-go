package graph
import (
    "strings"
    "errors"
)
// 
type DeviceAppManagementTaskCategory int

const (
    UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY DeviceAppManagementTaskCategory = iota
    ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
)

func (i DeviceAppManagementTaskCategory) String() string {
    return []string{"UNKNOWN", "ADVANCEDTHREATPROTECTION"}[i]
}
func ParseDeviceAppManagementTaskCategory(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY
        case "ADVANCEDTHREATPROTECTION":
            result = ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
        default:
            return 0, errors.New("Unknown DeviceAppManagementTaskCategory value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAppManagementTaskCategory(values []DeviceAppManagementTaskCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
