package networkaccess
type Status int

const (
    ENABLED_STATUS Status = iota
    DISABLED_STATUS
    UNKNOWNFUTUREVALUE_STATUS
)

func (i Status) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseStatus(v string) (any, error) {
    result := ENABLED_STATUS
    switch v {
        case "enabled":
            result = ENABLED_STATUS
        case "disabled":
            result = DISABLED_STATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStatus(values []Status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Status) isMultiValue() bool {
    return false
}
