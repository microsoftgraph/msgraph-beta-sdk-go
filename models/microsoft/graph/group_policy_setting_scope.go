package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicySettingScope int

const (
    UNKNOWN_GROUPPOLICYSETTINGSCOPE GroupPolicySettingScope = iota
    DEVICE_GROUPPOLICYSETTINGSCOPE
    USER_GROUPPOLICYSETTINGSCOPE
)

func (i GroupPolicySettingScope) String() string {
    return []string{"UNKNOWN", "DEVICE", "USER"}[i]
}
func ParseGroupPolicySettingScope(v string) (interface{}, error) {
    result := UNKNOWN_GROUPPOLICYSETTINGSCOPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_GROUPPOLICYSETTINGSCOPE
        case "DEVICE":
            result = DEVICE_GROUPPOLICYSETTINGSCOPE
        case "USER":
            result = USER_GROUPPOLICYSETTINGSCOPE
        default:
            return 0, errors.New("Unknown GroupPolicySettingScope value: " + v)
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
