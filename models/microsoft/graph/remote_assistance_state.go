package graph
import (
    "strings"
    "errors"
)
// 
type RemoteAssistanceState int

const (
    DISABLED_REMOTEASSISTANCESTATE RemoteAssistanceState = iota
    ENABLED_REMOTEASSISTANCESTATE
)

func (i RemoteAssistanceState) String() string {
    return []string{"DISABLED", "ENABLED"}[i]
}
func ParseRemoteAssistanceState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
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
