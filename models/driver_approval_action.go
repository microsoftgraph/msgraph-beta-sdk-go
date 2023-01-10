package models
import (
    "errors"
)
// An enum type to represent approval actions of single or list of drivers.
type DriverApprovalAction int

const (
    // This indicates the action to approve single or list of drivers.
    APPROVE_DRIVERAPPROVALACTION DriverApprovalAction = iota
    // This indicates the action to approve single or list of drivers.
    DECLINE_DRIVERAPPROVALACTION
    // This indicates the action to suspend single or list of drivers.
    SUSPEND_DRIVERAPPROVALACTION
)

func (i DriverApprovalAction) String() string {
    return []string{"approve", "decline", "suspend"}[i]
}
func ParseDriverApprovalAction(v string) (interface{}, error) {
    result := APPROVE_DRIVERAPPROVALACTION
    switch v {
        case "approve":
            result = APPROVE_DRIVERAPPROVALACTION
        case "decline":
            result = DECLINE_DRIVERAPPROVALACTION
        case "suspend":
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
