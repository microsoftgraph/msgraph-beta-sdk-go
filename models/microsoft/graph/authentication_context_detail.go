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
    switch strings.ToUpper(v) {
        case "REQUIRED":
            return REQUIRED_AUTHENTICATIONCONTEXTDETAIL, nil
        case "PREVIOUSLYSATISFIED":
            return PREVIOUSLYSATISFIED_AUTHENTICATIONCONTEXTDETAIL, nil
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_AUTHENTICATIONCONTEXTDETAIL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_AUTHENTICATIONCONTEXTDETAIL, nil
    }
    return 0, errors.New("Unknown AuthenticationContextDetail value: " + v)
}
func SerializeAuthenticationContextDetail(values []AuthenticationContextDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
