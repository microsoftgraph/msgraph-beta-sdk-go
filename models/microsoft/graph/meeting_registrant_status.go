package graph
import (
    "strings"
    "errors"
)
// 
type MeetingRegistrantStatus int

const (
    REGISTERED_MEETINGREGISTRANTSTATUS MeetingRegistrantStatus = iota
    CANCELED_MEETINGREGISTRANTSTATUS
    PROCESSING_MEETINGREGISTRANTSTATUS
    UNKNOWNFUTUREVALUE_MEETINGREGISTRANTSTATUS
)

func (i MeetingRegistrantStatus) String() string {
    return []string{"REGISTERED", "CANCELED", "PROCESSING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMeetingRegistrantStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "REGISTERED":
            return REGISTERED_MEETINGREGISTRANTSTATUS, nil
        case "CANCELED":
            return CANCELED_MEETINGREGISTRANTSTATUS, nil
        case "PROCESSING":
            return PROCESSING_MEETINGREGISTRANTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MEETINGREGISTRANTSTATUS, nil
    }
    return 0, errors.New("Unknown MeetingRegistrantStatus value: " + v)
}
func SerializeMeetingRegistrantStatus(values []MeetingRegistrantStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
