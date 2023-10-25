package models
import (
    "errors"
)
// 
type CallEventType int

const (
    CALLSTARTED_CALLEVENTTYPE CallEventType = iota
    CALLENDED_CALLEVENTTYPE
    UNKNOWNFUTUREVALUE_CALLEVENTTYPE
)

func (i CallEventType) String() string {
    return []string{"callStarted", "callEnded", "unknownFutureValue"}[i]
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
        default:
            return 0, errors.New("Unknown CallEventType value: " + v)
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
