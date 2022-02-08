package graph
import (
    "strings"
    "errors"
)
// 
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
    result := RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE
    switch strings.ToUpper(v) {
        case "RESOURCESCOPE":
            result = RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE
        case "ALLDEVICES":
            result = ALLDEVICES_ROLEASSIGNMENTSCOPETYPE
        case "ALLLICENSEDUSERS":
            result = ALLLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
        case "ALLDEVICESANDLICENSEDUSERS":
            result = ALLDEVICESANDLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
        default:
            return 0, errors.New("Unknown RoleAssignmentScopeType value: " + v)
    }
    return &result, nil
}
func SerializeRoleAssignmentScopeType(values []RoleAssignmentScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
