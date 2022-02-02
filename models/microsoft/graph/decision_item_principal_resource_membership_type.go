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
    switch strings.ToUpper(v) {
        case "DIRECT":
            return DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE, nil
        case "INDIRECT":
            return INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE, nil
    }
    return 0, errors.New("Unknown DecisionItemPrincipalResourceMembershipType value: " + v)
}
func SerializeDecisionItemPrincipalResourceMembershipType(values []DecisionItemPrincipalResourceMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
