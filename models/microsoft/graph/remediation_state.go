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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_REMEDIATIONSTATE, nil
        case "SKIPPED":
            return SKIPPED_REMEDIATIONSTATE, nil
        case "SUCCESS":
            return SUCCESS_REMEDIATIONSTATE, nil
        case "REMEDIATIONFAILED":
            return REMEDIATIONFAILED_REMEDIATIONSTATE, nil
        case "SCRIPTERROR":
            return SCRIPTERROR_REMEDIATIONSTATE, nil
    }
    return 0, errors.New("Unknown RemediationState value: " + v)
}
func SerializeRemediationState(values []RemediationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
