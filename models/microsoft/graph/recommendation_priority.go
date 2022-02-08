package graph
import (
    "strings"
    "errors"
)
// 
type RecommendationPriority int

const (
    LOW_RECOMMENDATIONPRIORITY RecommendationPriority = iota
    MEDIUM_RECOMMENDATIONPRIORITY
    HIGH_RECOMMENDATIONPRIORITY
)

func (i RecommendationPriority) String() string {
    return []string{"LOW", "MEDIUM", "HIGH"}[i]
}
func ParseRecommendationPriority(v string) (interface{}, error) {
    result := LOW_RECOMMENDATIONPRIORITY
    switch strings.ToUpper(v) {
        case "LOW":
            result = LOW_RECOMMENDATIONPRIORITY
        case "MEDIUM":
            result = MEDIUM_RECOMMENDATIONPRIORITY
        case "HIGH":
            result = HIGH_RECOMMENDATIONPRIORITY
        default:
            return 0, errors.New("Unknown RecommendationPriority value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationPriority(values []RecommendationPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
