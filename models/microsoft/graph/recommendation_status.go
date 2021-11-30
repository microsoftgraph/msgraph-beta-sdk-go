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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_RECOMMENDATIONSTATUS, nil
        case "COMPLETEDBYSYSTEM":
            return COMPLETEDBYSYSTEM_RECOMMENDATIONSTATUS, nil
        case "COMPLETEDBYUSER":
            return COMPLETEDBYUSER_RECOMMENDATIONSTATUS, nil
        case "DISMISSED":
            return DISMISSED_RECOMMENDATIONSTATUS, nil
        case "POSTPONED":
            return POSTPONED_RECOMMENDATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RECOMMENDATIONSTATUS, nil
    }
    return 0, errors.New("Unknown RecommendationStatus value: " + v)
}
func SerializeRecommendationStatus(values []RecommendationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
