package graph
import (
    "strings"
    "errors"
)
// 
type SignInUserType int

const (
    MEMBER_SIGNINUSERTYPE SignInUserType = iota
    GUEST_SIGNINUSERTYPE
    UNKNOWNFUTUREVALUE_SIGNINUSERTYPE
)

func (i SignInUserType) String() string {
    return []string{"MEMBER", "GUEST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSignInUserType(v string) (interface{}, error) {
    result := MEMBER_SIGNINUSERTYPE
    switch strings.ToUpper(v) {
        case "MEMBER":
            result = MEMBER_SIGNINUSERTYPE
        case "GUEST":
            result = GUEST_SIGNINUSERTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIGNINUSERTYPE
        default:
            return 0, errors.New("Unknown SignInUserType value: " + v)
    }
    return &result, nil
}
func SerializeSignInUserType(values []SignInUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
