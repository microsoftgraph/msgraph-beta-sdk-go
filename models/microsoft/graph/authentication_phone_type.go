package graph
import (
    "strings"
    "errors"
)
// 
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
    result := MOBILE_AUTHENTICATIONPHONETYPE
    switch strings.ToUpper(v) {
        case "MOBILE":
            result = MOBILE_AUTHENTICATIONPHONETYPE
        case "ALTERNATEMOBILE":
            result = ALTERNATEMOBILE_AUTHENTICATIONPHONETYPE
        case "OFFICE":
            result = OFFICE_AUTHENTICATIONPHONETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONPHONETYPE
        default:
            return 0, errors.New("Unknown AuthenticationPhoneType value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationPhoneType(values []AuthenticationPhoneType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
