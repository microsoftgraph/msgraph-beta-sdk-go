package graph
import (
    "strings"
    "errors"
)
// 
type ManagementState int

const (
    MANAGED_MANAGEMENTSTATE ManagementState = iota
    RETIREPENDING_MANAGEMENTSTATE
    RETIREFAILED_MANAGEMENTSTATE
    WIPEPENDING_MANAGEMENTSTATE
    WIPEFAILED_MANAGEMENTSTATE
    UNHEALTHY_MANAGEMENTSTATE
    DELETEPENDING_MANAGEMENTSTATE
    RETIREISSUED_MANAGEMENTSTATE
    WIPEISSUED_MANAGEMENTSTATE
    WIPECANCELED_MANAGEMENTSTATE
    RETIRECANCELED_MANAGEMENTSTATE
    DISCOVERED_MANAGEMENTSTATE
)

func (i ManagementState) String() string {
    return []string{"MANAGED", "RETIREPENDING", "RETIREFAILED", "WIPEPENDING", "WIPEFAILED", "UNHEALTHY", "DELETEPENDING", "RETIREISSUED", "WIPEISSUED", "WIPECANCELED", "RETIRECANCELED", "DISCOVERED"}[i]
}
func ParseManagementState(v string) (interface{}, error) {
    result := MANAGED_MANAGEMENTSTATE
    switch strings.ToUpper(v) {
        case "MANAGED":
            result = MANAGED_MANAGEMENTSTATE
        case "RETIREPENDING":
            result = RETIREPENDING_MANAGEMENTSTATE
        case "RETIREFAILED":
            result = RETIREFAILED_MANAGEMENTSTATE
        case "WIPEPENDING":
            result = WIPEPENDING_MANAGEMENTSTATE
        case "WIPEFAILED":
            result = WIPEFAILED_MANAGEMENTSTATE
        case "UNHEALTHY":
            result = UNHEALTHY_MANAGEMENTSTATE
        case "DELETEPENDING":
            result = DELETEPENDING_MANAGEMENTSTATE
        case "RETIREISSUED":
            result = RETIREISSUED_MANAGEMENTSTATE
        case "WIPEISSUED":
            result = WIPEISSUED_MANAGEMENTSTATE
        case "WIPECANCELED":
            result = WIPECANCELED_MANAGEMENTSTATE
        case "RETIRECANCELED":
            result = RETIRECANCELED_MANAGEMENTSTATE
        case "DISCOVERED":
            result = DISCOVERED_MANAGEMENTSTATE
        default:
            return 0, errors.New("Unknown ManagementState value: " + v)
    }
    return &result, nil
}
func SerializeManagementState(values []ManagementState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
