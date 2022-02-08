package graph
import (
    "strings"
    "errors"
)
// 
type AuthenticationMethodSignInState int

const (
    NOTSUPPORTED_AUTHENTICATIONMETHODSIGNINSTATE AuthenticationMethodSignInState = iota
    NOTALLOWEDBYPOLICY_AUTHENTICATIONMETHODSIGNINSTATE
    NOTENABLED_AUTHENTICATIONMETHODSIGNINSTATE
    PHONENUMBERNOTUNIQUE_AUTHENTICATIONMETHODSIGNINSTATE
    READY_AUTHENTICATIONMETHODSIGNINSTATE
    NOTCONFIGURED_AUTHENTICATIONMETHODSIGNINSTATE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODSIGNINSTATE
)

func (i AuthenticationMethodSignInState) String() string {
    return []string{"NOTSUPPORTED", "NOTALLOWEDBYPOLICY", "NOTENABLED", "PHONENUMBERNOTUNIQUE", "READY", "NOTCONFIGURED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAuthenticationMethodSignInState(v string) (interface{}, error) {
    result := NOTSUPPORTED_AUTHENTICATIONMETHODSIGNINSTATE
    switch strings.ToUpper(v) {
        case "NOTSUPPORTED":
            result = NOTSUPPORTED_AUTHENTICATIONMETHODSIGNINSTATE
        case "NOTALLOWEDBYPOLICY":
            result = NOTALLOWEDBYPOLICY_AUTHENTICATIONMETHODSIGNINSTATE
        case "NOTENABLED":
            result = NOTENABLED_AUTHENTICATIONMETHODSIGNINSTATE
        case "PHONENUMBERNOTUNIQUE":
            result = PHONENUMBERNOTUNIQUE_AUTHENTICATIONMETHODSIGNINSTATE
        case "READY":
            result = READY_AUTHENTICATIONMETHODSIGNINSTATE
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_AUTHENTICATIONMETHODSIGNINSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODSIGNINSTATE
        default:
            return 0, errors.New("Unknown AuthenticationMethodSignInState value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodSignInState(values []AuthenticationMethodSignInState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
