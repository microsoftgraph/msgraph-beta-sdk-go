package graph
import (
    "strings"
    "errors"
)
// 
type ManagedDeviceManagementFeatures int

const (
    NONE_MANAGEDDEVICEMANAGEMENTFEATURES ManagedDeviceManagementFeatures = iota
    MICROSOFTMANAGEDDESKTOP_MANAGEDDEVICEMANAGEMENTFEATURES
)

func (i ManagedDeviceManagementFeatures) String() string {
    return []string{"NONE", "MICROSOFTMANAGEDDESKTOP"}[i]
}
func ParseManagedDeviceManagementFeatures(v string) (interface{}, error) {
    result := NONE_MANAGEDDEVICEMANAGEMENTFEATURES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_MANAGEDDEVICEMANAGEMENTFEATURES
        case "MICROSOFTMANAGEDDESKTOP":
            result = MICROSOFTMANAGEDDESKTOP_MANAGEDDEVICEMANAGEMENTFEATURES
        default:
            return 0, errors.New("Unknown ManagedDeviceManagementFeatures value: " + v)
    }
    return &result, nil
}
func SerializeManagedDeviceManagementFeatures(values []ManagedDeviceManagementFeatures) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
