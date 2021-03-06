package models
import (
    "errors"
)
// Provides operations to call the setScheduledRetireState method.
type ScheduledRetireState int

const (
    // Cancel retire.
    CANCELRETIRE_SCHEDULEDRETIRESTATE ScheduledRetireState = iota
    // Retire these devices.
    COMFIRMRETIRE_SCHEDULEDRETIRESTATE
)

func (i ScheduledRetireState) String() string {
    return []string{"cancelRetire", "comfirmRetire"}[i]
}
func ParseScheduledRetireState(v string) (interface{}, error) {
    result := CANCELRETIRE_SCHEDULEDRETIRESTATE
    switch v {
        case "cancelRetire":
            result = CANCELRETIRE_SCHEDULEDRETIRESTATE
        case "comfirmRetire":
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
