package networkaccess
type AlertSeverity int

const (
    INFORMATIONAL_ALERTSEVERITY AlertSeverity = iota
    LOW_ALERTSEVERITY
    MEDIUM_ALERTSEVERITY
    HIGH_ALERTSEVERITY
    UNKNOWNFUTUREVALUE_ALERTSEVERITY
)

func (i AlertSeverity) String() string {
    return []string{"informational", "low", "medium", "high", "unknownFutureValue"}[i]
}
func ParseAlertSeverity(v string) (any, error) {
    result := INFORMATIONAL_ALERTSEVERITY
    switch v {
        case "informational":
            result = INFORMATIONAL_ALERTSEVERITY
        case "low":
            result = LOW_ALERTSEVERITY
        case "medium":
            result = MEDIUM_ALERTSEVERITY
        case "high":
            result = HIGH_ALERTSEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTSEVERITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlertSeverity(values []AlertSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlertSeverity) isMultiValue() bool {
    return false
}
