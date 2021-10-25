package graph
import (
    "strings"
    "errors"
)
type AuthenticationPhoneType int

const (
    MOBILE_AUTHENTICATIONPHONETYPE AuthenticationPhoneType = iota
    ALTERNATEMOBILE_AUTHENTICATIONPHONETYPE
    OFFICE_AUTHENTICATIONPHONETYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONPHONETYPE
)

func (i AuthenticationPhoneType) String() string {
    return []string{"MOBILE", "ALTERNATEMOBILE", "OFFICE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAuthenticationPhoneType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "MOBILE":
            return MOBILE_AUTHENTICATIONPHONETYPE, nil
        case "ALTERNATEMOBILE":
            return ALTERNATEMOBILE_AUTHENTICATIONPHONETYPE, nil
        case "OFFICE":
            return OFFICE_AUTHENTICATIONPHONETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_AUTHENTICATIONPHONETYPE, nil
    }
    return 0, errors.New("Unknown AuthenticationPhoneType value: " + v)
}
func SerializeAuthenticationPhoneType(values []AuthenticationPhoneType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
