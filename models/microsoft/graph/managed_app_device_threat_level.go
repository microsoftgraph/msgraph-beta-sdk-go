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
    result := NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL
        case "SECURED":
            result = SECURED_MANAGEDAPPDEVICETHREATLEVEL
        case "LOW":
            result = LOW_MANAGEDAPPDEVICETHREATLEVEL
        case "MEDIUM":
            result = MEDIUM_MANAGEDAPPDEVICETHREATLEVEL
        case "HIGH":
            result = HIGH_MANAGEDAPPDEVICETHREATLEVEL
        default:
            return 0, errors.New("Unknown ManagedAppDeviceThreatLevel value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDeviceThreatLevel(values []ManagedAppDeviceThreatLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
