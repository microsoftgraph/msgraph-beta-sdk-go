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
    switch strings.ToUpper(v) {
        case "LOW":
            return LOW_RECOMMENDATIONPRIORITY, nil
        case "MEDIUM":
            return MEDIUM_RECOMMENDATIONPRIORITY, nil
        case "HIGH":
            return HIGH_RECOMMENDATIONPRIORITY, nil
    }
    return 0, errors.New("Unknown RecommendationPriority value: " + v)
}
func SerializeRecommendationPriority(values []RecommendationPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
