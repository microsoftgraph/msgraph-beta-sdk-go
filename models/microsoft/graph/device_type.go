package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "DESKTOP":
            return DESKTOP_DEVICETYPE, nil
        case "WINDOWSRT":
            return WINDOWSRT_DEVICETYPE, nil
        case "WINMO6":
            return WINMO6_DEVICETYPE, nil
        case "NOKIA":
            return NOKIA_DEVICETYPE, nil
        case "WINDOWSPHONE":
            return WINDOWSPHONE_DEVICETYPE, nil
        case "MAC":
            return MAC_DEVICETYPE, nil
        case "WINCE":
            return WINCE_DEVICETYPE, nil
        case "WINEMBEDDED":
            return WINEMBEDDED_DEVICETYPE, nil
        case "IPHONE":
            return IPHONE_DEVICETYPE, nil
        case "IPAD":
            return IPAD_DEVICETYPE, nil
        case "IPOD":
            return IPOD_DEVICETYPE, nil
        case "ANDROID":
            return ANDROID_DEVICETYPE, nil
        case "ISOCCONSUMER":
            return ISOCCONSUMER_DEVICETYPE, nil
        case "UNIX":
            return UNIX_DEVICETYPE, nil
        case "MACMDM":
            return MACMDM_DEVICETYPE, nil
        case "HOLOLENS":
            return HOLOLENS_DEVICETYPE, nil
        case "SURFACEHUB":
            return SURFACEHUB_DEVICETYPE, nil
        case "ANDROIDFORWORK":
            return ANDROIDFORWORK_DEVICETYPE, nil
        case "ANDROIDENTERPRISE":
            return ANDROIDENTERPRISE_DEVICETYPE, nil
        case "WINDOWS10X":
            return WINDOWS10X_DEVICETYPE, nil
        case "ANDROIDNGMS":
            return ANDROIDNGMS_DEVICETYPE, nil
        case "CHROMEOS":
            return CHROMEOS_DEVICETYPE, nil
        case "LINUX":
            return LINUX_DEVICETYPE, nil
        case "BLACKBERRY":
            return BLACKBERRY_DEVICETYPE, nil
        case "PALM":
            return PALM_DEVICETYPE, nil
        case "UNKNOWN":
            return UNKNOWN_DEVICETYPE, nil
        case "CLOUDPC":
            return CLOUDPC_DEVICETYPE, nil
    }
    return 0, errors.New("Unknown DeviceType value: " + v)
}
func SerializeDeviceType(values []DeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
