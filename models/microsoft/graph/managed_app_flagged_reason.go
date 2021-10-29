package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppFlaggedReason int

const (
    NONE_MANAGEDAPPFLAGGEDREASON ManagedAppFlaggedReason = iota
    ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON
    ANDROIDBOOTLOADERUNLOCKED_MANAGEDAPPFLAGGEDREASON
    ANDROIDFACTORYROMMODIFIED_MANAGEDAPPFLAGGEDREASON
)

func (i ManagedAppFlaggedReason) String() string {
    return []string{"NONE", "ROOTEDDEVICE", "ANDROIDBOOTLOADERUNLOCKED", "ANDROIDFACTORYROMMODIFIED"}[i]
}
func ParseManagedAppFlaggedReason(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_MANAGEDAPPFLAGGEDREASON, nil
        case "ROOTEDDEVICE":
            return ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON, nil
        case "ANDROIDBOOTLOADERUNLOCKED":
            return ANDROIDBOOTLOADERUNLOCKED_MANAGEDAPPFLAGGEDREASON, nil
        case "ANDROIDFACTORYROMMODIFIED":
            return ANDROIDFACTORYROMMODIFIED_MANAGEDAPPFLAGGEDREASON, nil
    }
    return 0, errors.New("Unknown ManagedAppFlaggedReason value: " + v)
}
func SerializeManagedAppFlaggedReason(values []ManagedAppFlaggedReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
