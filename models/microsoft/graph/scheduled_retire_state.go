package graph
import (
    "strings"
    "errors"
)
// 
type ScheduledRetireState int

const (
    CANCELRETIRE_SCHEDULEDRETIRESTATE ScheduledRetireState = iota
    COMFIRMRETIRE_SCHEDULEDRETIRESTATE
)

func (i ScheduledRetireState) String() string {
    return []string{"CANCELRETIRE", "COMFIRMRETIRE"}[i]
}
func ParseScheduledRetireState(v string) (interface{}, error) {
    result := CANCELRETIRE_SCHEDULEDRETIRESTATE
    switch strings.ToUpper(v) {
        case "CANCELRETIRE":
            result = CANCELRETIRE_SCHEDULEDRETIRESTATE
        case "COMFIRMRETIRE":
            result = COMFIRMRETIRE_SCHEDULEDRETIRESTATE
        default:
            return 0, errors.New("Unknown ScheduledRetireState value: " + v)
    }
    return &result, nil
}
func SerializeScheduledRetireState(values []ScheduledRetireState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
