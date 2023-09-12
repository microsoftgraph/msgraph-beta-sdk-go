package models
import (
    "errors"
)
// Wi-Fi Proxy Settings.
type WiFiProxySetting int

const (
    // No Proxy.
    NONE_WIFIPROXYSETTING WiFiProxySetting = iota
    // Manual Proxy Settings via Address and Port.
    MANUAL_WIFIPROXYSETTING
    // Automatic Proxy Settings via URL.
    AUTOMATIC_WIFIPROXYSETTING
)

func (i WiFiProxySetting) String() string {
    return []string{"none", "manual", "automatic"}[i]
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
        default:
            return 0, errors.New("Unknown WiFiProxySetting value: " + v)
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
