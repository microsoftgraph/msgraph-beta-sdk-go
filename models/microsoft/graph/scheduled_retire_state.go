package graph
import (
    "strings"
    "errors"
)
type ScheduledRetireState int

const (
    CANCELRETIRE_SCHEDULEDRETIRESTATE ScheduledRetireState = iota
    COMFIRMRETIRE_SCHEDULEDRETIRESTATE
)

func (i ScheduledRetireState) String() string {
    return []string{"CANCELRETIRE", "COMFIRMRETIRE"}[i]
}
func ParseScheduledRetireState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CANCELRETIRE":
            return CANCELRETIRE_SCHEDULEDRETIRESTATE, nil
        case "COMFIRMRETIRE":
            return COMFIRMRETIRE_SCHEDULEDRETIRESTATE, nil
    }
    return 0, errors.New("Unknown ScheduledRetireState value: " + v)
}
func SerializeScheduledRetireState(values []ScheduledRetireState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
