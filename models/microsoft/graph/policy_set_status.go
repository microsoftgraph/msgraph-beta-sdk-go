package graph
import (
    "strings"
    "errors"
)
// 
type PolicySetStatus int

const (
    UNKNOWN_POLICYSETSTATUS PolicySetStatus = iota
    VALIDATING_POLICYSETSTATUS
    PARTIALSUCCESS_POLICYSETSTATUS
    SUCCESS_POLICYSETSTATUS
    ERROR_POLICYSETSTATUS
    NOTASSIGNED_POLICYSETSTATUS
)

func (i PolicySetStatus) String() string {
    return []string{"UNKNOWN", "VALIDATING", "PARTIALSUCCESS", "SUCCESS", "ERROR", "NOTASSIGNED"}[i]
}
func ParsePolicySetStatus(v string) (interface{}, error) {
    result := UNKNOWN_POLICYSETSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_POLICYSETSTATUS
        case "VALIDATING":
            result = VALIDATING_POLICYSETSTATUS
        case "PARTIALSUCCESS":
            result = PARTIALSUCCESS_POLICYSETSTATUS
        case "SUCCESS":
            result = SUCCESS_POLICYSETSTATUS
        case "ERROR":
            result = ERROR_POLICYSETSTATUS
        case "NOTASSIGNED":
            result = NOTASSIGNED_POLICYSETSTATUS
        default:
            return 0, errors.New("Unknown PolicySetStatus value: " + v)
    }
    return &result, nil
}
func SerializePolicySetStatus(values []PolicySetStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
