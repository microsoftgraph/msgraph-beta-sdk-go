package models
import (
    "errors"
    "strings"
)
// Flags representing which network profile types apply to a firewall rule.
type WindowsFirewallRuleNetworkProfileTypes int

const (
    // No flags set.
    NOTCONFIGURED_WINDOWSFIREWALLRULENETWORKPROFILETYPES WindowsFirewallRuleNetworkProfileTypes = iota
    // The profile for networks that are connected to domains.
    DOMAIN_WINDOWSFIREWALLRULENETWORKPROFILETYPES
    // The profile for private networks.
    PRIVATE_WINDOWSFIREWALLRULENETWORKPROFILETYPES
    // The profile for public networks.
    PUBLIC_WINDOWSFIREWALLRULENETWORKPROFILETYPES
)

func (i WindowsFirewallRuleNetworkProfileTypes) String() string {
    var values []string
    for p := WindowsFirewallRuleNetworkProfileTypes(1); p <= PUBLIC_WINDOWSFIREWALLRULENETWORKPROFILETYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"notConfigured", "domain", "private", "public"}[p])
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
                return 0, errors.New("Unknown WindowsFirewallRuleNetworkProfileTypes value: " + v)
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
