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
    switch strings.ToUpper(v) {
        case "MEMBER":
            return MEMBER_SIGNINUSERTYPE, nil
        case "GUEST":
            return GUEST_SIGNINUSERTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIGNINUSERTYPE, nil
    }
    return 0, errors.New("Unknown SignInUserType value: " + v)
}
func SerializeSignInUserType(values []SignInUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
