package graph
import (
    "strings"
    "errors"
)
type PolicyScope int

const (
    NONE_POLICYSCOPE PolicyScope = iota
    ALL_POLICYSCOPE
    SELECTED_POLICYSCOPE
    UNKNOWNFUTUREVALUE_POLICYSCOPE
)

func (i PolicyScope) String() string {
    return []string{"NONE", "ALL", "SELECTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePolicyScope(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_POLICYSCOPE, nil
        case "ALL":
            return ALL_POLICYSCOPE, nil
        case "SELECTED":
            return SELECTED_POLICYSCOPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_POLICYSCOPE, nil
    }
    return 0, errors.New("Unknown PolicyScope value: " + v)
}
func SerializePolicyScope(values []PolicyScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
