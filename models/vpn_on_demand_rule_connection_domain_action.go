package models
// VPN On-Demand Rule Connection Domain Action.
type VpnOnDemandRuleConnectionDomainAction int

const (
    // Connect if needed.
    CONNECTIFNEEDED_VPNONDEMANDRULECONNECTIONDOMAINACTION VpnOnDemandRuleConnectionDomainAction = iota
    // Never connect.
    NEVERCONNECT_VPNONDEMANDRULECONNECTIONDOMAINACTION
)

func (i VpnOnDemandRuleConnectionDomainAction) String() string {
    return []string{"connectIfNeeded", "neverConnect"}[i]
}
func ParseVpnOnDemandRuleConnectionDomainAction(v string) (any, error) {
    result := CONNECTIFNEEDED_VPNONDEMANDRULECONNECTIONDOMAINACTION
    switch v {
        case "connectIfNeeded":
            result = CONNECTIFNEEDED_VPNONDEMANDRULECONNECTIONDOMAINACTION
        case "neverConnect":
            result = NEVERCONNECT_VPNONDEMANDRULECONNECTIONDOMAINACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnOnDemandRuleConnectionDomainAction(values []VpnOnDemandRuleConnectionDomainAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnOnDemandRuleConnectionDomainAction) isMultiValue() bool {
    return false
}
