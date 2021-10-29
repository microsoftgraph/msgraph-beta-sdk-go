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
    switch strings.ToUpper(v) {
        case "DESKTOP":
            return DESKTOP_DEVICETYPES, nil
        case "WINDOWSRT":
            return WINDOWSRT_DEVICETYPES, nil
        case "WINMO6":
            return WINMO6_DEVICETYPES, nil
        case "NOKIA":
            return NOKIA_DEVICETYPES, nil
        case "WINDOWSPHONE":
            return WINDOWSPHONE_DEVICETYPES, nil
        case "MAC":
            return MAC_DEVICETYPES, nil
        case "WINCE":
            return WINCE_DEVICETYPES, nil
        case "WINEMBEDDED":
            return WINEMBEDDED_DEVICETYPES, nil
        case "IPHONE":
            return IPHONE_DEVICETYPES, nil
        case "IPAD":
            return IPAD_DEVICETYPES, nil
        case "IPOD":
            return IPOD_DEVICETYPES, nil
        case "ANDROID":
            return ANDROID_DEVICETYPES, nil
        case "ISOCCONSUMER":
            return ISOCCONSUMER_DEVICETYPES, nil
        case "UNIX":
            return UNIX_DEVICETYPES, nil
        case "MACMDM":
            return MACMDM_DEVICETYPES, nil
        case "HOLOLENS":
            return HOLOLENS_DEVICETYPES, nil
        case "SURFACEHUB":
            return SURFACEHUB_DEVICETYPES, nil
        case "ANDROIDFORWORK":
            return ANDROIDFORWORK_DEVICETYPES, nil
        case "ANDROIDENTERPRISE":
            return ANDROIDENTERPRISE_DEVICETYPES, nil
        case "BLACKBERRY":
            return BLACKBERRY_DEVICETYPES, nil
        case "PALM":
            return PALM_DEVICETYPES, nil
        case "UNKNOWN":
            return UNKNOWN_DEVICETYPES, nil
    }
    return 0, errors.New("Unknown DeviceTypes value: " + v)
}
func SerializeDeviceTypes(values []DeviceTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
