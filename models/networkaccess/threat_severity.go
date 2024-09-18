package networkaccess
type ThreatSeverity int

const (
    LOW_THREATSEVERITY ThreatSeverity = iota
    MEDIUM_THREATSEVERITY
    HIGH_THREATSEVERITY
    CRITICAL_THREATSEVERITY
    UNKNOWNFUTUREVALUE_THREATSEVERITY
)

func (i ThreatSeverity) String() string {
    return []string{"low", "medium", "high", "critical", "unknownFutureValue"}[i]
}
func ParseThreatSeverity(v string) (any, error) {
    result := LOW_THREATSEVERITY
    switch v {
        case "low":
            result = LOW_THREATSEVERITY
        case "medium":
            result = MEDIUM_THREATSEVERITY
        case "high":
            result = HIGH_THREATSEVERITY
        case "critical":
            result = CRITICAL_THREATSEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_THREATSEVERITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeThreatSeverity(values []ThreatSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ThreatSeverity) isMultiValue() bool {
    return false
}
