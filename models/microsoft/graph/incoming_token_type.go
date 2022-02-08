package graph
import (
    "strings"
    "errors"
)
// 
type IncomingTokenType int

const (
    NONE_INCOMINGTOKENTYPE IncomingTokenType = iota
    PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE
    SAML11_INCOMINGTOKENTYPE
    SAML20_INCOMINGTOKENTYPE
    UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE
)

func (i IncomingTokenType) String() string {
    return []string{"NONE", "PRIMARYREFRESHTOKEN", "SAML11", "SAML20", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIncomingTokenType(v string) (interface{}, error) {
    result := NONE_INCOMINGTOKENTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_INCOMINGTOKENTYPE
        case "PRIMARYREFRESHTOKEN":
            result = PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE
        case "SAML11":
            result = SAML11_INCOMINGTOKENTYPE
        case "SAML20":
            result = SAML20_INCOMINGTOKENTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE
        default:
            return 0, errors.New("Unknown IncomingTokenType value: " + v)
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
