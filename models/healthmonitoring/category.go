package healthmonitoring
type Category int

const (
    UNKNOWN_CATEGORY Category = iota
    AUTHENTICATION_CATEGORY
    UNKNOWNFUTUREVALUE_CATEGORY
)

func (i Category) String() string {
    return []string{"unknown", "authentication", "unknownFutureValue"}[i]
}
func ParseCategory(v string) (any, error) {
    result := UNKNOWN_CATEGORY
    switch v {
        case "unknown":
            result = UNKNOWN_CATEGORY
        case "authentication":
            result = AUTHENTICATION_CATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCategory(values []Category) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Category) isMultiValue() bool {
    return false
}
