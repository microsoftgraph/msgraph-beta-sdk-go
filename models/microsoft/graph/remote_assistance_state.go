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
    result := DISABLED_REMOTEASSISTANCESTATE
    switch strings.ToUpper(v) {
        case "DISABLED":
            result = DISABLED_REMOTEASSISTANCESTATE
        case "ENABLED":
            result = ENABLED_REMOTEASSISTANCESTATE
        default:
            return 0, errors.New("Unknown RemoteAssistanceState value: " + v)
    }
    return &result, nil
}
func SerializeRemoteAssistanceState(values []RemoteAssistanceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
