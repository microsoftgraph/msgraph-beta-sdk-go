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
    switch strings.ToUpper(v) {
        case "MANAGER":
            return MANAGER_PERSONRELATIONSHIP, nil
        case "COLLEAGUE":
            return COLLEAGUE_PERSONRELATIONSHIP, nil
        case "DIRECTREPORT":
            return DIRECTREPORT_PERSONRELATIONSHIP, nil
        case "DOTLINEREPORT":
            return DOTLINEREPORT_PERSONRELATIONSHIP, nil
        case "ASSISTANT":
            return ASSISTANT_PERSONRELATIONSHIP, nil
        case "DOTLINEMANAGER":
            return DOTLINEMANAGER_PERSONRELATIONSHIP, nil
        case "ALTERNATECONTACT":
            return ALTERNATECONTACT_PERSONRELATIONSHIP, nil
        case "FRIEND":
            return FRIEND_PERSONRELATIONSHIP, nil
        case "SPOUSE":
            return SPOUSE_PERSONRELATIONSHIP, nil
        case "SIBLING":
            return SIBLING_PERSONRELATIONSHIP, nil
        case "CHILD":
            return CHILD_PERSONRELATIONSHIP, nil
        case "PARENT":
            return PARENT_PERSONRELATIONSHIP, nil
        case "SPONSOR":
            return SPONSOR_PERSONRELATIONSHIP, nil
        case "EMERGENCYCONTACT":
            return EMERGENCYCONTACT_PERSONRELATIONSHIP, nil
        case "OTHER":
            return OTHER_PERSONRELATIONSHIP, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PERSONRELATIONSHIP, nil
    }
    return 0, errors.New("Unknown PersonRelationship value: " + v)
}
func SerializePersonRelationship(values []PersonRelationship) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
