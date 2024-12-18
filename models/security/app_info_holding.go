package security
type AppInfoHolding int

const (
    PRIVATE_APPINFOHOLDING AppInfoHolding = iota
    PUBLIC_APPINFOHOLDING
    UNKNOWN_APPINFOHOLDING
    UNKNOWNFUTUREVALUE_APPINFOHOLDING
)

func (i AppInfoHolding) String() string {
    return []string{"private", "public", "unknown", "unknownFutureValue"}[i]
}
func ParseAppInfoHolding(v string) (any, error) {
    result := PRIVATE_APPINFOHOLDING
    switch v {
        case "private":
            result = PRIVATE_APPINFOHOLDING
        case "public":
            result = PUBLIC_APPINFOHOLDING
        case "unknown":
            result = UNKNOWN_APPINFOHOLDING
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFOHOLDING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoHolding(values []AppInfoHolding) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoHolding) isMultiValue() bool {
    return false
}
