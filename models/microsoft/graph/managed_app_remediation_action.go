package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppRemediationAction int

const (
    BLOCK_MANAGEDAPPREMEDIATIONACTION ManagedAppRemediationAction = iota
    WIPE_MANAGEDAPPREMEDIATIONACTION
    WARN_MANAGEDAPPREMEDIATIONACTION
)

func (i ManagedAppRemediationAction) String() string {
    return []string{"BLOCK", "WIPE", "WARN"}[i]
}
func ParseManagedAppRemediationAction(v string) (interface{}, error) {
    result := BLOCK_MANAGEDAPPREMEDIATIONACTION
    switch strings.ToUpper(v) {
        case "BLOCK":
            result = BLOCK_MANAGEDAPPREMEDIATIONACTION
        case "WIPE":
            result = WIPE_MANAGEDAPPREMEDIATIONACTION
        case "WARN":
            result = WARN_MANAGEDAPPREMEDIATIONACTION
        default:
            return 0, errors.New("Unknown ManagedAppRemediationAction value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppRemediationAction(values []ManagedAppRemediationAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
