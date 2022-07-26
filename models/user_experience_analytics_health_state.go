package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type UserExperienceAnalyticsHealthState int

const (
    UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE UserExperienceAnalyticsHealthState = iota
    INSUFFICIENTDATA_USEREXPERIENCEANALYTICSHEALTHSTATE
    NEEDSATTENTION_USEREXPERIENCEANALYTICSHEALTHSTATE
    MEETINGGOALS_USEREXPERIENCEANALYTICSHEALTHSTATE
)

func (i UserExperienceAnalyticsHealthState) String() string {
    return []string{"unknown", "insufficientData", "needsAttention", "meetingGoals"}[i]
}
func ParseUserExperienceAnalyticsHealthState(v string) (interface{}, error) {
    result := UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE
        case "insufficientData":
            result = INSUFFICIENTDATA_USEREXPERIENCEANALYTICSHEALTHSTATE
        case "needsAttention":
            result = NEEDSATTENTION_USEREXPERIENCEANALYTICSHEALTHSTATE
        case "meetingGoals":
            result = MEETINGGOALS_USEREXPERIENCEANALYTICSHEALTHSTATE
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsHealthState value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsHealthState(values []UserExperienceAnalyticsHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
