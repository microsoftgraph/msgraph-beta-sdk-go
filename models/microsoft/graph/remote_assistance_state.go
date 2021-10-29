package graph
import (
    "strings"
    "errors"
)
// 
type RemoteAssistanceState int

const (
    NOTCONFIGURED_REMOTEASSISTANCESTATE RemoteAssistanceState = iota
    DISABLED_REMOTEASSISTANCESTATE
    ENABLED_REMOTEASSISTANCESTATE
)

func (i RemoteAssistanceState) String() string {
    return []string{"NOTCONFIGURED", "DISABLED", "ENABLED"}[i]
}
func ParseRemoteAssistanceState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_REMOTEASSISTANCESTATE, nil
        case "DISABLED":
            return DISABLED_REMOTEASSISTANCESTATE, nil
        case "ENABLED":
            return ENABLED_REMOTEASSISTANCESTATE, nil
    }
    return 0, errors.New("Unknown RemoteAssistanceState value: " + v)
}
func SerializeRemoteAssistanceState(values []RemoteAssistanceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
