package models
// Specify whether the rule applies to inbound traffic or outbound traffic.
type VpnTrafficDirection int

const (
    // The rule applies to all outbound traffic.
    OUTBOUND_VPNTRAFFICDIRECTION VpnTrafficDirection = iota
    // The rule applies to all inbound traffic.
    INBOUND_VPNTRAFFICDIRECTION
    // Evolvable enum member
    UNKNOWNFUTUREVALUE_VPNTRAFFICDIRECTION
)

func (i VpnTrafficDirection) String() string {
    return []string{"outbound", "inbound", "unknownFutureValue"}[i]
}
func ParseVpnTrafficDirection(v string) (any, error) {
    result := OUTBOUND_VPNTRAFFICDIRECTION
    switch v {
        case "outbound":
            result = OUTBOUND_VPNTRAFFICDIRECTION
        case "inbound":
            result = INBOUND_VPNTRAFFICDIRECTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_VPNTRAFFICDIRECTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnTrafficDirection(values []VpnTrafficDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnTrafficDirection) isMultiValue() bool {
    return false
}
