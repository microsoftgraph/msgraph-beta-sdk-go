package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the tenantRelationship singleton.
type DelegatedAdminRelationshipStatus int

const (
    ACTIVATING_DELEGATEDADMINRELATIONSHIPSTATUS DelegatedAdminRelationshipStatus = iota
    ACTIVE_DELEGATEDADMINRELATIONSHIPSTATUS
    APPROVALPENDING_DELEGATEDADMINRELATIONSHIPSTATUS
    APPROVED_DELEGATEDADMINRELATIONSHIPSTATUS
    CREATED_DELEGATEDADMINRELATIONSHIPSTATUS
    EXPIRED_DELEGATEDADMINRELATIONSHIPSTATUS
    EXPIRING_DELEGATEDADMINRELATIONSHIPSTATUS
    TERMINATED_DELEGATEDADMINRELATIONSHIPSTATUS
    TERMINATING_DELEGATEDADMINRELATIONSHIPSTATUS
    TERMINATIONREQUESTED_DELEGATEDADMINRELATIONSHIPSTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPSTATUS
)

func (i DelegatedAdminRelationshipStatus) String() string {
    return []string{"ACTIVATING", "ACTIVE", "APPROVALPENDING", "APPROVED", "CREATED", "EXPIRED", "EXPIRING", "TERMINATED", "TERMINATING", "TERMINATIONREQUESTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminRelationshipStatus(v string) (interface{}, error) {
    result := ACTIVATING_DELEGATEDADMINRELATIONSHIPSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVATING":
            result = ACTIVATING_DELEGATEDADMINRELATIONSHIPSTATUS
        case "ACTIVE":
            result = ACTIVE_DELEGATEDADMINRELATIONSHIPSTATUS
        case "APPROVALPENDING":
            result = APPROVALPENDING_DELEGATEDADMINRELATIONSHIPSTATUS
        case "APPROVED":
            result = APPROVED_DELEGATEDADMINRELATIONSHIPSTATUS
        case "CREATED":
            result = CREATED_DELEGATEDADMINRELATIONSHIPSTATUS
        case "EXPIRED":
            result = EXPIRED_DELEGATEDADMINRELATIONSHIPSTATUS
        case "EXPIRING":
            result = EXPIRING_DELEGATEDADMINRELATIONSHIPSTATUS
        case "TERMINATED":
            result = TERMINATED_DELEGATEDADMINRELATIONSHIPSTATUS
        case "TERMINATING":
            result = TERMINATING_DELEGATEDADMINRELATIONSHIPSTATUS
        case "TERMINATIONREQUESTED":
            result = TERMINATIONREQUESTED_DELEGATEDADMINRELATIONSHIPSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPSTATUS
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipStatus value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipStatus(values []DelegatedAdminRelationshipStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
