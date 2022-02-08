package graph
import (
    "strings"
    "errors"
)
// 
type DriverApprovalStatus int

const (
    NEEDSREVIEW_DRIVERAPPROVALSTATUS DriverApprovalStatus = iota
    DECLINED_DRIVERAPPROVALSTATUS
    APPROVED_DRIVERAPPROVALSTATUS
    SUSPENDED_DRIVERAPPROVALSTATUS
)

func (i DriverApprovalStatus) String() string {
    return []string{"NEEDSREVIEW", "DECLINED", "APPROVED", "SUSPENDED"}[i]
}
func ParseDriverApprovalStatus(v string) (interface{}, error) {
    result := NEEDSREVIEW_DRIVERAPPROVALSTATUS
    switch strings.ToUpper(v) {
        case "NEEDSREVIEW":
            result = NEEDSREVIEW_DRIVERAPPROVALSTATUS
        case "DECLINED":
            result = DECLINED_DRIVERAPPROVALSTATUS
        case "APPROVED":
            result = APPROVED_DRIVERAPPROVALSTATUS
        case "SUSPENDED":
            result = SUSPENDED_DRIVERAPPROVALSTATUS
        default:
            return 0, errors.New("Unknown DriverApprovalStatus value: " + v)
    }
    return &result, nil
}
func SerializeDriverApprovalStatus(values []DriverApprovalStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
