package models
import (
    "errors"
)
// Android Device Owner Kiosk Mode managed home screen orientation.
type AndroidDeviceOwnerKioskModeScreenOrientation int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION AndroidDeviceOwnerKioskModeScreenOrientation = iota
    // Portrait orientation.
    PORTRAIT_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
    // Landscape orientation.
    LANDSCAPE_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
    // Auto rotate between portrait and landscape orientations.
    AUTOROTATE_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
)

func (i AndroidDeviceOwnerKioskModeScreenOrientation) String() string {
    return []string{"notConfigured", "portrait", "landscape", "autoRotate"}[i]
}
func ParseAndroidDeviceOwnerKioskModeScreenOrientation(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
        case "portrait":
            result = PORTRAIT_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
        case "landscape":
            result = LANDSCAPE_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
        case "autoRotate":
            result = AUTOROTATE_ANDROIDDEVICEOWNERKIOSKMODESCREENORIENTATION
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerKioskModeScreenOrientation value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerKioskModeScreenOrientation(values []AndroidDeviceOwnerKioskModeScreenOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerKioskModeScreenOrientation) isMultiValue() bool {
    return false
}
