package models
// Wi-Fi Proxy Settings.
type WiFiProxySetting int

const (
    // Default. Indicates Wi-Fi Proxy is not set.
    NONE_WIFIPROXYSETTING WiFiProxySetting = iota
    // Indicates Wi-Fi Proxy is set by manually specifying an address and port as well as an optional list of hostnames that are exculded. This value is not supported for AndroidWorkProfileWiFiConfigurations.
    MANUAL_WIFIPROXYSETTING
    // Indicates Wi-Fi Proxy is set automatically by providing the URL to a PAC (Proxy Auto Configuration) file which contains a list of proxy servers to use.
    AUTOMATIC_WIFIPROXYSETTING
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WIFIPROXYSETTING
)

func (i WiFiProxySetting) String() string {
    return []string{"none", "manual", "automatic", "unknownFutureValue"}[i]
}
func ParseWiFiProxySetting(v string) (any, error) {
    result := NONE_WIFIPROXYSETTING
    switch v {
        case "none":
            result = NONE_WIFIPROXYSETTING
        case "manual":
            result = MANUAL_WIFIPROXYSETTING
        case "automatic":
            result = AUTOMATIC_WIFIPROXYSETTING
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WIFIPROXYSETTING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWiFiProxySetting(values []WiFiProxySetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WiFiProxySetting) isMultiValue() bool {
    return false
}
