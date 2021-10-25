package graph
import (
    "strings"
    "errors"
)
type ScopeOperatorType int

const (
    BINARY_SCOPEOPERATORTYPE ScopeOperatorType = iota
    UNARY_SCOPEOPERATORTYPE
)

func (i ScopeOperatorType) String() string {
    return []string{"BINARY", "UNARY"}[i]
}
func ParseScopeOperatorType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BINARY":
            return BINARY_SCOPEOPERATORTYPE, nil
        case "UNARY":
            return UNARY_SCOPEOPERATORTYPE, nil
    }
    return 0, errors.New("Unknown ScopeOperatorType value: " + v)
}
func SerializeScopeOperatorType(values []ScopeOperatorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
