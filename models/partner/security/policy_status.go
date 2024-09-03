package security
type PolicyStatus int

const (
    ENABLED_POLICYSTATUS PolicyStatus = iota
    DISABLED_POLICYSTATUS
    UNKNOWNFUTUREVALUE_POLICYSTATUS
)

func (i PolicyStatus) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParsePolicyStatus(v string) (any, error) {
    result := ENABLED_POLICYSTATUS
    switch v {
        case "enabled":
            result = ENABLED_POLICYSTATUS
        case "disabled":
            result = DISABLED_POLICYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_POLICYSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePolicyStatus(values []PolicyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PolicyStatus) isMultiValue() bool {
    return false
}
