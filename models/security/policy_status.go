package security
import (
    "strings"
    "errors"
)
// Provides operations to manage the cases property of the microsoft.graph.security entity.
type PolicyStatus int

const (
    PENDING_POLICYSTATUS PolicyStatus = iota
    ERROR_POLICYSTATUS
    SUCCESS_POLICYSTATUS
    UNKNOWNFUTUREVALUE_POLICYSTATUS
)

func (i PolicyStatus) String() string {
    return []string{"PENDING", "ERROR", "SUCCESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePolicyStatus(v string) (interface{}, error) {
    result := PENDING_POLICYSTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_POLICYSTATUS
        case "ERROR":
            result = ERROR_POLICYSTATUS
        case "SUCCESS":
            result = SUCCESS_POLICYSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_POLICYSTATUS
        default:
            return 0, errors.New("Unknown PolicyStatus value: " + v)
    }
    return &result, nil
}
func SerializePolicyStatus(values []PolicyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
