package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the tenantRelationship singleton.
type DelegatedAdminAccessAssignmentStatus int

const (
    PENDING_DELEGATEDADMINACCESSASSIGNMENTSTATUS DelegatedAdminAccessAssignmentStatus = iota
    ACTIVE_DELEGATEDADMINACCESSASSIGNMENTSTATUS
    DELETING_DELEGATEDADMINACCESSASSIGNMENTSTATUS
    DELETED_DELEGATEDADMINACCESSASSIGNMENTSTATUS
    ERROR_DELEGATEDADMINACCESSASSIGNMENTSTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDADMINACCESSASSIGNMENTSTATUS
)

func (i DelegatedAdminAccessAssignmentStatus) String() string {
    return []string{"PENDING", "ACTIVE", "DELETING", "DELETED", "ERROR", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminAccessAssignmentStatus(v string) (interface{}, error) {
    result := PENDING_DELEGATEDADMINACCESSASSIGNMENTSTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_DELEGATEDADMINACCESSASSIGNMENTSTATUS
        case "ACTIVE":
            result = ACTIVE_DELEGATEDADMINACCESSASSIGNMENTSTATUS
        case "DELETING":
            result = DELETING_DELEGATEDADMINACCESSASSIGNMENTSTATUS
        case "DELETED":
            result = DELETED_DELEGATEDADMINACCESSASSIGNMENTSTATUS
        case "ERROR":
            result = ERROR_DELEGATEDADMINACCESSASSIGNMENTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINACCESSASSIGNMENTSTATUS
        default:
            return 0, errors.New("Unknown DelegatedAdminAccessAssignmentStatus value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminAccessAssignmentStatus(values []DelegatedAdminAccessAssignmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
