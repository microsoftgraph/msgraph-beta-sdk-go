package graph
import (
    "strings"
    "errors"
)
// 
type DeviceTypes int

const (
    DESKTOP_DEVICETYPES DeviceTypes = iota
    WINDOWSRT_DEVICETYPES
    WINMO6_DEVICETYPES
    NOKIA_DEVICETYPES
    WINDOWSPHONE_DEVICETYPES
    MAC_DEVICETYPES
    WINCE_DEVICETYPES
    WINEMBEDDED_DEVICETYPES
    IPHONE_DEVICETYPES
    IPAD_DEVICETYPES
    IPOD_DEVICETYPES
    ANDROID_DEVICETYPES
    ISOCCONSUMER_DEVICETYPES
    UNIX_DEVICETYPES
    MACMDM_DEVICETYPES
    HOLOLENS_DEVICETYPES
    SURFACEHUB_DEVICETYPES
    ANDROIDFORWORK_DEVICETYPES
    ANDROIDENTERPRISE_DEVICETYPES
    BLACKBERRY_DEVICETYPES
    PALM_DEVICETYPES
    UNKNOWN_DEVICETYPES
)

func (i DeviceTypes) String() string {
    return []string{"DESKTOP", "WINDOWSRT", "WINMO6", "NOKIA", "WINDOWSPHONE", "MAC", "WINCE", "WINEMBEDDED", "IPHONE", "IPAD", "IPOD", "ANDROID", "ISOCCONSUMER", "UNIX", "MACMDM", "HOLOLENS", "SURFACEHUB", "ANDROIDFORWORK", "ANDROIDENTERPRISE", "BLACKBERRY", "PALM", "UNKNOWN"}[i]
}
func ParseDeviceTypes(v string) (interface{}, error) {
    result := DESKTOP_DEVICETYPES
    switch strings.ToUpper(v) {
        case "DESKTOP":
            result = DESKTOP_DEVICETYPES
        case "WINDOWSRT":
            result = WINDOWSRT_DEVICETYPES
        case "WINMO6":
            result = WINMO6_DEVICETYPES
        case "NOKIA":
            result = NOKIA_DEVICETYPES
        case "WINDOWSPHONE":
            result = WINDOWSPHONE_DEVICETYPES
        case "MAC":
            result = MAC_DEVICETYPES
        case "WINCE":
            result = WINCE_DEVICETYPES
        case "WINEMBEDDED":
            result = WINEMBEDDED_DEVICETYPES
        case "IPHONE":
            result = IPHONE_DEVICETYPES
        case "IPAD":
            result = IPAD_DEVICETYPES
        case "IPOD":
            result = IPOD_DEVICETYPES
        case "ANDROID":
            result = ANDROID_DEVICETYPES
        case "ISOCCONSUMER":
            result = ISOCCONSUMER_DEVICETYPES
        case "UNIX":
            result = UNIX_DEVICETYPES
        case "MACMDM":
            result = MACMDM_DEVICETYPES
        case "HOLOLENS":
            result = HOLOLENS_DEVICETYPES
        case "SURFACEHUB":
            result = SURFACEHUB_DEVICETYPES
        case "ANDROIDFORWORK":
            result = ANDROIDFORWORK_DEVICETYPES
        case "ANDROIDENTERPRISE":
            result = ANDROIDENTERPRISE_DEVICETYPES
        case "BLACKBERRY":
            result = BLACKBERRY_DEVICETYPES
        case "PALM":
            result = PALM_DEVICETYPES
        case "UNKNOWN":
            result = UNKNOWN_DEVICETYPES
        default:
            return 0, errors.New("Unknown DeviceTypes value: " + v)
    }
    return &result, nil
}
func SerializeDeviceTypes(values []DeviceTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
