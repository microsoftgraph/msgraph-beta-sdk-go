package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementResourceAccessProfileIntent int

const (
    APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT DeviceManagementResourceAccessProfileIntent = iota
    REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
)

func (i DeviceManagementResourceAccessProfileIntent) String() string {
    return []string{"APPLY", "REMOVE"}[i]
}
func ParseDeviceManagementResourceAccessProfileIntent(v string) (interface{}, error) {
    result := APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
    switch strings.ToUpper(v) {
        case "APPLY":
            result = APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
        case "REMOVE":
            result = REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
        default:
            return 0, errors.New("Unknown DeviceManagementResourceAccessProfileIntent value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementResourceAccessProfileIntent(values []DeviceManagementResourceAccessProfileIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
