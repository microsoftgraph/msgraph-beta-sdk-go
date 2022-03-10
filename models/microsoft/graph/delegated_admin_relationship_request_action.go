package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the tenantRelationship singleton.
type DelegatedAdminRelationshipRequestAction int

const (
    LOCKFORAPPROVAL_DELEGATEDADMINRELATIONSHIPREQUESTACTION DelegatedAdminRelationshipRequestAction = iota
    APPROVE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
    TERMINATE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
)

func (i DelegatedAdminRelationshipRequestAction) String() string {
    return []string{"LOCKFORAPPROVAL", "APPROVE", "TERMINATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDelegatedAdminRelationshipRequestAction(v string) (interface{}, error) {
    result := LOCKFORAPPROVAL_DELEGATEDADMINRELATIONSHIPREQUESTACTION
    switch strings.ToUpper(v) {
        case "LOCKFORAPPROVAL":
            result = LOCKFORAPPROVAL_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        case "APPROVE":
            result = APPROVE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        case "TERMINATE":
            result = TERMINATE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipRequestAction value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipRequestAction(values []DelegatedAdminRelationshipRequestAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
