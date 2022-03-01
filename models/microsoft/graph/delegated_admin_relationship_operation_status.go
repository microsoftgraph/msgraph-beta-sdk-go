package graph
import (
    "strings"
    "errors"
)
// 
type DelegatedAdminRelationshipOperationStatus int

const (
    NOTSTARTED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS DelegatedAdminRelationshipOperationStatus = iota
    RUNNING_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    COMPLETE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    FAILED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
)

func (i DelegatedAdminRelationshipOperationStatus) String() string {
    return []string{"NOTSTARTED", "RUNNING", "COMPLETE", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminRelationshipOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "RUNNING":
            result = RUNNING_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "COMPLETE":
            result = COMPLETE_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "FAILED":
            result = FAILED_DELEGATEDADMINRELATIONSHIPOPERATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
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
