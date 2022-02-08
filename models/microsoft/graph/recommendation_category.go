package graph
import (
    "strings"
    "errors"
)
// 
type RecommendationCategory int

const (
    USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY RecommendationCategory = iota
    SECURITY_RECOMMENDATIONCATEGORY
    PRODUCTIVITY_RECOMMENDATIONCATEGORY
    HEALTH_RECOMMENDATIONCATEGORY
    CONFIGURATION_RECOMMENDATIONCATEGORY
    UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY
)

func (i RecommendationCategory) String() string {
    return []string{"USAGEANDCOMPLIANCE", "SECURITY", "PRODUCTIVITY", "HEALTH", "CONFIGURATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRecommendationCategory(v string) (interface{}, error) {
    result := USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY
    switch strings.ToUpper(v) {
        case "USAGEANDCOMPLIANCE":
            result = USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY
        case "SECURITY":
            result = SECURITY_RECOMMENDATIONCATEGORY
        case "PRODUCTIVITY":
            result = PRODUCTIVITY_RECOMMENDATIONCATEGORY
        case "HEALTH":
            result = HEALTH_RECOMMENDATIONCATEGORY
        case "CONFIGURATION":
            result = CONFIGURATION_RECOMMENDATIONCATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY
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
