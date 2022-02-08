package graph
import (
    "strings"
    "errors"
)
// 
type DecisionItemPrincipalResourceMembershipType int

const (
    DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE DecisionItemPrincipalResourceMembershipType = iota
    INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
    UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
)

func (i DecisionItemPrincipalResourceMembershipType) String() string {
    return []string{"DIRECT", "INDIRECT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDecisionItemPrincipalResourceMembershipType(v string) (interface{}, error) {
    result := DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
    switch strings.ToUpper(v) {
        case "DIRECT":
            result = DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
        case "INDIRECT":
            result = INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
        default:
            return 0, errors.New("Unknown DecisionItemPrincipalResourceMembershipType value: " + v)
    }
    return &result, nil
}
func SerializeDecisionItemPrincipalResourceMembershipType(values []DecisionItemPrincipalResourceMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
