package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ALL":
            return ALL_INCLUDEDUSERTYPES, nil
        case "MEMBER":
            return MEMBER_INCLUDEDUSERTYPES, nil
        case "GUEST":
            return GUEST_INCLUDEDUSERTYPES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_INCLUDEDUSERTYPES, nil
    }
    return 0, errors.New("Unknown IncludedUserTypes value: " + v)
}
func SerializeIncludedUserTypes(values []IncludedUserTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
