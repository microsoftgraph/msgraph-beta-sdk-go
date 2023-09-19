package networkaccess
import (
    "errors"
)
// 
type ForwardingCategory int

const (
    DEFAULTESCAPED_FORWARDINGCATEGORY ForwardingCategory = iota
    OPTIMIZED_FORWARDINGCATEGORY
    ALLOW_FORWARDINGCATEGORY
    UNKNOWNFUTUREVALUE_FORWARDINGCATEGORY
)

func (i ForwardingCategory) String() string {
    return []string{"default", "optimized", "allow", "unknownFutureValue"}[i]
}
func ParseForwardingCategory(v string) (any, error) {
    result := DEFAULTESCAPED_FORWARDINGCATEGORY
    switch v {
        case "default":
            result = DEFAULTESCAPED_FORWARDINGCATEGORY
        case "optimized":
            result = OPTIMIZED_FORWARDINGCATEGORY
        case "allow":
            result = ALLOW_FORWARDINGCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FORWARDINGCATEGORY
        default:
            return 0, errors.New("Unknown ForwardingCategory value: " + v)
    }
    return &result, nil
}
func SerializeForwardingCategory(values []ForwardingCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ForwardingCategory) isMultiValue() bool {
    return false
}
