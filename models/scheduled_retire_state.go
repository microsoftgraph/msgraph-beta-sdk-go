package models
import (
    "errors"
)
// Cancel or confirm scheduled retire 
type ScheduledRetireState int

const (
    // CancelRetire
    CANCELRETIRE_SCHEDULEDRETIRESTATE ScheduledRetireState = iota
    // ConfirmRetire
    CONFIRMRETIRE_SCHEDULEDRETIRESTATE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_SCHEDULEDRETIRESTATE
)

func (i ScheduledRetireState) String() string {
    return []string{"cancelRetire", "confirmRetire", "unknownFutureValue"}[i]
}
func ParseScheduledRetireState(v string) (any, error) {
    result := CANCELRETIRE_SCHEDULEDRETIRESTATE
    switch v {
        case "cancelRetire":
            result = CANCELRETIRE_SCHEDULEDRETIRESTATE
        case "confirmRetire":
            result = CONFIRMRETIRE_SCHEDULEDRETIRESTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SCHEDULEDRETIRESTATE
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
