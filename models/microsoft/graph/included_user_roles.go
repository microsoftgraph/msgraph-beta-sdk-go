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
    switch strings.ToUpper(v) {
        case "ALL":
            return ALL_INCLUDEDUSERROLES, nil
        case "PRIVILEGEDADMIN":
            return PRIVILEGEDADMIN_INCLUDEDUSERROLES, nil
        case "ADMIN":
            return ADMIN_INCLUDEDUSERROLES, nil
        case "USER":
            return USER_INCLUDEDUSERROLES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_INCLUDEDUSERROLES, nil
    }
    return 0, errors.New("Unknown IncludedUserRoles value: " + v)
}
func SerializeIncludedUserRoles(values []IncludedUserRoles) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
