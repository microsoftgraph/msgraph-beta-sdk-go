package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type RecommendationCategory int

const (
    USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY RecommendationCategory = iota
    SECURITY_RECOMMENDATIONCATEGORY
    UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY
    PRODUCTIVITY_RECOMMENDATIONCATEGORY
    HEALTH_RECOMMENDATIONCATEGORY
    CONFIGURATION_RECOMMENDATIONCATEGORY
)

func (i RecommendationCategory) String() string {
    return []string{"usageAndCompliance", "security", "unknownFutureValue", "productivity", "health", "configuration"}[i]
}
func ParseRecommendationCategory(v string) (interface{}, error) {
    result := USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY
    switch v {
        case "usageAndCompliance":
            result = USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY
        case "security":
            result = SECURITY_RECOMMENDATIONCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY
        case "productivity":
            result = PRODUCTIVITY_RECOMMENDATIONCATEGORY
        case "health":
            result = HEALTH_RECOMMENDATIONCATEGORY
        case "configuration":
            result = CONFIGURATION_RECOMMENDATIONCATEGORY
        default:
            return 0, errors.New("Unknown RecommendationCategory value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationCategory(values []RecommendationCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
