package models
import (
    "math"
    "strings"
)
type OidcResponseType int

const (
    CODE_OIDCRESPONSETYPE = 1
    ID_TOKEN_OIDCRESPONSETYPE = 2
    TOKEN_OIDCRESPONSETYPE = 4
    UNKNOWNFUTUREVALUE_OIDCRESPONSETYPE = 8
)

func (i OidcResponseType) String() string {
    var values []string
    options := []string{"code", "id_token", "token", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := OidcResponseType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseOidcResponseType(v string) (any, error) {
    var result OidcResponseType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "code":
                result |= CODE_OIDCRESPONSETYPE
            case "id_token":
                result |= ID_TOKEN_OIDCRESPONSETYPE
            case "token":
                result |= TOKEN_OIDCRESPONSETYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_OIDCRESPONSETYPE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeOidcResponseType(values []OidcResponseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OidcResponseType) isMultiValue() bool {
    return true
}
