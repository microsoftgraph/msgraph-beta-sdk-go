package graph
import (
    "strings"
    "errors"
)
// 
type M365AlertSeverity int

const (
    UNKNOWN_M365ALERTSEVERITY M365AlertSeverity = iota
    INFORMATIONAL_M365ALERTSEVERITY
    LOW_M365ALERTSEVERITY
    MEDIUM_M365ALERTSEVERITY
    HIGH_M365ALERTSEVERITY
    UNKNOWNFUTUREVALUE_M365ALERTSEVERITY
)

func (i M365AlertSeverity) String() string {
    return []string{"UNKNOWN", "INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseM365AlertSeverity(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_M365ALERTSEVERITY, nil
        case "INFORMATIONAL":
            return INFORMATIONAL_M365ALERTSEVERITY, nil
        case "LOW":
            return LOW_M365ALERTSEVERITY, nil
        case "MEDIUM":
            return MEDIUM_M365ALERTSEVERITY, nil
        case "HIGH":
            return HIGH_M365ALERTSEVERITY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_M365ALERTSEVERITY, nil
    }
    return 0, errors.New("Unknown M365AlertSeverity value: " + v)
}
func SerializeM365AlertSeverity(values []M365AlertSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
