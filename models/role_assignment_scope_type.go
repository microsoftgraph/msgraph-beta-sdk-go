package models
import (
    "errors"
)
// Specifies the type of scope for a Role Assignment.
type RoleAssignmentScopeType int

const (
    // Allow assignments to the specified ResourceScopes.
    RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE RoleAssignmentScopeType = iota
    // Allow assignments to all Intune devices.
    ALLDEVICES_ROLEASSIGNMENTSCOPETYPE
    // Allow assignments to all Intune licensed users.
    ALLLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
    // Allow assignments to all Intune devices and licensed users.
    ALLDEVICESANDLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
)

func (i RoleAssignmentScopeType) String() string {
    return []string{"resourceScope", "allDevices", "allLicensedUsers", "allDevicesAndLicensedUsers"}[i]
}
func ParseRoleAssignmentScopeType(v string) (any, error) {
    result := RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE
    switch v {
        case "resourceScope":
            result = RESOURCESCOPE_ROLEASSIGNMENTSCOPETYPE
        case "allDevices":
            result = ALLDEVICES_ROLEASSIGNMENTSCOPETYPE
        case "allLicensedUsers":
            result = ALLLICENSEDUSERS_ROLEASSIGNMENTSCOPETYPE
        case "allDevicesAndLicensedUsers":
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
func (i RoleAssignmentScopeType) isMultiValue() bool {
    return false
}
