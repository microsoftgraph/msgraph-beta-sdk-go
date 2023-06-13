package models
import (
    "errors"
)
// Indicates the level of prevalence of the correlation group features in the anomaly. Possible values are: high, medium or low
type UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence int

const (
    // Indicates that we have a high prevalence in the correlation between the anomaly and correlation group.
    HIGH_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence = iota
    // Indicates that we have a medium prevalence in the correlation between the anomaly and correlation group.
    MEDIUM_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
    // Indicates that we have a low prevalence in the correlation between the anomaly and correlation group.
    LOW_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
)

func (i UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence) String() string {
    return []string{"high", "medium", "low", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsAnomalyCorrelationGroupPrevalence(v string) (any, error) {
    result := HIGH_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
    switch v {
        case "high":
            result = HIGH_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
        case "medium":
            result = MEDIUM_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
        case "low":
            result = LOW_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYCORRELATIONGROUPPREVALENCE
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsAnomalyCorrelationGroupPrevalence(values []UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
