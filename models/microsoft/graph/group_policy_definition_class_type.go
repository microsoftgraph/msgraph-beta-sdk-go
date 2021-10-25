package graph
import (
    "strings"
    "errors"
)
type GroupPolicyDefinitionClassType int

const (
    USER_GROUPPOLICYDEFINITIONCLASSTYPE GroupPolicyDefinitionClassType = iota
    MACHINE_GROUPPOLICYDEFINITIONCLASSTYPE
)

func (i GroupPolicyDefinitionClassType) String() string {
    return []string{"USER", "MACHINE"}[i]
}
func ParseGroupPolicyDefinitionClassType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "USER":
            return USER_GROUPPOLICYDEFINITIONCLASSTYPE, nil
        case "MACHINE":
            return MACHINE_GROUPPOLICYDEFINITIONCLASSTYPE, nil
    }
    return 0, errors.New("Unknown GroupPolicyDefinitionClassType value: " + v)
}
func SerializeGroupPolicyDefinitionClassType(values []GroupPolicyDefinitionClassType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
