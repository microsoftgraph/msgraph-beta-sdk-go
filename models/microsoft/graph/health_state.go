package graph
import (
    "strings"
    "errors"
)
// 
type HealthState int

const (
    UNKNOWN_HEALTHSTATE HealthState = iota
    HEALTHY_HEALTHSTATE
    UNHEALTHY_HEALTHSTATE
)

func (i HealthState) String() string {
    return []string{"UNKNOWN", "HEALTHY", "UNHEALTHY"}[i]
}
func ParseHealthState(v string) (interface{}, error) {
    result := UNKNOWN_HEALTHSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_HEALTHSTATE
        case "HEALTHY":
            result = HEALTHY_HEALTHSTATE
        case "UNHEALTHY":
            result = UNHEALTHY_HEALTHSTATE
        default:
            return 0, errors.New("Unknown HealthState value: " + v)
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
