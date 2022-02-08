package graph
import (
    "strings"
    "errors"
)
// 
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
    result := UNKNOWN_SECURITYBASELINECOMPLIANCESTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SECURITYBASELINECOMPLIANCESTATE
        case "SECURE":
            result = SECURE_SECURITYBASELINECOMPLIANCESTATE
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_SECURITYBASELINECOMPLIANCESTATE
        case "NOTSECURE":
            result = NOTSECURE_SECURITYBASELINECOMPLIANCESTATE
        case "ERROR":
            result = ERROR_SECURITYBASELINECOMPLIANCESTATE
        case "CONFLICT":
            result = CONFLICT_SECURITYBASELINECOMPLIANCESTATE
        default:
            return 0, errors.New("Unknown SecurityBaselineComplianceState value: " + v)
    }
    return &result, nil
}
func SerializeSecurityBaselineComplianceState(values []SecurityBaselineComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
