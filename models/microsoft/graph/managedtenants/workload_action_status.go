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
    switch strings.ToUpper(v) {
        case "TOADDRESS":
            return TOADDRESS_WORKLOADACTIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_WORKLOADACTIONSTATUS, nil
        case "ERROR":
            return ERROR_WORKLOADACTIONSTATUS, nil
        case "TIMEOUT":
            return TIMEOUT_WORKLOADACTIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_WORKLOADACTIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WORKLOADACTIONSTATUS, nil
    }
    return 0, errors.New("Unknown WorkloadActionStatus value: " + v)
}
func SerializeWorkloadActionStatus(values []WorkloadActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
