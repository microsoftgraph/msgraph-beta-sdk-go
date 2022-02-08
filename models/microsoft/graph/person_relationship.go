package graph
import (
    "strings"
    "errors"
)
// 
type PersonRelationship int

const (
    MANAGER_PERSONRELATIONSHIP PersonRelationship = iota
    COLLEAGUE_PERSONRELATIONSHIP
    DIRECTREPORT_PERSONRELATIONSHIP
    DOTLINEREPORT_PERSONRELATIONSHIP
    ASSISTANT_PERSONRELATIONSHIP
    DOTLINEMANAGER_PERSONRELATIONSHIP
    ALTERNATECONTACT_PERSONRELATIONSHIP
    FRIEND_PERSONRELATIONSHIP
    SPOUSE_PERSONRELATIONSHIP
    SIBLING_PERSONRELATIONSHIP
    CHILD_PERSONRELATIONSHIP
    PARENT_PERSONRELATIONSHIP
    SPONSOR_PERSONRELATIONSHIP
    EMERGENCYCONTACT_PERSONRELATIONSHIP
    OTHER_PERSONRELATIONSHIP
    UNKNOWNFUTUREVALUE_PERSONRELATIONSHIP
)

func (i PersonRelationship) String() string {
    return []string{"MANAGER", "COLLEAGUE", "DIRECTREPORT", "DOTLINEREPORT", "ASSISTANT", "DOTLINEMANAGER", "ALTERNATECONTACT", "FRIEND", "SPOUSE", "SIBLING", "CHILD", "PARENT", "SPONSOR", "EMERGENCYCONTACT", "OTHER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePersonRelationship(v string) (interface{}, error) {
    result := MANAGER_PERSONRELATIONSHIP
    switch strings.ToUpper(v) {
        case "MANAGER":
            result = MANAGER_PERSONRELATIONSHIP
        case "COLLEAGUE":
            result = COLLEAGUE_PERSONRELATIONSHIP
        case "DIRECTREPORT":
            result = DIRECTREPORT_PERSONRELATIONSHIP
        case "DOTLINEREPORT":
            result = DOTLINEREPORT_PERSONRELATIONSHIP
        case "ASSISTANT":
            result = ASSISTANT_PERSONRELATIONSHIP
        case "DOTLINEMANAGER":
            result = DOTLINEMANAGER_PERSONRELATIONSHIP
        case "ALTERNATECONTACT":
            result = ALTERNATECONTACT_PERSONRELATIONSHIP
        case "FRIEND":
            result = FRIEND_PERSONRELATIONSHIP
        case "SPOUSE":
            result = SPOUSE_PERSONRELATIONSHIP
        case "SIBLING":
            result = SIBLING_PERSONRELATIONSHIP
        case "CHILD":
            result = CHILD_PERSONRELATIONSHIP
        case "PARENT":
            result = PARENT_PERSONRELATIONSHIP
        case "SPONSOR":
            result = SPONSOR_PERSONRELATIONSHIP
        case "EMERGENCYCONTACT":
            result = EMERGENCYCONTACT_PERSONRELATIONSHIP
        case "OTHER":
            result = OTHER_PERSONRELATIONSHIP
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PERSONRELATIONSHIP
        default:
            return 0, errors.New("Unknown PersonRelationship value: " + v)
    }
    return &result, nil
}
func SerializePersonRelationship(values []PersonRelationship) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
