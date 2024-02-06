package models
import (
    "errors"
    "math"
    "strings"
)
// 
type AllowedRolePrincipalTypes int

const (
    USER_ALLOWEDROLEPRINCIPALTYPES = 1
    SERVICEPRINCIPAL_ALLOWEDROLEPRINCIPALTYPES = 2
    GROUP_ALLOWEDROLEPRINCIPALTYPES = 4
    UNKNOWNFUTUREVALUE_ALLOWEDROLEPRINCIPALTYPES = 8
)

func (i AllowedRolePrincipalTypes) String() string {
    var values []string
    options := []string{"user", "servicePrincipal", "group", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := AllowedRolePrincipalTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
