package models
import (
    "math"
    "strings"
)
// Flags representing which network profile types apply to a firewall rule.
type WindowsFirewallRuleNetworkProfileTypes int

const (
    // No flags set.
    NOTCONFIGURED_WINDOWSFIREWALLRULENETWORKPROFILETYPES = 1
    // The profile for networks that are connected to domains.
    DOMAIN_WINDOWSFIREWALLRULENETWORKPROFILETYPES = 2
    // The profile for private networks.
    PRIVATE_WINDOWSFIREWALLRULENETWORKPROFILETYPES = 4
    // The profile for public networks.
    PUBLIC_WINDOWSFIREWALLRULENETWORKPROFILETYPES = 8
)

func (i WindowsFirewallRuleNetworkProfileTypes) String() string {
    var values []string
    options := []string{"notConfigured", "domain", "private", "public"}
    for p := 0; p < 4; p++ {
        mantis := WindowsFirewallRuleNetworkProfileTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsFirewallRuleNetworkProfileTypes(v string) (any, error) {
    var result WindowsFirewallRuleNetworkProfileTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notConfigured":
                result |= NOTCONFIGURED_WINDOWSFIREWALLRULENETWORKPROFILETYPES
            case "domain":
                result |= DOMAIN_WINDOWSFIREWALLRULENETWORKPROFILETYPES
            case "private":
                result |= PRIVATE_WINDOWSFIREWALLRULENETWORKPROFILETYPES
            case "public":
                result |= PUBLIC_WINDOWSFIREWALLRULENETWORKPROFILETYPES
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeWindowsFirewallRuleNetworkProfileTypes(values []WindowsFirewallRuleNetworkProfileTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsFirewallRuleNetworkProfileTypes) isMultiValue() bool {
    return true
}
