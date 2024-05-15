package models
// This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
type AospWifiSecurityType int

const (
    // No security type.
    NONE_AOSPWIFISECURITYTYPE AospWifiSecurityType = iota
    // WPA-Pre-shared-key
    WPA_AOSPWIFISECURITYTYPE
    // WEP-Pre-shared-key
    WEP_AOSPWIFISECURITYTYPE
)

func (i AospWifiSecurityType) String() string {
    return []string{"none", "wpa", "wep"}[i]
}
func ParseAospWifiSecurityType(v string) (any, error) {
    result := NONE_AOSPWIFISECURITYTYPE
    switch v {
        case "none":
            result = NONE_AOSPWIFISECURITYTYPE
        case "wpa":
            result = WPA_AOSPWIFISECURITYTYPE
        case "wep":
            result = WEP_AOSPWIFISECURITYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAospWifiSecurityType(values []AospWifiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AospWifiSecurityType) isMultiValue() bool {
    return false
}
