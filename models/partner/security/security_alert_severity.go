package security
type SecurityAlertSeverity int

const (
    INFORMATIONAL_SECURITYALERTSEVERITY SecurityAlertSeverity = iota
    HIGH_SECURITYALERTSEVERITY
    MEDIUM_SECURITYALERTSEVERITY
    LOW_SECURITYALERTSEVERITY
    UNKNOWNFUTUREVALUE_SECURITYALERTSEVERITY
)

func (i SecurityAlertSeverity) String() string {
    return []string{"informational", "high", "medium", "low", "unknownFutureValue"}[i]
}
func ParseSecurityAlertSeverity(v string) (any, error) {
    result := INFORMATIONAL_SECURITYALERTSEVERITY
    switch v {
        case "informational":
            result = INFORMATIONAL_SECURITYALERTSEVERITY
        case "high":
            result = HIGH_SECURITYALERTSEVERITY
        case "medium":
            result = MEDIUM_SECURITYALERTSEVERITY
        case "low":
            result = LOW_SECURITYALERTSEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYALERTSEVERITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityAlertSeverity(values []SecurityAlertSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAlertSeverity) isMultiValue() bool {
    return false
}
