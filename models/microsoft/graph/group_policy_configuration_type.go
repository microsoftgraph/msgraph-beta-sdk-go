package graph
import (
    "strings"
    "errors"
)
type GroupPolicyConfigurationType int

const (
    POLICY_GROUPPOLICYCONFIGURATIONTYPE GroupPolicyConfigurationType = iota
    PREFERENCE_GROUPPOLICYCONFIGURATIONTYPE
)

func (i GroupPolicyConfigurationType) String() string {
    return []string{"POLICY", "PREFERENCE"}[i]
}
func ParseGroupPolicyConfigurationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "POLICY":
            return POLICY_GROUPPOLICYCONFIGURATIONTYPE, nil
        case "PREFERENCE":
            return PREFERENCE_GROUPPOLICYCONFIGURATIONTYPE, nil
    }
    return 0, errors.New("Unknown GroupPolicyConfigurationType value: " + v)
}
func SerializeGroupPolicyConfigurationType(values []GroupPolicyConfigurationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
