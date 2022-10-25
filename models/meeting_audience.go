package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type MeetingAudience int

const (
    EVERYONE_MEETINGAUDIENCE MeetingAudience = iota
    ORGANIZATION_MEETINGAUDIENCE
    UNKNOWNFUTUREVALUE_MEETINGAUDIENCE
)

func (i MeetingAudience) String() string {
    return []string{"everyone", "organization", "unknownFutureValue"}[i]
}
func ParseMeetingAudience(v string) (interface{}, error) {
    result := EVERYONE_MEETINGAUDIENCE
    switch v {
        case "everyone":
            result = EVERYONE_MEETINGAUDIENCE
        case "organization":
            result = ORGANIZATION_MEETINGAUDIENCE
        case "unknownFutureValue":
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
