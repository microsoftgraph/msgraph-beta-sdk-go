package security
import (
    "errors"
)
type SecurityAlertResolvedReason int

const (
    LEGITIMATE_SECURITYALERTRESOLVEDREASON SecurityAlertResolvedReason = iota
    IGNORE_SECURITYALERTRESOLVEDREASON
    FRAUD_SECURITYALERTRESOLVEDREASON
    UNKNOWNFUTUREVALUE_SECURITYALERTRESOLVEDREASON
)

func (i SecurityAlertResolvedReason) String() string {
    return []string{"legitimate", "ignore", "fraud", "unknownFutureValue"}[i]
}
func ParseSecurityAlertResolvedReason(v string) (any, error) {
    result := LEGITIMATE_SECURITYALERTRESOLVEDREASON
    switch v {
        case "legitimate":
            result = LEGITIMATE_SECURITYALERTRESOLVEDREASON
        case "ignore":
            result = IGNORE_SECURITYALERTRESOLVEDREASON
        case "fraud":
            result = FRAUD_SECURITYALERTRESOLVEDREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYALERTRESOLVEDREASON
        default:
            return 0, errors.New("Unknown SecurityAlertResolvedReason value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAlertResolvedReason(values []SecurityAlertResolvedReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAlertResolvedReason) isMultiValue() bool {
    return false
}
