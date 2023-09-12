package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := AllowedRolePrincipalTypes(1); p <= UNKNOWNFUTUREVALUE_ALLOWEDROLEPRINCIPALTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"user", "servicePrincipal", "group", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAllowedRolePrincipalTypes(v string) (any, error) {
    var result AllowedRolePrincipalTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "user":
                result |= USER_ALLOWEDROLEPRINCIPALTYPES
            case "servicePrincipal":
                result |= SERVICEPRINCIPAL_ALLOWEDROLEPRINCIPALTYPES
            case "group":
                result |= GROUP_ALLOWEDROLEPRINCIPALTYPES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ALLOWEDROLEPRINCIPALTYPES
            default:
                return 0, errors.New("Unknown AllowedRolePrincipalTypes value: " + v)
        }
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
func (i AllowedRolePrincipalTypes) isMultiValue() bool {
    return true
}
