package models
// Indicates health state of the Windows management app.
type HealthState int

const (
    // Unknown state.
    UNKNOWN_HEALTHSTATE HealthState = iota
    // Healthy state.
    HEALTHY_HEALTHSTATE
    // Unhealthy state.
    UNHEALTHY_HEALTHSTATE
)

func (i HealthState) String() string {
    return []string{"unknown", "healthy", "unhealthy"}[i]
}
func ParseHealthState(v string) (any, error) {
    result := UNKNOWN_HEALTHSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_HEALTHSTATE
        case "healthy":
            result = HEALTHY_HEALTHSTATE
        case "unhealthy":
            result = UNHEALTHY_HEALTHSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHealthState(values []HealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HealthState) isMultiValue() bool {
    return false
}
