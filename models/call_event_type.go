// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type CallEventType int

const (
    CALLSTARTED_CALLEVENTTYPE CallEventType = iota
    CALLENDED_CALLEVENTTYPE
    UNKNOWNFUTUREVALUE_CALLEVENTTYPE
    ROSTERUPDATED_CALLEVENTTYPE
    TRANSCRIPTIONSTARTED_CALLEVENTTYPE
    TRANSCRIPTIONSTOPPED_CALLEVENTTYPE
    RECORDINGSTARTED_CALLEVENTTYPE
    RECORDINGSTOPPED_CALLEVENTTYPE
)

func (i CallEventType) String() string {
    return []string{"callStarted", "callEnded", "unknownFutureValue", "rosterUpdated", "transcriptionStarted", "transcriptionStopped", "recordingStarted", "recordingStopped"}[i]
}
func ParseCallEventType(v string) (any, error) {
    result := CALLSTARTED_CALLEVENTTYPE
    switch v {
        case "callStarted":
            result = CALLSTARTED_CALLEVENTTYPE
        case "callEnded":
            result = CALLENDED_CALLEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CALLEVENTTYPE
        case "rosterUpdated":
            result = ROSTERUPDATED_CALLEVENTTYPE
        case "transcriptionStarted":
            result = TRANSCRIPTIONSTARTED_CALLEVENTTYPE
        case "transcriptionStopped":
            result = TRANSCRIPTIONSTOPPED_CALLEVENTTYPE
        case "recordingStarted":
            result = RECORDINGSTARTED_CALLEVENTTYPE
        case "recordingStopped":
            result = RECORDINGSTOPPED_CALLEVENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCallEventType(values []CallEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CallEventType) isMultiValue() bool {
    return false
}
