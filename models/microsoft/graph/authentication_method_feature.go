package graph
import (
    "strings"
    "errors"
)
// 
type AuthenticationMethodFeature int

const (
    SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE AuthenticationMethodFeature = iota
    SSPRENABLED_AUTHENTICATIONMETHODFEATURE
    SSPRCAPABLE_AUTHENTICATIONMETHODFEATURE
    PASSWORDLESSCAPABLE_AUTHENTICATIONMETHODFEATURE
    MFACAPABLE_AUTHENTICATIONMETHODFEATURE
)

func (i AuthenticationMethodFeature) String() string {
    return []string{"SSPRREGISTERED", "SSPRENABLED", "SSPRCAPABLE", "PASSWORDLESSCAPABLE", "MFACAPABLE"}[i]
}
func ParseAuthenticationMethodFeature(v string) (interface{}, error) {
    result := SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE
    switch strings.ToUpper(v) {
        case "SSPRREGISTERED":
            result = SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE
        case "SSPRENABLED":
            result = SSPRENABLED_AUTHENTICATIONMETHODFEATURE
        case "SSPRCAPABLE":
            result = SSPRCAPABLE_AUTHENTICATIONMETHODFEATURE
        case "PASSWORDLESSCAPABLE":
            result = PASSWORDLESSCAPABLE_AUTHENTICATIONMETHODFEATURE
        case "MFACAPABLE":
            result = MFACAPABLE_AUTHENTICATIONMETHODFEATURE
        default:
            return 0, errors.New("Unknown AuthenticationMethodFeature value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodFeature(values []AuthenticationMethodFeature) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
