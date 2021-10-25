package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SSPRREGISTERED":
            return SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE, nil
        case "SSPRENABLED":
            return SSPRENABLED_AUTHENTICATIONMETHODFEATURE, nil
        case "SSPRCAPABLE":
            return SSPRCAPABLE_AUTHENTICATIONMETHODFEATURE, nil
        case "PASSWORDLESSCAPABLE":
            return PASSWORDLESSCAPABLE_AUTHENTICATIONMETHODFEATURE, nil
        case "MFACAPABLE":
            return MFACAPABLE_AUTHENTICATIONMETHODFEATURE, nil
    }
    return 0, errors.New("Unknown AuthenticationMethodFeature value: " + v)
}
func SerializeAuthenticationMethodFeature(values []AuthenticationMethodFeature) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
