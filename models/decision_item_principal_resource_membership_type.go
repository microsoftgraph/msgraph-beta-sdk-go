package models
import (
    "errors"
    "math"
    "strings"
)
// 
type DecisionItemPrincipalResourceMembershipType int

const (
    DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE = 1
    INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE = 2
    UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE = 4
)

func (i DecisionItemPrincipalResourceMembershipType) String() string {
    var values []string
    options := []string{"direct", "indirect", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := DecisionItemPrincipalResourceMembershipType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
