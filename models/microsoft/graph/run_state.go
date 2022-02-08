package graph
import (
    "strings"
    "errors"
)
// 
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
    result := UNKNOWN_RUNSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_RUNSTATE
        case "SUCCESS":
            result = SUCCESS_RUNSTATE
        case "FAIL":
            result = FAIL_RUNSTATE
        case "SCRIPTERROR":
            result = SCRIPTERROR_RUNSTATE
        case "PENDING":
            result = PENDING_RUNSTATE
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_RUNSTATE
        default:
            return 0, errors.New("Unknown RunState value: " + v)
    }
    return &result, nil
}
func SerializeRunState(values []RunState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
