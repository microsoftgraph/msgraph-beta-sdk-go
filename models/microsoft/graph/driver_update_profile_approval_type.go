package graph
import (
    "strings"
    "errors"
)
// 
type DriverUpdateProfileApprovalType int

const (
    MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE DriverUpdateProfileApprovalType = iota
    AUTOMATIC_DRIVERUPDATEPROFILEAPPROVALTYPE
)

func (i DriverUpdateProfileApprovalType) String() string {
    return []string{"MANUAL", "AUTOMATIC"}[i]
}
func ParseDriverUpdateProfileApprovalType(v string) (interface{}, error) {
    result := MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE
    switch strings.ToUpper(v) {
        case "MANUAL":
            result = MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE
        case "AUTOMATIC":
            result = AUTOMATIC_DRIVERUPDATEPROFILEAPPROVALTYPE
        default:
            return 0, errors.New("Unknown DriverUpdateProfileApprovalType value: " + v)
    }
    return &result, nil
}
func SerializeDriverUpdateProfileApprovalType(values []DriverUpdateProfileApprovalType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
