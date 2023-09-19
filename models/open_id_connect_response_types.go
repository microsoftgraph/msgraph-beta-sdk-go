package models
import (
    "errors"
    "strings"
)
// 
type OpenIdConnectResponseTypes int

const (
    CODE_OPENIDCONNECTRESPONSETYPES OpenIdConnectResponseTypes = iota
    ID_TOKEN_OPENIDCONNECTRESPONSETYPES
    TOKEN_OPENIDCONNECTRESPONSETYPES
)

func (i OpenIdConnectResponseTypes) String() string {
    var values []string
    for p := OpenIdConnectResponseTypes(1); p <= TOKEN_OPENIDCONNECTRESPONSETYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"code", "id_token", "token"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseOpenIdConnectResponseTypes(v string) (any, error) {
    var result OpenIdConnectResponseTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "code":
                result |= CODE_OPENIDCONNECTRESPONSETYPES
            case "id_token":
                result |= ID_TOKEN_OPENIDCONNECTRESPONSETYPES
            case "token":
                result |= TOKEN_OPENIDCONNECTRESPONSETYPES
            default:
                return 0, errors.New("Unknown OpenIdConnectResponseTypes value: " + v)
        }
    }
    return &result, nil
}
func SerializeOpenIdConnectResponseTypes(values []OpenIdConnectResponseTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OpenIdConnectResponseTypes) isMultiValue() bool {
    return true
}
