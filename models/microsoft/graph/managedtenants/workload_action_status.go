package managedtenants
import (
    "strings"
    "errors"
)
// 
type WorkloadActionStatus int

const (
    TOADDRESS_WORKLOADACTIONSTATUS WorkloadActionStatus = iota
    COMPLETED_WORKLOADACTIONSTATUS
    ERROR_WORKLOADACTIONSTATUS
    TIMEOUT_WORKLOADACTIONSTATUS
    INPROGRESS_WORKLOADACTIONSTATUS
    UNKNOWNFUTUREVALUE_WORKLOADACTIONSTATUS
)

func (i WorkloadActionStatus) String() string {
    return []string{"TOADDRESS", "COMPLETED", "ERROR", "TIMEOUT", "INPROGRESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWorkloadActionStatus(v string) (interface{}, error) {
    result := TOADDRESS_WORKLOADACTIONSTATUS
    switch strings.ToUpper(v) {
        case "TOADDRESS":
            result = TOADDRESS_WORKLOADACTIONSTATUS
        case "COMPLETED":
            result = COMPLETED_WORKLOADACTIONSTATUS
        case "ERROR":
            result = ERROR_WORKLOADACTIONSTATUS
        case "TIMEOUT":
            result = TIMEOUT_WORKLOADACTIONSTATUS
        case "INPROGRESS":
            result = INPROGRESS_WORKLOADACTIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WORKLOADACTIONSTATUS
        default:
            return 0, errors.New("Unknown WorkloadActionStatus value: " + v)
    }
    return &result, nil
}
func SerializeWorkloadActionStatus(values []WorkloadActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
