package graph
import (
    "strings"
    "errors"
)
type RunState int

const (
    UNKNOWN_RUNSTATE RunState = iota
    SUCCESS_RUNSTATE
    FAIL_RUNSTATE
    SCRIPTERROR_RUNSTATE
    PENDING_RUNSTATE
    NOTAPPLICABLE_RUNSTATE
)

func (i RunState) String() string {
    return []string{"UNKNOWN", "SUCCESS", "FAIL", "SCRIPTERROR", "PENDING", "NOTAPPLICABLE"}[i]
}
func ParseRunState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_RUNSTATE, nil
        case "SUCCESS":
            return SUCCESS_RUNSTATE, nil
        case "FAIL":
            return FAIL_RUNSTATE, nil
        case "SCRIPTERROR":
            return SCRIPTERROR_RUNSTATE, nil
        case "PENDING":
            return PENDING_RUNSTATE, nil
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_RUNSTATE, nil
    }
    return 0, errors.New("Unknown RunState value: " + v)
}
func SerializeRunState(values []RunState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
