package graph
import (
    "strings"
    "errors"
)
// 
type ScopeOperatorMultiValuedComparisonType int

const (
    ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE ScopeOperatorMultiValuedComparisonType = iota
    ANY_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
)

func (i ScopeOperatorMultiValuedComparisonType) String() string {
    return []string{"ALL", "ANY"}[i]
}
func ParseScopeOperatorMultiValuedComparisonType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ALL":
            return ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE, nil
        case "ANY":
            return ANY_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE, nil
    }
    return 0, errors.New("Unknown ScopeOperatorMultiValuedComparisonType value: " + v)
}
func SerializeScopeOperatorMultiValuedComparisonType(values []ScopeOperatorMultiValuedComparisonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
