package models
// Firewall rule traffic directions.
type WindowsFirewallRuleTrafficDirectionType int

const (
    // Not configured.
    NOTCONFIGURED_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE WindowsFirewallRuleTrafficDirectionType = iota
    // The rule applies to outbound traffic.
    OUT_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
    // The rule applies to inbound traffic.
    IN_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
)

func (i WindowsFirewallRuleTrafficDirectionType) String() string {
    return []string{"notConfigured", "out", "in"}[i]
}
func ParseWindowsFirewallRuleTrafficDirectionType(v string) (any, error) {
    result := NOTCONFIGURED_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
        case "out":
            result = OUT_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
        case "in":
            result = IN_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsFirewallRuleTrafficDirectionType(values []WindowsFirewallRuleTrafficDirectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsFirewallRuleTrafficDirectionType) isMultiValue() bool {
    return false
}
