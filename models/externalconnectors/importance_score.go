package externalconnectors
import (
    "errors"
)
// 
type ImportanceScore int

const (
    LOW_IMPORTANCESCORE ImportanceScore = iota
    MEDIUM_IMPORTANCESCORE
    HIGH_IMPORTANCESCORE
    VERYHIGH_IMPORTANCESCORE
    UNKNOWNFUTUREVALUE_IMPORTANCESCORE
)

func (i ImportanceScore) String() string {
    return []string{"low", "medium", "high", "veryHigh", "unknownFutureValue"}[i]
}
func ParseImportanceScore(v string) (any, error) {
    result := LOW_IMPORTANCESCORE
    switch v {
        case "low":
            result = LOW_IMPORTANCESCORE
        case "medium":
            result = MEDIUM_IMPORTANCESCORE
        case "high":
            result = HIGH_IMPORTANCESCORE
        case "veryHigh":
            result = VERYHIGH_IMPORTANCESCORE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IMPORTANCESCORE
        default:
            return 0, errors.New("Unknown ImportanceScore value: " + v)
    }
    return &result, nil
}
func SerializeImportanceScore(values []ImportanceScore) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ImportanceScore) isMultiValue() bool {
    return false
}
