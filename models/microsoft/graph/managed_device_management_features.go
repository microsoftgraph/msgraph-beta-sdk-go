package graph
import (
    "strings"
    "errors"
)
type ManagedDeviceManagementFeatures int

const (
    NONE_MANAGEDDEVICEMANAGEMENTFEATURES ManagedDeviceManagementFeatures = iota
    MICROSOFTMANAGEDDESKTOP_MANAGEDDEVICEMANAGEMENTFEATURES
)

func (i ManagedDeviceManagementFeatures) String() string {
    return []string{"NONE", "MICROSOFTMANAGEDDESKTOP"}[i]
}
func ParseManagedDeviceManagementFeatures(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_MANAGEDDEVICEMANAGEMENTFEATURES, nil
        case "MICROSOFTMANAGEDDESKTOP":
            return MICROSOFTMANAGEDDESKTOP_MANAGEDDEVICEMANAGEMENTFEATURES, nil
    }
    return 0, errors.New("Unknown ManagedDeviceManagementFeatures value: " + v)
}
func SerializeManagedDeviceManagementFeatures(values []ManagedDeviceManagementFeatures) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
