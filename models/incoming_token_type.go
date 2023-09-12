package models
import (
    "errors"
    "strings"
)
// 
type IncomingTokenType int

const (
    NONE_INCOMINGTOKENTYPE IncomingTokenType = iota
    PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE
    SAML11_INCOMINGTOKENTYPE
    SAML20_INCOMINGTOKENTYPE
    UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE
    REMOTEDESKTOPTOKEN_INCOMINGTOKENTYPE
)

func (i IncomingTokenType) String() string {
    var values []string
    for p := IncomingTokenType(1); p <= REMOTEDESKTOPTOKEN_INCOMINGTOKENTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "primaryRefreshToken", "saml11", "saml20", "unknownFutureValue", "remoteDesktopToken"}[p])
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
