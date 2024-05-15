package models
// Device app management task category.
type DeviceAppManagementTaskCategory int

const (
    // Unknown source.
    UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY DeviceAppManagementTaskCategory = iota
    // Windows Defender ATP Threat & Vulnerability Management.
    ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
)

func (i DeviceAppManagementTaskCategory) String() string {
    return []string{"unknown", "advancedThreatProtection"}[i]
}
func ParseDeviceAppManagementTaskCategory(v string) (any, error) {
    result := UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY
        case "advancedThreatProtection":
            result = ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
        default:
            return nil, nil
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
func (i DeviceAppManagementTaskCategory) isMultiValue() bool {
    return false
}
