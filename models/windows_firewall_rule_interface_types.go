package models
import (
    "errors"
    "math"
    "strings"
)
// Flags representing firewall rule interface types.
type WindowsFirewallRuleInterfaceTypes int

const (
    // No flags set.
    NOTCONFIGURED_WINDOWSFIREWALLRULEINTERFACETYPES = 1
    // The Remote Access interface type.
    REMOTEACCESS_WINDOWSFIREWALLRULEINTERFACETYPES = 2
    // The Wireless interface type.
    WIRELESS_WINDOWSFIREWALLRULEINTERFACETYPES = 4
    // The LAN interface type.
    LAN_WINDOWSFIREWALLRULEINTERFACETYPES = 8
)

func (i WindowsFirewallRuleInterfaceTypes) String() string {
    var values []string
    options := []string{"notConfigured", "remoteAccess", "wireless", "lan"}
    for p := 0; p < 4; p++ {
        mantis := WindowsFirewallRuleInterfaceTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
