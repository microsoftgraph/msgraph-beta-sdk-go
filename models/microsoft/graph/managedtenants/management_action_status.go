package managedtenants
import (
    "strings"
    "errors"
)
// 
type ManagementActionStatus int

const (
    TOADDRESS_MANAGEMENTACTIONSTATUS ManagementActionStatus = iota
    COMPLETED_MANAGEMENTACTIONSTATUS
    ERROR_MANAGEMENTACTIONSTATUS
    TIMEOUT_MANAGEMENTACTIONSTATUS
    INPROGRESS_MANAGEMENTACTIONSTATUS
    PLANNED_MANAGEMENTACTIONSTATUS
    RESOLVEDBY3RDPARTY_MANAGEMENTACTIONSTATUS
    RESOLVEDTHROUGHALTERNATEMITIGATION_MANAGEMENTACTIONSTATUS
    RISKACCEPTED_MANAGEMENTACTIONSTATUS
    UNKNOWNFUTUREVALUE_MANAGEMENTACTIONSTATUS
)

func (i ManagementActionStatus) String() string {
    return []string{"TOADDRESS", "COMPLETED", "ERROR", "TIMEOUT", "INPROGRESS", "PLANNED", "RESOLVEDBY3RDPARTY", "RESOLVEDTHROUGHALTERNATEMITIGATION", "RISKACCEPTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseManagementActionStatus(v string) (interface{}, error) {
    result := TOADDRESS_MANAGEMENTACTIONSTATUS
    switch strings.ToUpper(v) {
        case "TOADDRESS":
            result = TOADDRESS_MANAGEMENTACTIONSTATUS
        case "COMPLETED":
            result = COMPLETED_MANAGEMENTACTIONSTATUS
        case "ERROR":
            result = ERROR_MANAGEMENTACTIONSTATUS
        case "TIMEOUT":
            result = TIMEOUT_MANAGEMENTACTIONSTATUS
        case "INPROGRESS":
            result = INPROGRESS_MANAGEMENTACTIONSTATUS
        case "PLANNED":
            result = PLANNED_MANAGEMENTACTIONSTATUS
        case "RESOLVEDBY3RDPARTY":
            result = RESOLVEDBY3RDPARTY_MANAGEMENTACTIONSTATUS
        case "RESOLVEDTHROUGHALTERNATEMITIGATION":
            result = RESOLVEDTHROUGHALTERNATEMITIGATION_MANAGEMENTACTIONSTATUS
        case "RISKACCEPTED":
            result = RISKACCEPTED_MANAGEMENTACTIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTACTIONSTATUS
        default:
            return 0, errors.New("Unknown ManagementActionStatus value: " + v)
    }
    return &result, nil
}
func SerializeManagementActionStatus(values []ManagementActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
