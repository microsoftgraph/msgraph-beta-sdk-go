package networkaccess
type ForwardingCategory int

const (
    DEFAULT_FORWARDINGCATEGORY ForwardingCategory = iota
    OPTIMIZED_FORWARDINGCATEGORY
    ALLOW_FORWARDINGCATEGORY
    UNKNOWNFUTUREVALUE_FORWARDINGCATEGORY
)

func (i ForwardingCategory) String() string {
    return []string{"default", "optimized", "allow", "unknownFutureValue"}[i]
}
func ParseForwardingCategory(v string) (any, error) {
    result := DEFAULT_FORWARDINGCATEGORY
    switch v {
        case "default":
            result = DEFAULT_FORWARDINGCATEGORY
        case "optimized":
            result = OPTIMIZED_FORWARDINGCATEGORY
        case "allow":
            result = ALLOW_FORWARDINGCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FORWARDINGCATEGORY
        default:
            return nil, nil
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
