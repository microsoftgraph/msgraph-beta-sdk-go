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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY, nil
        case "INFORMATIONAL":
            return INFORMATIONAL_USEREXPERIENCEANALYTICSINSIGHTSEVERITY, nil
        case "WARNING":
            return WARNING_USEREXPERIENCEANALYTICSINSIGHTSEVERITY, nil
        case "ERROR":
            return ERROR_USEREXPERIENCEANALYTICSINSIGHTSEVERITY, nil
    }
    return 0, errors.New("Unknown UserExperienceAnalyticsInsightSeverity value: " + v)
}
func SerializeUserExperienceAnalyticsInsightSeverity(values []UserExperienceAnalyticsInsightSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
