package networkaccess
import (
    "errors"
)
// 
type ForwardingRuleAction int

const (
    BYPASS_FORWARDINGRULEACTION ForwardingRuleAction = iota
    FORWARD_FORWARDINGRULEACTION
    UNKNOWNFUTUREVALUE_FORWARDINGRULEACTION
)

func (i ForwardingRuleAction) String() string {
    return []string{"bypass", "forward", "unknownFutureValue"}[i]
}
func ParseForwardingRuleAction(v string) (any, error) {
    result := BYPASS_FORWARDINGRULEACTION
    switch v {
        case "bypass":
            result = BYPASS_FORWARDINGRULEACTION
        case "forward":
            result = FORWARD_FORWARDINGRULEACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FORWARDINGRULEACTION
        default:
            return 0, errors.New("Unknown ForwardingRuleAction value: " + v)
    }
    return &result, nil
}
func SerializeForwardingRuleAction(values []ForwardingRuleAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ForwardingRuleAction) isMultiValue() bool {
    return false
}
