package models
import (
    "errors"
    "strings"
)
// 
type DecisionItemPrincipalResourceMembershipType int

const (
    DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE DecisionItemPrincipalResourceMembershipType = iota
    INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
    UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
)

func (i DecisionItemPrincipalResourceMembershipType) String() string {
    var values []string
    for p := DecisionItemPrincipalResourceMembershipType(1); p <= UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"direct", "indirect", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDecisionItemPrincipalResourceMembershipType(v string) (any, error) {
    var result DecisionItemPrincipalResourceMembershipType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "direct":
                result |= DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
            case "indirect":
                result |= INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
            default:
                return 0, errors.New("Unknown DecisionItemPrincipalResourceMembershipType value: " + v)
        }
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
func (i DecisionItemPrincipalResourceMembershipType) isMultiValue() bool {
    return true
}
