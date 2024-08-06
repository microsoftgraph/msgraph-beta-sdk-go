package models
type PerUserMfaState int

const (
    DISABLED_PERUSERMFASTATE PerUserMfaState = iota
    ENFORCED_PERUSERMFASTATE
    ENABLED_PERUSERMFASTATE
    UNKNOWNFUTUREVALUE_PERUSERMFASTATE
)

func (i PerUserMfaState) String() string {
    return []string{"disabled", "enforced", "enabled", "unknownFutureValue"}[i]
}
func ParsePerUserMfaState(v string) (any, error) {
    result := DISABLED_PERUSERMFASTATE
    switch v {
        case "disabled":
            result = DISABLED_PERUSERMFASTATE
        case "enforced":
            result = ENFORCED_PERUSERMFASTATE
        case "enabled":
            result = ENABLED_PERUSERMFASTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERUSERMFASTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePerUserMfaState(values []PerUserMfaState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PerUserMfaState) isMultiValue() bool {
    return false
}
