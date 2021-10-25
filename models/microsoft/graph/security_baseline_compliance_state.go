package graph
import (
    "strings"
    "errors"
)
type SecurityBaselineComplianceState int

const (
    UNKNOWN_SECURITYBASELINECOMPLIANCESTATE SecurityBaselineComplianceState = iota
    SECURE_SECURITYBASELINECOMPLIANCESTATE
    NOTAPPLICABLE_SECURITYBASELINECOMPLIANCESTATE
    NOTSECURE_SECURITYBASELINECOMPLIANCESTATE
    ERROR_SECURITYBASELINECOMPLIANCESTATE
    CONFLICT_SECURITYBASELINECOMPLIANCESTATE
)

func (i SecurityBaselineComplianceState) String() string {
    return []string{"UNKNOWN", "SECURE", "NOTAPPLICABLE", "NOTSECURE", "ERROR", "CONFLICT"}[i]
}
func ParseSecurityBaselineComplianceState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SECURITYBASELINECOMPLIANCESTATE, nil
        case "SECURE":
            return SECURE_SECURITYBASELINECOMPLIANCESTATE, nil
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_SECURITYBASELINECOMPLIANCESTATE, nil
        case "NOTSECURE":
            return NOTSECURE_SECURITYBASELINECOMPLIANCESTATE, nil
        case "ERROR":
            return ERROR_SECURITYBASELINECOMPLIANCESTATE, nil
        case "CONFLICT":
            return CONFLICT_SECURITYBASELINECOMPLIANCESTATE, nil
    }
    return 0, errors.New("Unknown SecurityBaselineComplianceState value: " + v)
}
func SerializeSecurityBaselineComplianceState(values []SecurityBaselineComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
