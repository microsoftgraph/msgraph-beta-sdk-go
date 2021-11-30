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
    switch strings.ToUpper(v) {
        case "USAGEANDCOMPLIANCE":
            return USAGEANDCOMPLIANCE_RECOMMENDATIONCATEGORY, nil
        case "SECURITY":
            return SECURITY_RECOMMENDATIONCATEGORY, nil
        case "PRODUCTIVITY":
            return PRODUCTIVITY_RECOMMENDATIONCATEGORY, nil
        case "HEALTH":
            return HEALTH_RECOMMENDATIONCATEGORY, nil
        case "CONFIGURATION":
            return CONFIGURATION_RECOMMENDATIONCATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY, nil
    }
    return 0, errors.New("Unknown RecommendationCategory value: " + v)
}
func SerializeRecommendationCategory(values []RecommendationCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
