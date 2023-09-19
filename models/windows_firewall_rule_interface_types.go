package models
import (
    "errors"
    "strings"
)
// Flags representing firewall rule interface types.
type WindowsFirewallRuleInterfaceTypes int

const (
    // No flags set.
    NOTCONFIGURED_WINDOWSFIREWALLRULEINTERFACETYPES WindowsFirewallRuleInterfaceTypes = iota
    // The Remote Access interface type.
    REMOTEACCESS_WINDOWSFIREWALLRULEINTERFACETYPES
    // The Wireless interface type.
    WIRELESS_WINDOWSFIREWALLRULEINTERFACETYPES
    // The LAN interface type.
    LAN_WINDOWSFIREWALLRULEINTERFACETYPES
)

func (i WindowsFirewallRuleInterfaceTypes) String() string {
    var values []string
    for p := WindowsFirewallRuleInterfaceTypes(1); p <= LAN_WINDOWSFIREWALLRULEINTERFACETYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"notConfigured", "remoteAccess", "wireless", "lan"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsFirewallRuleInterfaceTypes(v string) (any, error) {
    var result WindowsFirewallRuleInterfaceTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notConfigured":
                result |= NOTCONFIGURED_WINDOWSFIREWALLRULEINTERFACETYPES
            case "remoteAccess":
                result |= REMOTEACCESS_WINDOWSFIREWALLRULEINTERFACETYPES
            case "wireless":
                result |= WIRELESS_WINDOWSFIREWALLRULEINTERFACETYPES
            case "lan":
                result |= LAN_WINDOWSFIREWALLRULEINTERFACETYPES
            default:
                return 0, errors.New("Unknown WindowsFirewallRuleInterfaceTypes value: " + v)
        }
    }
    return &result, nil
}
func SerializeWindowsFirewallRuleInterfaceTypes(values []WindowsFirewallRuleInterfaceTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsFirewallRuleInterfaceTypes) isMultiValue() bool {
    return true
}
