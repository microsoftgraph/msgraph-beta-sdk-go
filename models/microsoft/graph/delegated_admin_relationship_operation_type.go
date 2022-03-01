package graph
import (
    "strings"
    "errors"
)
// 
type DelegatedAdminRelationshipOperationType int

const (
    DELEGATEDADMINACCESSASSIGNMENTUPDATE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE DelegatedAdminRelationshipOperationType = iota
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
)

func (i DelegatedAdminRelationshipOperationType) String() string {
    return []string{"DELEGATEDADMINACCESSASSIGNMENTUPDATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminRelationshipOperationType(v string) (interface{}, error) {
    result := DELEGATEDADMINACCESSASSIGNMENTUPDATE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
    switch strings.ToUpper(v) {
        case "DELEGATEDADMINACCESSASSIGNMENTUPDATE":
            result = DELEGATEDADMINACCESSASSIGNMENTUPDATE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipOperationType value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipOperationType(values []DelegatedAdminRelationshipOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
