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
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE, nil
        case "NOTRUN":
            return NOTRUN_SYNCHRONIZATIONSTATUSCODE, nil
        case "ACTIVE":
            return ACTIVE_SYNCHRONIZATIONSTATUSCODE, nil
        case "PAUSED":
            return PAUSED_SYNCHRONIZATIONSTATUSCODE, nil
        case "QUARANTINE":
            return QUARANTINE_SYNCHRONIZATIONSTATUSCODE, nil
    }
    return 0, errors.New("Unknown SynchronizationStatusCode value: " + v)
}
func SerializeSynchronizationStatusCode(values []SynchronizationStatusCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
