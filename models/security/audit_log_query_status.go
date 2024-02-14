package security
import (
    "errors"
)
type AuditLogQueryStatus int

const (
    NOTSTARTED_AUDITLOGQUERYSTATUS AuditLogQueryStatus = iota
    RUNNING_AUDITLOGQUERYSTATUS
    SUCCEEDED_AUDITLOGQUERYSTATUS
    FAILED_AUDITLOGQUERYSTATUS
    CANCELLED_AUDITLOGQUERYSTATUS
    UNKNOWNFUTUREVALUE_AUDITLOGQUERYSTATUS
)

func (i AuditLogQueryStatus) String() string {
    return []string{"notStarted", "running", "succeeded", "failed", "cancelled", "unknownFutureValue"}[i]
}
func ParseAuditLogQueryStatus(v string) (any, error) {
    result := NOTSTARTED_AUDITLOGQUERYSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_AUDITLOGQUERYSTATUS
        case "running":
            result = RUNNING_AUDITLOGQUERYSTATUS
        case "succeeded":
            result = SUCCEEDED_AUDITLOGQUERYSTATUS
        case "failed":
            result = FAILED_AUDITLOGQUERYSTATUS
        case "cancelled":
            result = CANCELLED_AUDITLOGQUERYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUDITLOGQUERYSTATUS
        default:
            return 0, errors.New("Unknown AuditLogQueryStatus value: " + v)
    }
    return &result, nil
}
func SerializeAuditLogQueryStatus(values []AuditLogQueryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuditLogQueryStatus) isMultiValue() bool {
    return false
}
