package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ACCOUNTSTATUS, nil
        case "STAGED":
            return STAGED_ACCOUNTSTATUS, nil
        case "ACTIVE":
            return ACTIVE_ACCOUNTSTATUS, nil
        case "SUSPENDED":
            return SUSPENDED_ACCOUNTSTATUS, nil
        case "DELETED":
            return DELETED_ACCOUNTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCOUNTSTATUS, nil
    }
    return 0, errors.New("Unknown AccountStatus value: " + v)
}
func SerializeAccountStatus(values []AccountStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
