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
    result := ACTIVE_SYNCHRONIZATIONSCHEDULESTATE
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_SYNCHRONIZATIONSCHEDULESTATE
        case "DISABLED":
            result = DISABLED_SYNCHRONIZATIONSCHEDULESTATE
        case "PAUSED":
            result = PAUSED_SYNCHRONIZATIONSCHEDULESTATE
        default:
            return 0, errors.New("Unknown SynchronizationScheduleState value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationScheduleState(values []SynchronizationScheduleState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
