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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE, nil
        case "BUSY":
            return BUSY_TEAMWORKDEVICEACTIVITYSTATE, nil
        case "IDLE":
            return IDLE_TEAMWORKDEVICEACTIVITYSTATE, nil
        case "UNAVAILABLE":
            return UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE, nil
    }
    return 0, errors.New("Unknown TeamworkDeviceActivityState value: " + v)
}
func SerializeTeamworkDeviceActivityState(values []TeamworkDeviceActivityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
