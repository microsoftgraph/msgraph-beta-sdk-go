package models
import (
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
    return []string{"created", "pending", "complete", "failed", "unknownFutureValue"}[i]
}
func ParseDelegatedAdminRelationshipRequestStatus(v string) (interface{}, error) {
    result := CREATED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
    switch v {
        case "created":
            result = CREATED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "pending":
            result = PENDING_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "complete":
            result = COMPLETE_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "failed":
            result = FAILED_DELEGATEDADMINRELATIONSHIPREQUESTSTATUS
        case "unknownFutureValue":
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
