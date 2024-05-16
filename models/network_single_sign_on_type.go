package models
// Wi-Fi Network Single Sign On Type Settings.
type NetworkSingleSignOnType int

const (
    // Disabled
    DISABLED_NETWORKSINGLESIGNONTYPE NetworkSingleSignOnType = iota
    // Pre-Logon
    PRELOGON_NETWORKSINGLESIGNONTYPE
    // Post-Logon
    POSTLOGON_NETWORKSINGLESIGNONTYPE
)

func (i NetworkSingleSignOnType) String() string {
    return []string{"disabled", "prelogon", "postlogon"}[i]
}
func ParseNetworkSingleSignOnType(v string) (any, error) {
    result := DISABLED_NETWORKSINGLESIGNONTYPE
    switch v {
        case "disabled":
            result = DISABLED_NETWORKSINGLESIGNONTYPE
        case "prelogon":
            result = PRELOGON_NETWORKSINGLESIGNONTYPE
        case "postlogon":
            result = POSTLOGON_NETWORKSINGLESIGNONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworkSingleSignOnType(values []NetworkSingleSignOnType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworkSingleSignOnType) isMultiValue() bool {
    return false
}
