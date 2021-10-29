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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_USEREXPERIENCEANALYTICSHEALTHSTATE, nil
        case "INSUFFICIENTDATA":
            return INSUFFICIENTDATA_USEREXPERIENCEANALYTICSHEALTHSTATE, nil
        case "NEEDSATTENTION":
            return NEEDSATTENTION_USEREXPERIENCEANALYTICSHEALTHSTATE, nil
        case "MEETINGGOALS":
            return MEETINGGOALS_USEREXPERIENCEANALYTICSHEALTHSTATE, nil
    }
    return 0, errors.New("Unknown UserExperienceAnalyticsHealthState value: " + v)
}
func SerializeUserExperienceAnalyticsHealthState(values []UserExperienceAnalyticsHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
