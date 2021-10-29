package graph
import (
    "strings"
    "errors"
)
// 
type SynchronizationScheduleState int

const (
    ACTIVE_SYNCHRONIZATIONSCHEDULESTATE SynchronizationScheduleState = iota
    DISABLED_SYNCHRONIZATIONSCHEDULESTATE
    PAUSED_SYNCHRONIZATIONSCHEDULESTATE
)

func (i SynchronizationScheduleState) String() string {
    return []string{"ACTIVE", "DISABLED", "PAUSED"}[i]
}
func ParseSynchronizationScheduleState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_SYNCHRONIZATIONSCHEDULESTATE, nil
        case "DISABLED":
            return DISABLED_SYNCHRONIZATIONSCHEDULESTATE, nil
        case "PAUSED":
            return PAUSED_SYNCHRONIZATIONSCHEDULESTATE, nil
    }
    return 0, errors.New("Unknown SynchronizationScheduleState value: " + v)
}
func SerializeSynchronizationScheduleState(values []SynchronizationScheduleState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
