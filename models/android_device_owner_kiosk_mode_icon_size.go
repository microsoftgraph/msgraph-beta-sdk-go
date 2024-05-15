package models
// Android Device Owner Kiosk Mode managed home screen icon size.
type AndroidDeviceOwnerKioskModeIconSize int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE AndroidDeviceOwnerKioskModeIconSize = iota
    // Smallest icon size.
    SMALLEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Small icon size.
    SMALL_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Regular icon size.
    REGULAR_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Large icon size.
    LARGE_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Largest icon size.
    LARGEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
)

func (i AndroidDeviceOwnerKioskModeIconSize) String() string {
    return []string{"notConfigured", "smallest", "small", "regular", "large", "largest"}[i]
}
func ParseAndroidDeviceOwnerKioskModeIconSize(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "smallest":
            result = SMALLEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "small":
            result = SMALL_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "regular":
            result = REGULAR_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "large":
            result = LARGE_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "largest":
            result = LARGEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerKioskModeIconSize(values []AndroidDeviceOwnerKioskModeIconSize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerKioskModeIconSize) isMultiValue() bool {
    return false
}
