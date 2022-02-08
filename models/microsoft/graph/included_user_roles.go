package graph
import (
    "strings"
    "errors"
)
// 
type IncludedUserRoles int

const (
    ALL_INCLUDEDUSERROLES IncludedUserRoles = iota
    PRIVILEGEDADMIN_INCLUDEDUSERROLES
    ADMIN_INCLUDEDUSERROLES
    USER_INCLUDEDUSERROLES
    UNKNOWNFUTUREVALUE_INCLUDEDUSERROLES
)

func (i IncludedUserRoles) String() string {
    return []string{"ALL", "PRIVILEGEDADMIN", "ADMIN", "USER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIncludedUserRoles(v string) (interface{}, error) {
    result := ALL_INCLUDEDUSERROLES
    switch strings.ToUpper(v) {
        case "ALL":
            result = ALL_INCLUDEDUSERROLES
        case "PRIVILEGEDADMIN":
            result = PRIVILEGEDADMIN_INCLUDEDUSERROLES
        case "ADMIN":
            result = ADMIN_INCLUDEDUSERROLES
        case "USER":
            result = USER_INCLUDEDUSERROLES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_INCLUDEDUSERROLES
        default:
            return 0, errors.New("Unknown IncludedUserRoles value: " + v)
    }
    return &result, nil
}
func SerializeIncludedUserRoles(values []IncludedUserRoles) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
