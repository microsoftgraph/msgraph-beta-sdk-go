package models
// VPN On-Demand Rule Connection network interface type.
type VpnOnDemandRuleInterfaceTypeMatch int

const (
    // NotConfigured
    NOTCONFIGURED_VPNONDEMANDRULEINTERFACETYPEMATCH VpnOnDemandRuleInterfaceTypeMatch = iota
    // Ethernet.
    ETHERNET_VPNONDEMANDRULEINTERFACETYPEMATCH
    // WiFi.
    WIFI_VPNONDEMANDRULEINTERFACETYPEMATCH
    // Cellular.
    CELLULAR_VPNONDEMANDRULEINTERFACETYPEMATCH
)

func (i VpnOnDemandRuleInterfaceTypeMatch) String() string {
    return []string{"notConfigured", "ethernet", "wiFi", "cellular"}[i]
}
func ParseVpnOnDemandRuleInterfaceTypeMatch(v string) (any, error) {
    result := NOTCONFIGURED_VPNONDEMANDRULEINTERFACETYPEMATCH
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_VPNONDEMANDRULEINTERFACETYPEMATCH
        case "ethernet":
            result = ETHERNET_VPNONDEMANDRULEINTERFACETYPEMATCH
        case "wiFi":
            result = WIFI_VPNONDEMANDRULEINTERFACETYPEMATCH
        case "cellular":
            result = CELLULAR_VPNONDEMANDRULEINTERFACETYPEMATCH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnOnDemandRuleInterfaceTypeMatch(values []VpnOnDemandRuleInterfaceTypeMatch) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnOnDemandRuleInterfaceTypeMatch) isMultiValue() bool {
    return false
}
