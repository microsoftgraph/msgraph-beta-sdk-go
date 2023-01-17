package models
import (
    "errors"
)
// Group Policy Definition Class Type.
type GroupPolicyDefinitionClassType int

const (
    // Identifies placement of the policy setting under the user configuration node.
    USER_GROUPPOLICYDEFINITIONCLASSTYPE GroupPolicyDefinitionClassType = iota
    // Identifies placement of the policy setting under the computer configuration node.
    MACHINE_GROUPPOLICYDEFINITIONCLASSTYPE
)

func (i GroupPolicyDefinitionClassType) String() string {
    return []string{"user", "machine"}[i]
}
func ParseGroupPolicyDefinitionClassType(v string) (any, error) {
    result := USER_GROUPPOLICYDEFINITIONCLASSTYPE
    switch v {
        case "user":
            result = USER_GROUPPOLICYDEFINITIONCLASSTYPE
        case "machine":
            result = MACHINE_GROUPPOLICYDEFINITIONCLASSTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyDefinitionClassType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyDefinitionClassType(values []GroupPolicyDefinitionClassType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
