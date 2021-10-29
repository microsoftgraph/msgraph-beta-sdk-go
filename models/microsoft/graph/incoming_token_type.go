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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_INCOMINGTOKENTYPE, nil
        case "PRIMARYREFRESHTOKEN":
            return PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE, nil
        case "SAML11":
            return SAML11_INCOMINGTOKENTYPE, nil
        case "SAML20":
            return SAML20_INCOMINGTOKENTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE, nil
    }
    return 0, errors.New("Unknown IncomingTokenType value: " + v)
}
func SerializeIncomingTokenType(values []IncomingTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
