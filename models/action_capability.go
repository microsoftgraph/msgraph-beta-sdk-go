package models
type ActionCapability int

const (
    ENABLED_ACTIONCAPABILITY ActionCapability = iota
    DISABLED_ACTIONCAPABILITY
    UNKNOWNFUTUREVALUE_ACTIONCAPABILITY
)

func (i ActionCapability) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseActionCapability(v string) (any, error) {
    result := ENABLED_ACTIONCAPABILITY
    switch v {
        case "enabled":
            result = ENABLED_ACTIONCAPABILITY
        case "disabled":
            result = DISABLED_ACTIONCAPABILITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACTIONCAPABILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeActionCapability(values []ActionCapability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionCapability) isMultiValue() bool {
    return false
}
