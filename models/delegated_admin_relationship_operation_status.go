package models
import (
    "errors"
)
// Provides operations to manage the tenantRelationship singleton.
type DelegatedAdminRelationshipOperationStatus int

const (
    NOTSTARTED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS DelegatedAdminRelationshipOperationStatus = iota
    RUNNING_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    COMPLETE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    FAILED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
)

func (i DelegatedAdminRelationshipOperationStatus) String() string {
    return []string{"notStarted", "running", "complete", "failed", "unknownFutureValue"}[i]
}
func ParseDelegatedAdminRelationshipOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "running":
            result = RUNNING_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "complete":
            result = COMPLETE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "failed":
            result = FAILED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipOperationStatus(values []DelegatedAdminRelationshipOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
