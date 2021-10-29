package graph
import (
    "strings"
    "errors"
)
// 
type DriverApprovalAction int

const (
    APPROVE_DRIVERAPPROVALACTION DriverApprovalAction = iota
    DECLINE_DRIVERAPPROVALACTION
    SUSPEND_DRIVERAPPROVALACTION
)

func (i DriverApprovalAction) String() string {
    return []string{"APPROVE", "DECLINE", "SUSPEND"}[i]
}
func ParseDriverApprovalAction(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "APPROVE":
            return APPROVE_DRIVERAPPROVALACTION, nil
        case "DECLINE":
            return DECLINE_DRIVERAPPROVALACTION, nil
        case "SUSPEND":
            return SUSPEND_DRIVERAPPROVALACTION, nil
    }
    return 0, errors.New("Unknown DriverApprovalAction value: " + v)
}
func SerializeDriverApprovalAction(values []DriverApprovalAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
