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
    result := ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
    switch strings.ToUpper(v) {
        case "ALL":
            result = ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
        case "ANY":
            result = ANY_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
        default:
            return 0, errors.New("Unknown ScopeOperatorMultiValuedComparisonType value: " + v)
    }
    return &result, nil
}
func SerializeScopeOperatorMultiValuedComparisonType(values []ScopeOperatorMultiValuedComparisonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
