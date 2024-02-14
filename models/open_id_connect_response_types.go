package models
import (
    "errors"
    "math"
    "strings"
)
type OpenIdConnectResponseTypes int

const (
    CODE_OPENIDCONNECTRESPONSETYPES = 1
    ID_TOKEN_OPENIDCONNECTRESPONSETYPES = 2
    TOKEN_OPENIDCONNECTRESPONSETYPES = 4
)

func (i OpenIdConnectResponseTypes) String() string {
    var values []string
    options := []string{"code", "id_token", "token"}
    for p := 0; p < 3; p++ {
        mantis := OpenIdConnectResponseTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
