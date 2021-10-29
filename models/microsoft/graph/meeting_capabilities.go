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
    switch strings.ToUpper(v) {
        case "QUESTIONANDANSWER":
            return QUESTIONANDANSWER_MEETINGCAPABILITIES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES, nil
    }
    return 0, errors.New("Unknown MeetingCapabilities value: " + v)
}
func SerializeMeetingCapabilities(values []MeetingCapabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
