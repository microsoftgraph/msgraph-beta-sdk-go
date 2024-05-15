package models
// Group Policy Configuration Type
type GroupPolicyConfigurationType int

const (
    // The policy type does not tattoo the value, which means the value is removed allowing the original configuration value to be used. The policy type supercedes application configuration setting so the application is always aware of the value. The policy type prevents the user from modifying the value through the application's user interface.
    POLICY_GROUPPOLICYCONFIGURATIONTYPE GroupPolicyConfigurationType = iota
    // The preference type does tattoo the value, which means the value is not removed from the registry. The preference type will overwrite the user configured-value and does not retain the previous value. The preference type does not prevent the user from modifying the value through the application's user interface.
    PREFERENCE_GROUPPOLICYCONFIGURATIONTYPE
)

func (i GroupPolicyConfigurationType) String() string {
    return []string{"policy", "preference"}[i]
}
func ParseGroupPolicyConfigurationType(v string) (any, error) {
    result := POLICY_GROUPPOLICYCONFIGURATIONTYPE
    switch v {
        case "policy":
            result = POLICY_GROUPPOLICYCONFIGURATIONTYPE
        case "preference":
            result = PREFERENCE_GROUPPOLICYCONFIGURATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGroupPolicyConfigurationType(values []GroupPolicyConfigurationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupPolicyConfigurationType) isMultiValue() bool {
    return false
}
