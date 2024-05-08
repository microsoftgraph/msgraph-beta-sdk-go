package security
import (
    "errors"
)
type SecurityAlertStatus int

const (
    ACTIVE_SECURITYALERTSTATUS SecurityAlertStatus = iota
    RESOLVED_SECURITYALERTSTATUS
    INVESTIGATING_SECURITYALERTSTATUS
    UNKNOWNFUTUREVALUE_SECURITYALERTSTATUS
)

func (i SecurityAlertStatus) String() string {
    return []string{"active", "resolved", "investigating", "unknownFutureValue"}[i]
}
func ParseSecurityAlertStatus(v string) (any, error) {
    result := ACTIVE_SECURITYALERTSTATUS
    switch v {
        case "active":
            result = ACTIVE_SECURITYALERTSTATUS
        case "resolved":
            result = RESOLVED_SECURITYALERTSTATUS
        case "investigating":
            result = INVESTIGATING_SECURITYALERTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYALERTSTATUS
        default:
            return 0, errors.New("Unknown SecurityAlertStatus value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAlertStatus(values []SecurityAlertStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAlertStatus) isMultiValue() bool {
    return false
}
