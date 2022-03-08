package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the tenantRelationship singleton.
type DelegatedAdminRelationshipRequestStatus int

const (
    CREATED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS DelegatedAdminRelationshipRequestStatus = iota
    PENDING_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
    COMPLETE_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
    FAILED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
)

func (i DelegatedAdminRelationshipRequestStatus) String() string {
    return []string{"CREATED", "PENDING", "COMPLETE", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminRelationshipRequestStatus(v string) (interface{}, error) {
    result := CREATED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
    switch strings.ToUpper(v) {
        case "CREATED":
            result = CREATED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "PENDING":
            result = PENDING_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "COMPLETE":
            result = COMPLETE_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "FAILED":
            result = FAILED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipRequestStatus value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipRequestStatus(values []DelegatedAdminRelationshipRequestStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
