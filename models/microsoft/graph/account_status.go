package graph
import (
    "strings"
    "errors"
)
// 
type AccountStatus int

const (
    UNKNOWN_ACCOUNTSTATUS AccountStatus = iota
    STAGED_ACCOUNTSTATUS
    ACTIVE_ACCOUNTSTATUS
    SUSPENDED_ACCOUNTSTATUS
    DELETED_ACCOUNTSTATUS
    UNKNOWNFUTUREVALUE_ACCOUNTSTATUS
)

func (i AccountStatus) String() string {
    return []string{"UNKNOWN", "STAGED", "ACTIVE", "SUSPENDED", "DELETED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccountStatus(v string) (interface{}, error) {
    result := UNKNOWN_ACCOUNTSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ACCOUNTSTATUS
        case "STAGED":
            result = STAGED_ACCOUNTSTATUS
        case "ACTIVE":
            result = ACTIVE_ACCOUNTSTATUS
        case "SUSPENDED":
            result = SUSPENDED_ACCOUNTSTATUS
        case "DELETED":
            result = DELETED_ACCOUNTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCOUNTSTATUS
        default:
            return 0, errors.New("Unknown AccountStatus value: " + v)
    }
    return &result, nil
}
func SerializeAccountStatus(values []AccountStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
