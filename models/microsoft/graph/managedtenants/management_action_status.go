package managedtenants
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "TOADDRESS":
            return TOADDRESS_MANAGEMENTACTIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_MANAGEMENTACTIONSTATUS, nil
        case "ERROR":
            return ERROR_MANAGEMENTACTIONSTATUS, nil
        case "TIMEOUT":
            return TIMEOUT_MANAGEMENTACTIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_MANAGEMENTACTIONSTATUS, nil
        case "PLANNED":
            return PLANNED_MANAGEMENTACTIONSTATUS, nil
        case "RESOLVEDBY3RDPARTY":
            return RESOLVEDBY3RDPARTY_MANAGEMENTACTIONSTATUS, nil
        case "RESOLVEDTHROUGHALTERNATEMITIGATION":
            return RESOLVEDTHROUGHALTERNATEMITIGATION_MANAGEMENTACTIONSTATUS, nil
        case "RISKACCEPTED":
            return RISKACCEPTED_MANAGEMENTACTIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MANAGEMENTACTIONSTATUS, nil
    }
    return 0, errors.New("Unknown ManagementActionStatus value: " + v)
}
func SerializeManagementActionStatus(values []ManagementActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
