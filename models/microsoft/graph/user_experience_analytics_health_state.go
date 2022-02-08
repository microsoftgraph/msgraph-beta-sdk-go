package graph
import (
    "strings"
    "errors"
)
// 
type UserExperienceAnalyticsHealthState int

const (
    UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE UserExperienceAnalyticsHealthState = iota
    INSUFFICIENTDATA_USEREXPERIENCEANALYTICSHEALTHSTATE
    NEEDSATTENTION_USEREXPERIENCEANALYTICSHEALTHSTATE
    MEETINGGOALS_USEREXPERIENCEANALYTICSHEALTHSTATE
)

func (i UserExperienceAnalyticsHealthState) String() string {
    return []string{"UNKNOWN", "INSUFFICIENTDATA", "NEEDSATTENTION", "MEETINGGOALS"}[i]
}
func ParseUserExperienceAnalyticsHealthState(v string) (interface{}, error) {
    result := UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE
        case "INSUFFICIENTDATA":
            result = INSUFFICIENTDATA_USEREXPERIENCEANALYTICSHEALTHSTATE
        case "NEEDSATTENTION":
            result = NEEDSATTENTION_USEREXPERIENCEANALYTICSHEALTHSTATE
        case "MEETINGGOALS":
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
