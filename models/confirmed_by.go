package models
import (
    "errors"
    "strings"
)
// 
type ConfirmedBy int

const (
    NONE_CONFIRMEDBY ConfirmedBy = iota
    USER_CONFIRMEDBY
    MANAGER_CONFIRMEDBY
    UNKNOWNFUTUREVALUE_CONFIRMEDBY
)

func (i ConfirmedBy) String() string {
    var values []string
    for p := ConfirmedBy(1); p <= UNKNOWNFUTUREVALUE_CONFIRMEDBY; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "user", "manager", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseConfirmedBy(v string) (any, error) {
    var result ConfirmedBy
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CONFIRMEDBY
            case "user":
                result |= USER_CONFIRMEDBY
            case "manager":
                result |= MANAGER_CONFIRMEDBY
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONFIRMEDBY
            default:
                return 0, errors.New("Unknown ConfirmedBy value: " + v)
        }
    }
    return &result, nil
}
func SerializeConfirmedBy(values []ConfirmedBy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfirmedBy) isMultiValue() bool {
    return true
}
