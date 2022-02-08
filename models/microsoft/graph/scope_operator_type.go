package graph
import (
    "strings"
    "errors"
)
// 
type ScopeOperatorType int

const (
    BINARY_SCOPEOPERATORTYPE ScopeOperatorType = iota
    UNARY_SCOPEOPERATORTYPE
)

func (i ScopeOperatorType) String() string {
    return []string{"BINARY", "UNARY"}[i]
}
func ParseScopeOperatorType(v string) (interface{}, error) {
    result := BINARY_SCOPEOPERATORTYPE
    switch strings.ToUpper(v) {
        case "BINARY":
            result = BINARY_SCOPEOPERATORTYPE
        case "UNARY":
            result = UNARY_SCOPEOPERATORTYPE
        default:
            return 0, errors.New("Unknown ScopeOperatorType value: " + v)
    }
    return &result, nil
}
func SerializeScopeOperatorType(values []ScopeOperatorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
