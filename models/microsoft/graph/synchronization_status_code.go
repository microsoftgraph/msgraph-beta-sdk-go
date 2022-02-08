package graph
import (
    "strings"
    "errors"
)
// 
type SynchronizationStatusCode int

const (
    NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE SynchronizationStatusCode = iota
    NOTRUN_SYNCHRONIZATIONSTATUSCODE
    ACTIVE_SYNCHRONIZATIONSTATUSCODE
    PAUSED_SYNCHRONIZATIONSTATUSCODE
    QUARANTINE_SYNCHRONIZATIONSTATUSCODE
)

func (i SynchronizationStatusCode) String() string {
    return []string{"NOTCONFIGURED", "NOTRUN", "ACTIVE", "PAUSED", "QUARANTINE"}[i]
}
func ParseSynchronizationStatusCode(v string) (interface{}, error) {
    result := NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE
        case "NOTRUN":
            result = NOTRUN_SYNCHRONIZATIONSTATUSCODE
        case "ACTIVE":
            result = ACTIVE_SYNCHRONIZATIONSTATUSCODE
        case "PAUSED":
            result = PAUSED_SYNCHRONIZATIONSTATUSCODE
        case "QUARANTINE":
            result = QUARANTINE_SYNCHRONIZATIONSTATUSCODE
        default:
            return 0, errors.New("Unknown SynchronizationStatusCode value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationStatusCode(values []SynchronizationStatusCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
