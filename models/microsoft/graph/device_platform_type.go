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
    switch strings.ToUpper(v) {
        case "ANDROID":
            return ANDROID_DEVICEPLATFORMTYPE, nil
        case "ANDROIDFORWORK":
            return ANDROIDFORWORK_DEVICEPLATFORMTYPE, nil
        case "IOS":
            return IOS_DEVICEPLATFORMTYPE, nil
        case "MACOS":
            return MACOS_DEVICEPLATFORMTYPE, nil
        case "WINDOWSPHONE81":
            return WINDOWSPHONE81_DEVICEPLATFORMTYPE, nil
        case "WINDOWS81ANDLATER":
            return WINDOWS81ANDLATER_DEVICEPLATFORMTYPE, nil
        case "WINDOWS10ANDLATER":
            return WINDOWS10ANDLATER_DEVICEPLATFORMTYPE, nil
        case "ANDROIDWORKPROFILE":
            return ANDROIDWORKPROFILE_DEVICEPLATFORMTYPE, nil
        case "UNKNOWN":
            return UNKNOWN_DEVICEPLATFORMTYPE, nil
        case "ANDROIDAOSP":
            return ANDROIDAOSP_DEVICEPLATFORMTYPE, nil
    }
    return 0, errors.New("Unknown DevicePlatformType value: " + v)
}
func SerializeDevicePlatformType(values []DevicePlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
