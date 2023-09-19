package models
import (
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
    return []string{"registered", "canceled", "processing", "unknownFutureValue"}[i]
}
func ParseMeetingRegistrantStatus(v string) (any, error) {
    result := REGISTERED_MEETINGREGISTRANTSTATUS
    switch v {
        case "registered":
            result = REGISTERED_MEETINGREGISTRANTSTATUS
        case "canceled":
            result = CANCELED_MEETINGREGISTRANTSTATUS
        case "processing":
            result = PROCESSING_MEETINGREGISTRANTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MEETINGREGISTRANTSTATUS
        default:
            return 0, errors.New("Unknown MeetingRegistrantStatus value: " + v)
    }
    return &result, nil
}
func SerializeMeetingRegistrantStatus(values []MeetingRegistrantStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MeetingRegistrantStatus) isMultiValue() bool {
    return false
}
