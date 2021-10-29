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
    switch strings.ToUpper(v) {
        case "ALLOW":
            return ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION, nil
        case "BLOCKORGANIZATIONALDATA":
            return BLOCKORGANIZATIONALDATA_MANAGEDAPPNOTIFICATIONRESTRICTION, nil
        case "BLOCK":
            return BLOCK_MANAGEDAPPNOTIFICATIONRESTRICTION, nil
    }
    return 0, errors.New("Unknown ManagedAppNotificationRestriction value: " + v)
}
func SerializeManagedAppNotificationRestriction(values []ManagedAppNotificationRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
