package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
type SignInFrequencyAuthenticationType int

const (
    PRIMARYANDSECONDARYAUTHENTICATION_SIGNINFREQUENCYAUTHENTICATIONTYPE SignInFrequencyAuthenticationType = iota
    SECONDARYAUTHENTICATION_SIGNINFREQUENCYAUTHENTICATIONTYPE
    UNKNOWNFUTUREVALUE_SIGNINFREQUENCYAUTHENTICATIONTYPE
)

func (i SignInFrequencyAuthenticationType) String() string {
    return []string{"PRIMARYANDSECONDARYAUTHENTICATION", "SECONDARYAUTHENTICATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSignInFrequencyAuthenticationType(v string) (interface{}, error) {
    result := PRIMARYANDSECONDARYAUTHENTICATION_SIGNINFREQUENCYAUTHENTICATIONTYPE
    switch strings.ToUpper(v) {
        case "PRIMARYANDSECONDARYAUTHENTICATION":
            result = PRIMARYANDSECONDARYAUTHENTICATION_SIGNINFREQUENCYAUTHENTICATIONTYPE
        case "SECONDARYAUTHENTICATION":
            result = SECONDARYAUTHENTICATION_SIGNINFREQUENCYAUTHENTICATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIGNINFREQUENCYAUTHENTICATIONTYPE
        default:
            return 0, errors.New("Unknown SignInFrequencyAuthenticationType value: " + v)
    }
    return &result, nil
}
func SerializeSignInFrequencyAuthenticationType(values []SignInFrequencyAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
