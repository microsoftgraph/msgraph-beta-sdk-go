package models
// Possible values of the ConfigurationUsage list.
type ConfigurationUsage int

const (
    // Disallowed.
    BLOCKED_CONFIGURATIONUSAGE ConfigurationUsage = iota
    // Required.
    REQUIRED_CONFIGURATIONUSAGE
    // Optional.
    ALLOWED_CONFIGURATIONUSAGE
    // Not Configured.
    NOTCONFIGURED_CONFIGURATIONUSAGE
)

func (i ConfigurationUsage) String() string {
    return []string{"blocked", "required", "allowed", "notConfigured"}[i]
}
func ParseConfigurationUsage(v string) (any, error) {
    result := BLOCKED_CONFIGURATIONUSAGE
    switch v {
        case "blocked":
            result = BLOCKED_CONFIGURATIONUSAGE
        case "required":
            result = REQUIRED_CONFIGURATIONUSAGE
        case "allowed":
            result = ALLOWED_CONFIGURATIONUSAGE
        case "notConfigured":
            result = NOTCONFIGURED_CONFIGURATIONUSAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationUsage(values []ConfigurationUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationUsage) isMultiValue() bool {
    return false
}
