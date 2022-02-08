package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkConnectionStatus int

const (
    UNKNOWN_TEAMWORKCONNECTIONSTATUS TeamworkConnectionStatus = iota
    CONNECTED_TEAMWORKCONNECTIONSTATUS
    DISCONNECTED_TEAMWORKCONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_TEAMWORKCONNECTIONSTATUS
)

func (i TeamworkConnectionStatus) String() string {
    return []string{"UNKNOWN", "CONNECTED", "DISCONNECTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkConnectionStatus(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKCONNECTIONSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TEAMWORKCONNECTIONSTATUS
        case "CONNECTED":
            result = CONNECTED_TEAMWORKCONNECTIONSTATUS
        case "DISCONNECTED":
            result = DISCONNECTED_TEAMWORKCONNECTIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKCONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown TeamworkConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkConnectionStatus(values []TeamworkConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
