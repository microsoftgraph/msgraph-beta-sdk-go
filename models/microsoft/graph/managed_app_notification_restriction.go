package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppNotificationRestriction int

const (
    ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION ManagedAppNotificationRestriction = iota
    BLOCKORGANIZATIONALDATA_MANAGEDAPPNOTIFICATIONRESTRICTION
    BLOCK_MANAGEDAPPNOTIFICATIONRESTRICTION
)

func (i ManagedAppNotificationRestriction) String() string {
    return []string{"ALLOW", "BLOCKORGANIZATIONALDATA", "BLOCK"}[i]
}
func ParseManagedAppNotificationRestriction(v string) (interface{}, error) {
    result := ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION
    switch strings.ToUpper(v) {
        case "ALLOW":
            result = ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION
        case "BLOCKORGANIZATIONALDATA":
            result = BLOCKORGANIZATIONALDATA_MANAGEDAPPNOTIFICATIONRESTRICTION
        case "BLOCK":
            result = BLOCK_MANAGEDAPPNOTIFICATIONRESTRICTION
        default:
            return 0, errors.New("Unknown ManagedAppNotificationRestriction value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppNotificationRestriction(values []ManagedAppNotificationRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
