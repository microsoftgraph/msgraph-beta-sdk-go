package graph
import (
    "strings"
    "errors"
)
// 
type AuthenticationProtocol int

const (
    WSFED_AUTHENTICATIONPROTOCOL AuthenticationProtocol = iota
    SAML_AUTHENTICATIONPROTOCOL
    UNKNOWNFUTUREVALUE_AUTHENTICATIONPROTOCOL
)

func (i AuthenticationProtocol) String() string {
    return []string{"WSFED", "SAML", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAuthenticationProtocol(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "WSFED":
            return WSFED_AUTHENTICATIONPROTOCOL, nil
        case "SAML":
            return SAML_AUTHENTICATIONPROTOCOL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_AUTHENTICATIONPROTOCOL, nil
    }
    return 0, errors.New("Unknown AuthenticationProtocol value: " + v)
}
func SerializeAuthenticationProtocol(values []AuthenticationProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
