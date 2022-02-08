package graph
import (
    "strings"
    "errors"
)
// 
type MeetingAudience int

const (
    EVERYONE_MEETINGAUDIENCE MeetingAudience = iota
    ORGANIZATION_MEETINGAUDIENCE
    UNKNOWNFUTUREVALUE_MEETINGAUDIENCE
)

func (i MeetingAudience) String() string {
    return []string{"EVERYONE", "ORGANIZATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMeetingAudience(v string) (interface{}, error) {
    result := EVERYONE_MEETINGAUDIENCE
    switch strings.ToUpper(v) {
        case "EVERYONE":
            result = EVERYONE_MEETINGAUDIENCE
        case "ORGANIZATION":
            result = ORGANIZATION_MEETINGAUDIENCE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MEETINGAUDIENCE
        default:
            return 0, errors.New("Unknown MeetingAudience value: " + v)
    }
    return &result, nil
}
func SerializeMeetingAudience(values []MeetingAudience) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
