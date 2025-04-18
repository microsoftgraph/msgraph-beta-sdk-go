// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// VPN On-Demand Rule Connection Action.
type VpnOnDemandRuleConnectionAction int

const (
    // Connect.
    CONNECT_VPNONDEMANDRULECONNECTIONACTION VpnOnDemandRuleConnectionAction = iota
    // Evaluate Connection.
    EVALUATECONNECTION_VPNONDEMANDRULECONNECTIONACTION
    // Ignore.
    IGNORE_VPNONDEMANDRULECONNECTIONACTION
    // Disconnect.
    DISCONNECT_VPNONDEMANDRULECONNECTIONACTION
)

func (i VpnOnDemandRuleConnectionAction) String() string {
    return []string{"connect", "evaluateConnection", "ignore", "disconnect"}[i]
}
func ParseVpnOnDemandRuleConnectionAction(v string) (any, error) {
    result := CONNECT_VPNONDEMANDRULECONNECTIONACTION
    switch v {
        case "connect":
            result = CONNECT_VPNONDEMANDRULECONNECTIONACTION
        case "evaluateConnection":
            result = EVALUATECONNECTION_VPNONDEMANDRULECONNECTIONACTION
        case "ignore":
            result = IGNORE_VPNONDEMANDRULECONNECTIONACTION
        case "disconnect":
            result = DISCONNECT_VPNONDEMANDRULECONNECTIONACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnOnDemandRuleConnectionAction(values []VpnOnDemandRuleConnectionAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnOnDemandRuleConnectionAction) isMultiValue() bool {
    return false
}
