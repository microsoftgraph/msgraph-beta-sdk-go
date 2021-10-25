package graph
import (
    "strings"
    "errors"
)
type DeviceAppManagementTaskCategory int

const (
    UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY DeviceAppManagementTaskCategory = iota
    ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
)

func (i DeviceAppManagementTaskCategory) String() string {
    return []string{"UNKNOWN", "ADVANCEDTHREATPROTECTION"}[i]
}
func ParseDeviceAppManagementTaskCategory(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY, nil
        case "ADVANCEDTHREATPROTECTION":
            return ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY, nil
    }
    return 0, errors.New("Unknown DeviceAppManagementTaskCategory value: " + v)
}
func SerializeDeviceAppManagementTaskCategory(values []DeviceAppManagementTaskCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
