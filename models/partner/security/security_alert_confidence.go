package security
type SecurityAlertConfidence int

const (
    LOW_SECURITYALERTCONFIDENCE SecurityAlertConfidence = iota
    MEDIUM_SECURITYALERTCONFIDENCE
    HIGH_SECURITYALERTCONFIDENCE
    UNKNOWNFUTUREVALUE_SECURITYALERTCONFIDENCE
)

func (i SecurityAlertConfidence) String() string {
    return []string{"low", "medium", "high", "unknownFutureValue"}[i]
}
func ParseSecurityAlertConfidence(v string) (any, error) {
    result := LOW_SECURITYALERTCONFIDENCE
    switch v {
        case "low":
            result = LOW_SECURITYALERTCONFIDENCE
        case "medium":
            result = MEDIUM_SECURITYALERTCONFIDENCE
        case "high":
            result = HIGH_SECURITYALERTCONFIDENCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYALERTCONFIDENCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityAlertConfidence(values []SecurityAlertConfidence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAlertConfidence) isMultiValue() bool {
    return false
}
