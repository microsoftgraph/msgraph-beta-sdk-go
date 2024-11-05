package models
// The possible security types for Android Wi-Fi profiles. Default value 'Open', indicates no authentication required for the network. The security protocols supported are WEP, WPA and WPA2. 'WpaEnterprise' and 'Wpa2Enterprise' options are available for Enterprise Wi-Fi profiles. 'Wep' and 'WpaPersonal' (supports WPA and WPA2) options are available for Basic Wi-Fi profiles.
type AndroidWiFiSecurityType int

const (
    // Default. Indicates Android Wifi Security Type is set to "Open" i.e. no authentication is required. (No Authentication).
    OPEN_ANDROIDWIFISECURITYTYPE AndroidWiFiSecurityType = iota
    // Indicates Android Wifi Security Type is set to WPA encryption. Must use AndroidWorkProfileEnterpriseWifiConfiguration type to configure enterprise options.
    WPAENTERPRISE_ANDROIDWIFISECURITYTYPE
    // Indicates Android Wifi Security Type is set to WPA2 encryption. Must use AndroidWorkProfileEnterpriseWifiConfiguration type to configure enterprise options.
    WPA2ENTERPRISE_ANDROIDWIFISECURITYTYPE
    // Indicates Android Wifi Security Type is set to WEP encryption. This restricts the preSharedKey to a valid passphrase (5 or 13 characters) or a valid HEX key (10 or 26 hexidecimal characters). Use AndroidWorkProfileWifiConfiguration to configure basic Wi-Fi options.
    WEP_ANDROIDWIFISECURITYTYPE
    //  Indicates Android Wifi Security Type is set to WPA encryption. This restricts the preSharedKey to a string between 8 and 64 characters long. Use AndroidWorkProfileWifiConfiguration to configure basic Wi-Fi options.
    WPAPERSONAL_ANDROIDWIFISECURITYTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ANDROIDWIFISECURITYTYPE
)

func (i AndroidWiFiSecurityType) String() string {
    return []string{"open", "wpaEnterprise", "wpa2Enterprise", "wep", "wpaPersonal", "unknownFutureValue"}[i]
}
func ParseAndroidWiFiSecurityType(v string) (any, error) {
    result := OPEN_ANDROIDWIFISECURITYTYPE
    switch v {
        case "open":
            result = OPEN_ANDROIDWIFISECURITYTYPE
        case "wpaEnterprise":
            result = WPAENTERPRISE_ANDROIDWIFISECURITYTYPE
        case "wpa2Enterprise":
            result = WPA2ENTERPRISE_ANDROIDWIFISECURITYTYPE
        case "wep":
            result = WEP_ANDROIDWIFISECURITYTYPE
        case "wpaPersonal":
            result = WPAPERSONAL_ANDROIDWIFISECURITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDWIFISECURITYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidWiFiSecurityType(values []AndroidWiFiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidWiFiSecurityType) isMultiValue() bool {
    return false
}
