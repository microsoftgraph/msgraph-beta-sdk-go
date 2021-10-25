package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_POLICYSETSTATUS, nil
        case "VALIDATING":
            return VALIDATING_POLICYSETSTATUS, nil
        case "PARTIALSUCCESS":
            return PARTIALSUCCESS_POLICYSETSTATUS, nil
        case "SUCCESS":
            return SUCCESS_POLICYSETSTATUS, nil
        case "ERROR":
            return ERROR_POLICYSETSTATUS, nil
        case "NOTASSIGNED":
            return NOTASSIGNED_POLICYSETSTATUS, nil
    }
    return 0, errors.New("Unknown PolicySetStatus value: " + v)
}
func SerializePolicySetStatus(values []PolicySetStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
