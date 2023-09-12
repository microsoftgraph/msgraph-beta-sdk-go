package models
import (
    "errors"
)
// Specifies the routing policy for a VPN traffic rule.
type VpnTrafficRuleRoutingPolicyType int

const (
    // No routing policy specified.
    NONE_VPNTRAFFICRULEROUTINGPOLICYTYPE VpnTrafficRuleRoutingPolicyType = iota
    // Network traffic for the specified app will be routed through the VPN.
    SPLITTUNNEL_VPNTRAFFICRULEROUTINGPOLICYTYPE
    // All network traffic will be routed through the VPN.
    FORCETUNNEL_VPNTRAFFICRULEROUTINGPOLICYTYPE
)

func (i VpnTrafficRuleRoutingPolicyType) String() string {
    return []string{"none", "splitTunnel", "forceTunnel"}[i]
}
func ParseVpnTrafficRuleRoutingPolicyType(v string) (any, error) {
    result := NONE_VPNTRAFFICRULEROUTINGPOLICYTYPE
    switch v {
        case "none":
            result = NONE_VPNTRAFFICRULEROUTINGPOLICYTYPE
        case "splitTunnel":
            result = SPLITTUNNEL_VPNTRAFFICRULEROUTINGPOLICYTYPE
        case "forceTunnel":
            result = FORCETUNNEL_VPNTRAFFICRULEROUTINGPOLICYTYPE
        default:
            return 0, errors.New("Unknown VpnTrafficRuleRoutingPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeVpnTrafficRuleRoutingPolicyType(values []VpnTrafficRuleRoutingPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnTrafficRuleRoutingPolicyType) isMultiValue() bool {
    return false
}
