package graph
import (
    "strings"
    "errors"
)
// 
type RemediationState int

const (
    UNKNOWN_REMEDIATIONSTATE RemediationState = iota
    SKIPPED_REMEDIATIONSTATE
    SUCCESS_REMEDIATIONSTATE
    REMEDIATIONFAILED_REMEDIATIONSTATE
    SCRIPTERROR_REMEDIATIONSTATE
)

func (i RemediationState) String() string {
    return []string{"UNKNOWN", "SKIPPED", "SUCCESS", "REMEDIATIONFAILED", "SCRIPTERROR"}[i]
}
func ParseRemediationState(v string) (interface{}, error) {
    result := UNKNOWN_REMEDIATIONSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_REMEDIATIONSTATE
        case "SKIPPED":
            result = SKIPPED_REMEDIATIONSTATE
        case "SUCCESS":
            result = SUCCESS_REMEDIATIONSTATE
        case "REMEDIATIONFAILED":
            result = REMEDIATIONFAILED_REMEDIATIONSTATE
        case "SCRIPTERROR":
            result = SCRIPTERROR_REMEDIATIONSTATE
        default:
            return 0, errors.New("Unknown RemediationState value: " + v)
    }
    return &result, nil
}
func SerializeRemediationState(values []RemediationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
