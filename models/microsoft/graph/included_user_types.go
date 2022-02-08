package graph
import (
    "strings"
    "errors"
)
// 
type IncludedUserTypes int

const (
    ALL_INCLUDEDUSERTYPES IncludedUserTypes = iota
    MEMBER_INCLUDEDUSERTYPES
    GUEST_INCLUDEDUSERTYPES
    UNKNOWNFUTUREVALUE_INCLUDEDUSERTYPES
)

func (i IncludedUserTypes) String() string {
    return []string{"ALL", "MEMBER", "GUEST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIncludedUserTypes(v string) (interface{}, error) {
    result := ALL_INCLUDEDUSERTYPES
    switch strings.ToUpper(v) {
        case "ALL":
            result = ALL_INCLUDEDUSERTYPES
        case "MEMBER":
            result = MEMBER_INCLUDEDUSERTYPES
        case "GUEST":
            result = GUEST_INCLUDEDUSERTYPES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_INCLUDEDUSERTYPES
        default:
            return 0, errors.New("Unknown IncludedUserTypes value: " + v)
    }
    return &result, nil
}
func SerializeIncludedUserTypes(values []IncludedUserTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
