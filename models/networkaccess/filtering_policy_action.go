package networkaccess
import (
    "errors"
)
// 
type FilteringPolicyAction int

const (
    BLOCK_FILTERINGPOLICYACTION FilteringPolicyAction = iota
    ALLOW_FILTERINGPOLICYACTION
    UNKNOWNFUTUREVALUE_FILTERINGPOLICYACTION
)

func (i FilteringPolicyAction) String() string {
    return []string{"block", "allow", "unknownFutureValue"}[i]
}
func ParseFilteringPolicyAction(v string) (any, error) {
    result := BLOCK_FILTERINGPOLICYACTION
    switch v {
        case "block":
            result = BLOCK_FILTERINGPOLICYACTION
        case "allow":
            result = ALLOW_FILTERINGPOLICYACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FILTERINGPOLICYACTION
        default:
            return 0, errors.New("Unknown FilteringPolicyAction value: " + v)
    }
    return &result, nil
}
func SerializeFilteringPolicyAction(values []FilteringPolicyAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FilteringPolicyAction) isMultiValue() bool {
    return false
}
