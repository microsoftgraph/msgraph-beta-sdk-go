package models
// Scope of the group policy setting.
type GroupPolicySettingScope int

const (
    // Device scope unknown
    UNKNOWN_GROUPPOLICYSETTINGSCOPE GroupPolicySettingScope = iota
    // Device scope
    DEVICE_GROUPPOLICYSETTINGSCOPE
    // User scope
    USER_GROUPPOLICYSETTINGSCOPE
)

func (i GroupPolicySettingScope) String() string {
    return []string{"unknown", "device", "user"}[i]
}
func ParseGroupPolicySettingScope(v string) (any, error) {
    result := UNKNOWN_GROUPPOLICYSETTINGSCOPE
    switch v {
        case "unknown":
            result = UNKNOWN_GROUPPOLICYSETTINGSCOPE
        case "device":
            result = DEVICE_GROUPPOLICYSETTINGSCOPE
        case "user":
            result = USER_GROUPPOLICYSETTINGSCOPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGroupPolicySettingScope(values []GroupPolicySettingScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupPolicySettingScope) isMultiValue() bool {
    return false
}
