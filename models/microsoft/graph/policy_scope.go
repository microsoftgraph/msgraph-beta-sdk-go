package graph
import (
    "strings"
    "errors"
)
// 
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
    result := NONE_POLICYSCOPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_POLICYSCOPE
        case "ALL":
            result = ALL_POLICYSCOPE
        case "SELECTED":
            result = SELECTED_POLICYSCOPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_POLICYSCOPE
        default:
            return 0, errors.New("Unknown PolicyScope value: " + v)
    }
    return &result, nil
}
func SerializePolicyScope(values []PolicyScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
