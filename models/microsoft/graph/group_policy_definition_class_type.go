package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyDefinitionClassType int

const (
    USER_GROUPPOLICYDEFINITIONCLASSTYPE GroupPolicyDefinitionClassType = iota
    MACHINE_GROUPPOLICYDEFINITIONCLASSTYPE
)

func (i GroupPolicyDefinitionClassType) String() string {
    return []string{"USER", "MACHINE"}[i]
}
func ParseGroupPolicyDefinitionClassType(v string) (interface{}, error) {
    result := USER_GROUPPOLICYDEFINITIONCLASSTYPE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_GROUPPOLICYDEFINITIONCLASSTYPE
        case "MACHINE":
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
