package models
// Device management features.
type ManagedDeviceManagementFeatures int

const (
    // Unknown device management features.
    NONE_MANAGEDDEVICEMANAGEMENTFEATURES ManagedDeviceManagementFeatures = iota
    // Microsoft Managed Desktop
    MICROSOFTMANAGEDDESKTOP_MANAGEDDEVICEMANAGEMENTFEATURES
)

func (i ManagedDeviceManagementFeatures) String() string {
    return []string{"none", "microsoftManagedDesktop"}[i]
}
func ParseManagedDeviceManagementFeatures(v string) (any, error) {
    result := NONE_MANAGEDDEVICEMANAGEMENTFEATURES
    switch v {
        case "none":
            result = NONE_MANAGEDDEVICEMANAGEMENTFEATURES
        case "microsoftManagedDesktop":
            result = MICROSOFTMANAGEDDESKTOP_MANAGEDDEVICEMANAGEMENTFEATURES
        default:
            return nil, nil
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
func (i ManagedDeviceManagementFeatures) isMultiValue() bool {
    return false
}
