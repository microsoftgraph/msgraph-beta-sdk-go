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
    switch strings.ToUpper(v) {
        case "NEEDSREVIEW":
            return NEEDSREVIEW_DRIVERAPPROVALSTATUS, nil
        case "DECLINED":
            return DECLINED_DRIVERAPPROVALSTATUS, nil
        case "APPROVED":
            return APPROVED_DRIVERAPPROVALSTATUS, nil
        case "SUSPENDED":
            return SUSPENDED_DRIVERAPPROVALSTATUS, nil
    }
    return 0, errors.New("Unknown DriverApprovalStatus value: " + v)
}
func SerializeDriverApprovalStatus(values []DriverApprovalStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
