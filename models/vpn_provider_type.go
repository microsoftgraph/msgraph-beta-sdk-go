package models
// Provider type for per-app VPN.
type VpnProviderType int

const (
    // Tunnel traffic is not explicitly configured.
    NOTCONFIGURED_VPNPROVIDERTYPE VpnProviderType = iota
    // Tunnel traffic at the application layer.
    APPPROXY_VPNPROVIDERTYPE
    // Tunnel traffic at the IP layer.
    PACKETTUNNEL_VPNPROVIDERTYPE
)

func (i VpnProviderType) String() string {
    return []string{"notConfigured", "appProxy", "packetTunnel"}[i]
}
func ParseVpnProviderType(v string) (any, error) {
    result := NOTCONFIGURED_VPNPROVIDERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_VPNPROVIDERTYPE
        case "appProxy":
            result = APPPROXY_VPNPROVIDERTYPE
        case "packetTunnel":
            result = PACKETTUNNEL_VPNPROVIDERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnProviderType(values []VpnProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnProviderType) isMultiValue() bool {
    return false
}
