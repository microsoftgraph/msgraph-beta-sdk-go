package graph
import (
    "strings"
    "errors"
)
// 
type AlertSeverity_v2 int

const (
    UNKNOWN_ALERTSEVERITY_V2 AlertSeverity_v2 = iota
    INFORMATIONAL_ALERTSEVERITY_V2
    LOW_ALERTSEVERITY_V2
    MEDIUM_ALERTSEVERITY_V2
    HIGH_ALERTSEVERITY_V2
    UNKNOWNFUTUREVALUE_ALERTSEVERITY_V2
)

func (i AlertSeverity_v2) String() string {
    return []string{"UNKNOWN", "INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertSeverity_v2(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ALERTSEVERITY_V2, nil
        case "INFORMATIONAL":
            return INFORMATIONAL_ALERTSEVERITY_V2, nil
        case "LOW":
            return LOW_ALERTSEVERITY_V2, nil
        case "MEDIUM":
            return MEDIUM_ALERTSEVERITY_V2, nil
        case "HIGH":
            return HIGH_ALERTSEVERITY_V2, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALERTSEVERITY_V2, nil
    }
    return 0, errors.New("Unknown AlertSeverity_v2 value: " + v)
}
func SerializeAlertSeverity_v2(values []AlertSeverity_v2) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
