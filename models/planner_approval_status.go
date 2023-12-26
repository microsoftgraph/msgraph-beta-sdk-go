package models
import (
    "errors"
)
// 
type PlannerApprovalStatus int

const (
    REQUESTED_PLANNERAPPROVALSTATUS PlannerApprovalStatus = iota
    APPROVED_PLANNERAPPROVALSTATUS
    REJECTED_PLANNERAPPROVALSTATUS
    CANCELLED_PLANNERAPPROVALSTATUS
    UNKNOWNFUTUREVALUE_PLANNERAPPROVALSTATUS
)

func (i PlannerApprovalStatus) String() string {
    return []string{"requested", "approved", "rejected", "cancelled", "unknownFutureValue"}[i]
}
func ParsePlannerApprovalStatus(v string) (any, error) {
    result := REQUESTED_PLANNERAPPROVALSTATUS
    switch v {
        case "requested":
            result = REQUESTED_PLANNERAPPROVALSTATUS
        case "approved":
            result = APPROVED_PLANNERAPPROVALSTATUS
        case "rejected":
            result = REJECTED_PLANNERAPPROVALSTATUS
        case "cancelled":
            result = CANCELLED_PLANNERAPPROVALSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERAPPROVALSTATUS
        default:
            return 0, errors.New("Unknown PlannerApprovalStatus value: " + v)
    }
    return &result, nil
}
func SerializePlannerApprovalStatus(values []PlannerApprovalStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerApprovalStatus) isMultiValue() bool {
    return false
}
