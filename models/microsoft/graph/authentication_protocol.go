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
    result := WSFED_AUTHENTICATIONPROTOCOL
    switch strings.ToUpper(v) {
        case "WSFED":
            result = WSFED_AUTHENTICATIONPROTOCOL
        case "SAML":
            result = SAML_AUTHENTICATIONPROTOCOL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONPROTOCOL
        default:
            return 0, errors.New("Unknown AuthenticationProtocol value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationProtocol(values []AuthenticationProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
