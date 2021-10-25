package graph
import (
    "strings"
    "errors"
)
type RoleAssignmentScopeType int

const (
    RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE RoleAssignmentScopeType = iota
    ALLDEVICES_ROLEASSIGNMENTSCOPETYPE
    ALLLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
    ALLDEVICESANDLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
)

func (i RoleAssignmentScopeType) String() string {
    return []string{"RESOURCESCOPE", "ALLDEVICES", "ALLLICENSEDUSERS", "ALLDEVICESANDLICENSEDUSERS"}[i]
}
func ParseRoleAssignmentScopeType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "RESOURCESCOPE":
            return RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE, nil
        case "ALLDEVICES":
            return ALLDEVICES_ROLEASSIGNMENTSCOPETYPE, nil
        case "ALLLICENSEDUSERS":
            return ALLLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE, nil
        case "ALLDEVICESANDLICENSEDUSERS":
            return ALLDEVICESANDLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE, nil
    }
    return 0, errors.New("Unknown RoleAssignmentScopeType value: " + v)
}
func SerializeRoleAssignmentScopeType(values []RoleAssignmentScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
