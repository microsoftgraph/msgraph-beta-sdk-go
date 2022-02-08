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
    result := APPROVE_DRIVERAPPROVALACTION
    switch strings.ToUpper(v) {
        case "APPROVE":
            result = APPROVE_DRIVERAPPROVALACTION
        case "DECLINE":
            result = DECLINE_DRIVERAPPROVALACTION
        case "SUSPEND":
            result = SUSPEND_DRIVERAPPROVALACTION
        default:
            return 0, errors.New("Unknown DriverApprovalAction value: " + v)
    }
    return &result, nil
}
func SerializeDriverApprovalAction(values []DriverApprovalAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
