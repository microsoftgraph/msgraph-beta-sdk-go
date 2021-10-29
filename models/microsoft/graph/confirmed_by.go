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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CONFIRMEDBY, nil
        case "USER":
            return USER_CONFIRMEDBY, nil
        case "MANAGER":
            return MANAGER_CONFIRMEDBY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONFIRMEDBY, nil
    }
    return 0, errors.New("Unknown ConfirmedBy value: " + v)
}
func SerializeConfirmedBy(values []ConfirmedBy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
