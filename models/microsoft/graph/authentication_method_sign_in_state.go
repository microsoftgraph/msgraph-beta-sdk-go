package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTSUPPORTED":
            return NOTSUPPORTED_AUTHENTICATIONMETHODSIGNINSTATE, nil
        case "NOTALLOWEDBYPOLICY":
            return NOTALLOWEDBYPOLICY_AUTHENTICATIONMETHODSIGNINSTATE, nil
        case "NOTENABLED":
            return NOTENABLED_AUTHENTICATIONMETHODSIGNINSTATE, nil
        case "PHONENUMBERNOTUNIQUE":
            return PHONENUMBERNOTUNIQUE_AUTHENTICATIONMETHODSIGNINSTATE, nil
        case "READY":
            return READY_AUTHENTICATIONMETHODSIGNINSTATE, nil
        case "NOTCONFIGURED":
            return NOTCONFIGURED_AUTHENTICATIONMETHODSIGNINSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODSIGNINSTATE, nil
    }
    return 0, errors.New("Unknown AuthenticationMethodSignInState value: " + v)
}
func SerializeAuthenticationMethodSignInState(values []AuthenticationMethodSignInState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
