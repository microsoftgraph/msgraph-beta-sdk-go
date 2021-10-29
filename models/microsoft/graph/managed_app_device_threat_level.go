package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppDeviceThreatLevel int

const (
    NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL ManagedAppDeviceThreatLevel = iota
    SECURED_MANAGEDAPPDEVICETHREATLEVEL
    LOW_MANAGEDAPPDEVICETHREATLEVEL
    MEDIUM_MANAGEDAPPDEVICETHREATLEVEL
    HIGH_MANAGEDAPPDEVICETHREATLEVEL
)

func (i ManagedAppDeviceThreatLevel) String() string {
    return []string{"NOTCONFIGURED", "SECURED", "LOW", "MEDIUM", "HIGH"}[i]
}
func ParseManagedAppDeviceThreatLevel(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL, nil
        case "SECURED":
            return SECURED_MANAGEDAPPDEVICETHREATLEVEL, nil
        case "LOW":
            return LOW_MANAGEDAPPDEVICETHREATLEVEL, nil
        case "MEDIUM":
            return MEDIUM_MANAGEDAPPDEVICETHREATLEVEL, nil
        case "HIGH":
            return HIGH_MANAGEDAPPDEVICETHREATLEVEL, nil
    }
    return 0, errors.New("Unknown ManagedAppDeviceThreatLevel value: " + v)
}
func SerializeManagedAppDeviceThreatLevel(values []ManagedAppDeviceThreatLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
