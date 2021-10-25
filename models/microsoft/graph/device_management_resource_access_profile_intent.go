package graph
import (
    "strings"
    "errors"
)
type DeviceManagementResourceAccessProfileIntent int

const (
    APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT DeviceManagementResourceAccessProfileIntent = iota
    REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
)

func (i DeviceManagementResourceAccessProfileIntent) String() string {
    return []string{"APPLY", "REMOVE"}[i]
}
func ParseDeviceManagementResourceAccessProfileIntent(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "APPLY":
            return APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT, nil
        case "REMOVE":
            return REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT, nil
    }
    return 0, errors.New("Unknown DeviceManagementResourceAccessProfileIntent value: " + v)
}
func SerializeDeviceManagementResourceAccessProfileIntent(values []DeviceManagementResourceAccessProfileIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
