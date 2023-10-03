package models
import (
    "errors"
)
// 
type AuthorizationSystemActionSeverity int

const (
    NORMAL_AUTHORIZATIONSYSTEMACTIONSEVERITY AuthorizationSystemActionSeverity = iota
    HIGH_AUTHORIZATIONSYSTEMACTIONSEVERITY
    UNKNOWNFUTUREVALUE_AUTHORIZATIONSYSTEMACTIONSEVERITY
)

func (i AuthorizationSystemActionSeverity) String() string {
    return []string{"normal", "high", "unknownFutureValue"}[i]
}
func ParseAuthorizationSystemActionSeverity(v string) (any, error) {
    result := NORMAL_AUTHORIZATIONSYSTEMACTIONSEVERITY
    switch v {
        case "normal":
            result = NORMAL_AUTHORIZATIONSYSTEMACTIONSEVERITY
        case "high":
            result = HIGH_AUTHORIZATIONSYSTEMACTIONSEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHORIZATIONSYSTEMACTIONSEVERITY
        default:
            return 0, errors.New("Unknown AuthorizationSystemActionSeverity value: " + v)
    }
    return &result, nil
}
func SerializeAuthorizationSystemActionSeverity(values []AuthorizationSystemActionSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthorizationSystemActionSeverity) isMultiValue() bool {
    return false
}
