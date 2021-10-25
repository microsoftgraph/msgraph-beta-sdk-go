package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "MANAGED":
            return MANAGED_MANAGEMENTSTATE, nil
        case "RETIREPENDING":
            return RETIREPENDING_MANAGEMENTSTATE, nil
        case "RETIREFAILED":
            return RETIREFAILED_MANAGEMENTSTATE, nil
        case "WIPEPENDING":
            return WIPEPENDING_MANAGEMENTSTATE, nil
        case "WIPEFAILED":
            return WIPEFAILED_MANAGEMENTSTATE, nil
        case "UNHEALTHY":
            return UNHEALTHY_MANAGEMENTSTATE, nil
        case "DELETEPENDING":
            return DELETEPENDING_MANAGEMENTSTATE, nil
        case "RETIREISSUED":
            return RETIREISSUED_MANAGEMENTSTATE, nil
        case "WIPEISSUED":
            return WIPEISSUED_MANAGEMENTSTATE, nil
        case "WIPECANCELED":
            return WIPECANCELED_MANAGEMENTSTATE, nil
        case "RETIRECANCELED":
            return RETIRECANCELED_MANAGEMENTSTATE, nil
        case "DISCOVERED":
            return DISCOVERED_MANAGEMENTSTATE, nil
    }
    return 0, errors.New("Unknown ManagementState value: " + v)
}
func SerializeManagementState(values []ManagementState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
