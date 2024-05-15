package security
type RemediationSeverity int

const (
    LOW_REMEDIATIONSEVERITY RemediationSeverity = iota
    MEDIUM_REMEDIATIONSEVERITY
    HIGH_REMEDIATIONSEVERITY
    UNKNOWNFUTUREVALUE_REMEDIATIONSEVERITY
)

func (i RemediationSeverity) String() string {
    return []string{"low", "medium", "high", "unknownFutureValue"}[i]
}
func ParseRemediationSeverity(v string) (any, error) {
    result := LOW_REMEDIATIONSEVERITY
    switch v {
        case "low":
            result = LOW_REMEDIATIONSEVERITY
        case "medium":
            result = MEDIUM_REMEDIATIONSEVERITY
        case "high":
            result = HIGH_REMEDIATIONSEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REMEDIATIONSEVERITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRemediationSeverity(values []RemediationSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RemediationSeverity) isMultiValue() bool {
    return false
}
