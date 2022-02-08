package graph
import (
    "strings"
    "errors"
)
// 
type UserExperienceAnalyticsInsightSeverity int

const (
    NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY UserExperienceAnalyticsInsightSeverity = iota
    INFORMATIONAL_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
    WARNING_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
    ERROR_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
)

func (i UserExperienceAnalyticsInsightSeverity) String() string {
    return []string{"NONE", "INFORMATIONAL", "WARNING", "ERROR"}[i]
}
func ParseUserExperienceAnalyticsInsightSeverity(v string) (interface{}, error) {
    result := NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        case "INFORMATIONAL":
            result = INFORMATIONAL_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        case "WARNING":
            result = WARNING_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        case "ERROR":
            result = ERROR_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsInsightSeverity value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsInsightSeverity(values []UserExperienceAnalyticsInsightSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
