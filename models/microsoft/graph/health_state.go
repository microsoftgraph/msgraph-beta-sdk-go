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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_HEALTHSTATE, nil
        case "HEALTHY":
            return HEALTHY_HEALTHSTATE, nil
        case "UNHEALTHY":
            return UNHEALTHY_HEALTHSTATE, nil
    }
    return 0, errors.New("Unknown HealthState value: " + v)
}
func SerializeHealthState(values []HealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
