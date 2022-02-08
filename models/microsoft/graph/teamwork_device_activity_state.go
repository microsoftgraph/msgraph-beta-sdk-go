package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkDeviceActivityState int

const (
    UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE TeamworkDeviceActivityState = iota
    BUSY_TEAMWORKDEVICEACTIVITYSTATE
    IDLE_TEAMWORKDEVICEACTIVITYSTATE
    UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE
)

func (i TeamworkDeviceActivityState) String() string {
    return []string{"UNKNOWN", "BUSY", "IDLE", "UNAVAILABLE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkDeviceActivityState(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE
        case "BUSY":
            result = BUSY_TEAMWORKDEVICEACTIVITYSTATE
        case "IDLE":
            result = IDLE_TEAMWORKDEVICEACTIVITYSTATE
        case "UNAVAILABLE":
            result = UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE
        default:
            return 0, errors.New("Unknown TeamworkDeviceActivityState value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkDeviceActivityState(values []TeamworkDeviceActivityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
