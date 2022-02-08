package graph
import (
    "strings"
    "errors"
)
// 
type DevicePlatformType int

const (
    ANDROID_DEVICEPLATFORMTYPE DevicePlatformType = iota
    ANDROIDFORWORK_DEVICEPLATFORMTYPE
    IOS_DEVICEPLATFORMTYPE
    MACOS_DEVICEPLATFORMTYPE
    WINDOWSPHONE81_DEVICEPLATFORMTYPE
    WINDOWS81ANDLATER_DEVICEPLATFORMTYPE
    WINDOWS10ANDLATER_DEVICEPLATFORMTYPE
    ANDROIDWORKPROFILE_DEVICEPLATFORMTYPE
    UNKNOWN_DEVICEPLATFORMTYPE
    ANDROIDAOSP_DEVICEPLATFORMTYPE
)

func (i DevicePlatformType) String() string {
    return []string{"ANDROID", "ANDROIDFORWORK", "IOS", "MACOS", "WINDOWSPHONE81", "WINDOWS81ANDLATER", "WINDOWS10ANDLATER", "ANDROIDWORKPROFILE", "UNKNOWN", "ANDROIDAOSP"}[i]
}
func ParseDevicePlatformType(v string) (interface{}, error) {
    result := ANDROID_DEVICEPLATFORMTYPE
    switch strings.ToUpper(v) {
        case "ANDROID":
            result = ANDROID_DEVICEPLATFORMTYPE
        case "ANDROIDFORWORK":
            result = ANDROIDFORWORK_DEVICEPLATFORMTYPE
        case "IOS":
            result = IOS_DEVICEPLATFORMTYPE
        case "MACOS":
            result = MACOS_DEVICEPLATFORMTYPE
        case "WINDOWSPHONE81":
            result = WINDOWSPHONE81_DEVICEPLATFORMTYPE
        case "WINDOWS81ANDLATER":
            result = WINDOWS81ANDLATER_DEVICEPLATFORMTYPE
        case "WINDOWS10ANDLATER":
            result = WINDOWS10ANDLATER_DEVICEPLATFORMTYPE
        case "ANDROIDWORKPROFILE":
            result = ANDROIDWORKPROFILE_DEVICEPLATFORMTYPE
        case "UNKNOWN":
            result = UNKNOWN_DEVICEPLATFORMTYPE
        case "ANDROIDAOSP":
            result = ANDROIDAOSP_DEVICEPLATFORMTYPE
        default:
            return 0, errors.New("Unknown DevicePlatformType value: " + v)
    }
    return &result, nil
}
func SerializeDevicePlatformType(values []DevicePlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
