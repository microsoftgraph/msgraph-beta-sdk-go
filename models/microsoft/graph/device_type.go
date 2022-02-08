package graph
import (
    "strings"
    "errors"
)
// 
type DeviceType int

const (
    DESKTOP_DEVICETYPE DeviceType = iota
    WINDOWSRT_DEVICETYPE
    WINMO6_DEVICETYPE
    NOKIA_DEVICETYPE
    WINDOWSPHONE_DEVICETYPE
    MAC_DEVICETYPE
    WINCE_DEVICETYPE
    WINEMBEDDED_DEVICETYPE
    IPHONE_DEVICETYPE
    IPAD_DEVICETYPE
    IPOD_DEVICETYPE
    ANDROID_DEVICETYPE
    ISOCCONSUMER_DEVICETYPE
    UNIX_DEVICETYPE
    MACMDM_DEVICETYPE
    HOLOLENS_DEVICETYPE
    SURFACEHUB_DEVICETYPE
    ANDROIDFORWORK_DEVICETYPE
    ANDROIDENTERPRISE_DEVICETYPE
    WINDOWS10X_DEVICETYPE
    ANDROIDNGMS_DEVICETYPE
    CHROMEOS_DEVICETYPE
    LINUX_DEVICETYPE
    BLACKBERRY_DEVICETYPE
    PALM_DEVICETYPE
    UNKNOWN_DEVICETYPE
    CLOUDPC_DEVICETYPE
)

func (i DeviceType) String() string {
    return []string{"DESKTOP", "WINDOWSRT", "WINMO6", "NOKIA", "WINDOWSPHONE", "MAC", "WINCE", "WINEMBEDDED", "IPHONE", "IPAD", "IPOD", "ANDROID", "ISOCCONSUMER", "UNIX", "MACMDM", "HOLOLENS", "SURFACEHUB", "ANDROIDFORWORK", "ANDROIDENTERPRISE", "WINDOWS10X", "ANDROIDNGMS", "CHROMEOS", "LINUX", "BLACKBERRY", "PALM", "UNKNOWN", "CLOUDPC"}[i]
}
func ParseDeviceType(v string) (interface{}, error) {
    result := DESKTOP_DEVICETYPE
    switch strings.ToUpper(v) {
        case "DESKTOP":
            result = DESKTOP_DEVICETYPE
        case "WINDOWSRT":
            result = WINDOWSRT_DEVICETYPE
        case "WINMO6":
            result = WINMO6_DEVICETYPE
        case "NOKIA":
            result = NOKIA_DEVICETYPE
        case "WINDOWSPHONE":
            result = WINDOWSPHONE_DEVICETYPE
        case "MAC":
            result = MAC_DEVICETYPE
        case "WINCE":
            result = WINCE_DEVICETYPE
        case "WINEMBEDDED":
            result = WINEMBEDDED_DEVICETYPE
        case "IPHONE":
            result = IPHONE_DEVICETYPE
        case "IPAD":
            result = IPAD_DEVICETYPE
        case "IPOD":
            result = IPOD_DEVICETYPE
        case "ANDROID":
            result = ANDROID_DEVICETYPE
        case "ISOCCONSUMER":
            result = ISOCCONSUMER_DEVICETYPE
        case "UNIX":
            result = UNIX_DEVICETYPE
        case "MACMDM":
            result = MACMDM_DEVICETYPE
        case "HOLOLENS":
            result = HOLOLENS_DEVICETYPE
        case "SURFACEHUB":
            result = SURFACEHUB_DEVICETYPE
        case "ANDROIDFORWORK":
            result = ANDROIDFORWORK_DEVICETYPE
        case "ANDROIDENTERPRISE":
            result = ANDROIDENTERPRISE_DEVICETYPE
        case "WINDOWS10X":
            result = WINDOWS10X_DEVICETYPE
        case "ANDROIDNGMS":
            result = ANDROIDNGMS_DEVICETYPE
        case "CHROMEOS":
            result = CHROMEOS_DEVICETYPE
        case "LINUX":
            result = LINUX_DEVICETYPE
        case "BLACKBERRY":
            result = BLACKBERRY_DEVICETYPE
        case "PALM":
            result = PALM_DEVICETYPE
        case "UNKNOWN":
            result = UNKNOWN_DEVICETYPE
        case "CLOUDPC":
            result = CLOUDPC_DEVICETYPE
        default:
            return 0, errors.New("Unknown DeviceType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceType(values []DeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
