package models
// The VPN action to take for a specific service.
type VpnServiceExceptionAction int

const (
    // Make all traffic from that service go through the VPN
    FORCETRAFFICVIAVPN_VPNSERVICEEXCEPTIONACTION VpnServiceExceptionAction = iota
    // Allow the service outside of the VPN
    ALLOWTRAFFICOUTSIDE_VPNSERVICEEXCEPTIONACTION
    // Drop all traffic from the service
    DROPTRAFFIC_VPNSERVICEEXCEPTIONACTION
)

func (i VpnServiceExceptionAction) String() string {
    return []string{"forceTrafficViaVPN", "allowTrafficOutside", "dropTraffic"}[i]
}
func ParseVpnServiceExceptionAction(v string) (any, error) {
    result := FORCETRAFFICVIAVPN_VPNSERVICEEXCEPTIONACTION
    switch v {
        case "forceTrafficViaVPN":
            result = FORCETRAFFICVIAVPN_VPNSERVICEEXCEPTIONACTION
        case "allowTrafficOutside":
            result = ALLOWTRAFFICOUTSIDE_VPNSERVICEEXCEPTIONACTION
        case "dropTraffic":
            result = DROPTRAFFIC_VPNSERVICEEXCEPTIONACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnServiceExceptionAction(values []VpnServiceExceptionAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnServiceExceptionAction) isMultiValue() bool {
    return false
}
