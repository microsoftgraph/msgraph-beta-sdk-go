package models
import (
    "errors"
)
// 
type AllowedRolePrincipalTypes int

const (
    USER_ALLOWEDROLEPRINCIPALTYPES AllowedRolePrincipalTypes = iota
    SERVICEPRINCIPAL_ALLOWEDROLEPRINCIPALTYPES
    GROUP_ALLOWEDROLEPRINCIPALTYPES
    UNKNOWNFUTUREVALUE_ALLOWEDROLEPRINCIPALTYPES
)

func (i AllowedRolePrincipalTypes) String() string {
    return []string{"user", "servicePrincipal", "group", "unknownFutureValue"}[i]
}
func ParseAllowedRolePrincipalTypes(v string) (any, error) {
    result := USER_ALLOWEDROLEPRINCIPALTYPES
    switch v {
        case "user":
            result = USER_ALLOWEDROLEPRINCIPALTYPES
        case "servicePrincipal":
            result = SERVICEPRINCIPAL_ALLOWEDROLEPRINCIPALTYPES
        case "group":
            result = GROUP_ALLOWEDROLEPRINCIPALTYPES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALLOWEDROLEPRINCIPALTYPES
        default:
            return 0, errors.New("Unknown AllowedRolePrincipalTypes value: " + v)
    }
    return &result, nil
}
func SerializeAllowedRolePrincipalTypes(values []AllowedRolePrincipalTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
