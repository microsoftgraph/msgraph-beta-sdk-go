package graph
import (
    "strings"
    "errors"
)
// 
type AuthenticationContextDetail int

const (
    REQUIRED_AUTHENTICATIONCONTEXTDETAIL AuthenticationContextDetail = iota
    PREVIOUSLYSATISFIED_AUTHENTICATIONCONTEXTDETAIL
    NOTAPPLICABLE_AUTHENTICATIONCONTEXTDETAIL
    UNKNOWNFUTUREVALUE_AUTHENTICATIONCONTEXTDETAIL
)

func (i AuthenticationContextDetail) String() string {
    return []string{"REQUIRED", "PREVIOUSLYSATISFIED", "NOTAPPLICABLE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAuthenticationContextDetail(v string) (interface{}, error) {
    result := REQUIRED_AUTHENTICATIONCONTEXTDETAIL
    switch strings.ToUpper(v) {
        case "REQUIRED":
            result = REQUIRED_AUTHENTICATIONCONTEXTDETAIL
        case "PREVIOUSLYSATISFIED":
            result = PREVIOUSLYSATISFIED_AUTHENTICATIONCONTEXTDETAIL
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_AUTHENTICATIONCONTEXTDETAIL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONCONTEXTDETAIL
        default:
            return 0, errors.New("Unknown AuthenticationContextDetail value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationContextDetail(values []AuthenticationContextDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
