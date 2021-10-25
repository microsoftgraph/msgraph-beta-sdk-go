package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_GROUPPOLICYSETTINGSCOPE, nil
        case "DEVICE":
            return DEVICE_GROUPPOLICYSETTINGSCOPE, nil
        case "USER":
            return USER_GROUPPOLICYSETTINGSCOPE, nil
    }
    return 0, errors.New("Unknown GroupPolicySettingScope value: " + v)
}
func SerializeGroupPolicySettingScope(values []GroupPolicySettingScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
