package models
import (
    "errors"
)
// The type of tunnels that will be present to the VPN client for configuration
type VpnTunnelConfigurationType int

const (
    // WiFi and Cellular Tunnels
    WIFIANDCELLULAR_VPNTUNNELCONFIGURATIONTYPE VpnTunnelConfigurationType = iota
    // Cellular Tunnel Only
    CELLULAR_VPNTUNNELCONFIGURATIONTYPE
    // WiFi Tunnel Only
    WIFI_VPNTUNNELCONFIGURATIONTYPE
)

func (i VpnTunnelConfigurationType) String() string {
    return []string{"wifiAndCellular", "cellular", "wifi"}[i]
}
func ParseVpnTunnelConfigurationType(v string) (any, error) {
    result := WIFIANDCELLULAR_VPNTUNNELCONFIGURATIONTYPE
    switch v {
        case "wifiAndCellular":
            result = WIFIANDCELLULAR_VPNTUNNELCONFIGURATIONTYPE
        case "cellular":
            result = CELLULAR_VPNTUNNELCONFIGURATIONTYPE
        case "wifi":
            result = WIFI_VPNTUNNELCONFIGURATIONTYPE
        default:
            return 0, errors.New("Unknown VpnTunnelConfigurationType value: " + v)
    }
    return &result, nil
}
func SerializeVpnTunnelConfigurationType(values []VpnTunnelConfigurationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
