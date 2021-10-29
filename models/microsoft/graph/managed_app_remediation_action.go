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
    switch strings.ToUpper(v) {
        case "BLOCK":
            return BLOCK_MANAGEDAPPREMEDIATIONACTION, nil
        case "WIPE":
            return WIPE_MANAGEDAPPREMEDIATIONACTION, nil
        case "WARN":
            return WARN_MANAGEDAPPREMEDIATIONACTION, nil
    }
    return 0, errors.New("Unknown ManagedAppRemediationAction value: " + v)
}
func SerializeManagedAppRemediationAction(values []ManagedAppRemediationAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
