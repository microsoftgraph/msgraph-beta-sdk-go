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
    switch strings.ToUpper(v) {
        case "EVERYONE":
            return EVERYONE_MEETINGAUDIENCE, nil
        case "ORGANIZATION":
            return ORGANIZATION_MEETINGAUDIENCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MEETINGAUDIENCE, nil
    }
    return 0, errors.New("Unknown MeetingAudience value: " + v)
}
func SerializeMeetingAudience(values []MeetingAudience) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
