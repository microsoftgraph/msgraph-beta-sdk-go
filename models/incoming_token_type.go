package models
import (
    "errors"
    "math"
    "strings"
)
// 
type IncomingTokenType int

const (
    NONE_INCOMINGTOKENTYPE = 1
    PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE = 2
    SAML11_INCOMINGTOKENTYPE = 4
    SAML20_INCOMINGTOKENTYPE = 8
    UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE = 16
    REMOTEDESKTOPTOKEN_INCOMINGTOKENTYPE = 32
)

func (i IncomingTokenType) String() string {
    var values []string
    options := []string{"none", "primaryRefreshToken", "saml11", "saml20", "unknownFutureValue", "remoteDesktopToken"}
    for p := 0; p < 6; p++ {
        mantis := IncomingTokenType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseIncomingTokenType(v string) (any, error) {
    var result IncomingTokenType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_INCOMINGTOKENTYPE
            case "primaryRefreshToken":
                result |= PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE
            case "saml11":
                result |= SAML11_INCOMINGTOKENTYPE
            case "saml20":
                result |= SAML20_INCOMINGTOKENTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE
            case "remoteDesktopToken":
                result |= REMOTEDESKTOPTOKEN_INCOMINGTOKENTYPE
            default:
                return 0, errors.New("Unknown IncomingTokenType value: " + v)
        }
    }
    return &result, nil
}
func SerializeIncomingTokenType(values []IncomingTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IncomingTokenType) isMultiValue() bool {
    return true
}
