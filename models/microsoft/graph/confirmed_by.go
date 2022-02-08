package graph
import (
    "strings"
    "errors"
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
    return []string{"NONE", "USER", "MANAGER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConfirmedBy(v string) (interface{}, error) {
    result := NONE_CONFIRMEDBY
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CONFIRMEDBY
        case "USER":
            result = USER_CONFIRMEDBY
        case "MANAGER":
            result = MANAGER_CONFIRMEDBY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONFIRMEDBY
        default:
            return 0, errors.New("Unknown ConfirmedBy value: " + v)
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
