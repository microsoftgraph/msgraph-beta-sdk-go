package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyConfigurationType int

const (
    POLICY_GROUPPOLICYCONFIGURATIONTYPE GroupPolicyConfigurationType = iota
    PREFERENCE_GROUPPOLICYCONFIGURATIONTYPE
)

func (i GroupPolicyConfigurationType) String() string {
    return []string{"POLICY", "PREFERENCE"}[i]
}
func ParseGroupPolicyConfigurationType(v string) (interface{}, error) {
    result := POLICY_GROUPPOLICYCONFIGURATIONTYPE
    switch strings.ToUpper(v) {
        case "POLICY":
            result = POLICY_GROUPPOLICYCONFIGURATIONTYPE
        case "PREFERENCE":
            result = PREFERENCE_GROUPPOLICYCONFIGURATIONTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyConfigurationType value: " + v)
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
