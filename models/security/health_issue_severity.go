package security
import (
    "errors"
)
type HealthIssueSeverity int

const (
    LOW_HEALTHISSUESEVERITY HealthIssueSeverity = iota
    MEDIUM_HEALTHISSUESEVERITY
    HIGH_HEALTHISSUESEVERITY
    UNKNOWNFUTUREVALUE_HEALTHISSUESEVERITY
)

func (i HealthIssueSeverity) String() string {
    return []string{"low", "medium", "high", "unknownFutureValue"}[i]
}
func ParseHealthIssueSeverity(v string) (any, error) {
    result := LOW_HEALTHISSUESEVERITY
    switch v {
        case "low":
            result = LOW_HEALTHISSUESEVERITY
        case "medium":
            result = MEDIUM_HEALTHISSUESEVERITY
        case "high":
            result = HIGH_HEALTHISSUESEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_HEALTHISSUESEVERITY
        default:
            return 0, errors.New("Unknown HealthIssueSeverity value: " + v)
    }
    return &result, nil
}
func SerializeHealthIssueSeverity(values []HealthIssueSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HealthIssueSeverity) isMultiValue() bool {
    return false
}
