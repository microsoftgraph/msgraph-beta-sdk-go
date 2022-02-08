package graph
import (
    "strings"
    "errors"
)
// 
type RecommendationStatus int

const (
    ACTIVE_RECOMMENDATIONSTATUS RecommendationStatus = iota
    COMPLETEDBYSYSTEM_RECOMMENDATIONSTATUS
    COMPLETEDBYUSER_RECOMMENDATIONSTATUS
    DISMISSED_RECOMMENDATIONSTATUS
    POSTPONED_RECOMMENDATIONSTATUS
    UNKNOWNFUTUREVALUE_RECOMMENDATIONSTATUS
)

func (i RecommendationStatus) String() string {
    return []string{"ACTIVE", "COMPLETEDBYSYSTEM", "COMPLETEDBYUSER", "DISMISSED", "POSTPONED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRecommendationStatus(v string) (interface{}, error) {
    result := ACTIVE_RECOMMENDATIONSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_RECOMMENDATIONSTATUS
        case "COMPLETEDBYSYSTEM":
            result = COMPLETEDBYSYSTEM_RECOMMENDATIONSTATUS
        case "COMPLETEDBYUSER":
            result = COMPLETEDBYUSER_RECOMMENDATIONSTATUS
        case "DISMISSED":
            result = DISMISSED_RECOMMENDATIONSTATUS
        case "POSTPONED":
            result = POSTPONED_RECOMMENDATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONSTATUS
        default:
            return 0, errors.New("Unknown RecommendationStatus value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationStatus(values []RecommendationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
