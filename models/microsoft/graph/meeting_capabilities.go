package graph
import (
    "strings"
    "errors"
)
// 
type MeetingCapabilities int

const (
    QUESTIONANDANSWER_MEETINGCAPABILITIES MeetingCapabilities = iota
    UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES
)

func (i MeetingCapabilities) String() string {
    return []string{"QUESTIONANDANSWER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMeetingCapabilities(v string) (interface{}, error) {
    result := QUESTIONANDANSWER_MEETINGCAPABILITIES
    switch strings.ToUpper(v) {
        case "QUESTIONANDANSWER":
            result = QUESTIONANDANSWER_MEETINGCAPABILITIES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES
        default:
            return 0, errors.New("Unknown MeetingCapabilities value: " + v)
    }
    return &result, nil
}
func SerializeMeetingCapabilities(values []MeetingCapabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
