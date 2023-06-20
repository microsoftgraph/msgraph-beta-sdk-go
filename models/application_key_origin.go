package models
import (
    "errors"
)
// 
type ApplicationKeyOrigin int

const (
    APPLICATION_APPLICATIONKEYORIGIN ApplicationKeyOrigin = iota
    SERVICEPRINCIPAL_APPLICATIONKEYORIGIN
    UNKNOWNFUTUREVALUE_APPLICATIONKEYORIGIN
)

func (i ApplicationKeyOrigin) String() string {
    return []string{"application", "servicePrincipal", "unknownFutureValue"}[i]
}
func ParseApplicationKeyOrigin(v string) (any, error) {
    result := APPLICATION_APPLICATIONKEYORIGIN
    switch v {
        case "application":
            result = APPLICATION_APPLICATIONKEYORIGIN
        case "servicePrincipal":
            result = SERVICEPRINCIPAL_APPLICATIONKEYORIGIN
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLICATIONKEYORIGIN
        default:
            return 0, errors.New("Unknown ApplicationKeyOrigin value: " + v)
    }
    return &result, nil
}
func SerializeApplicationKeyOrigin(values []ApplicationKeyOrigin) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
