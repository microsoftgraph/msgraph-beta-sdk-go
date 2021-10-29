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
    switch strings.ToUpper(v) {
        case "MANUAL":
            return MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE, nil
        case "AUTOMATIC":
            return AUTOMATIC_DRIVERUPDATEPROFILEAPPROVALTYPE, nil
    }
    return 0, errors.New("Unknown DriverUpdateProfileApprovalType value: " + v)
}
func SerializeDriverUpdateProfileApprovalType(values []DriverUpdateProfileApprovalType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
