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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TEAMWORKCONNECTIONSTATUS, nil
        case "CONNECTED":
            return CONNECTED_TEAMWORKCONNECTIONSTATUS, nil
        case "DISCONNECTED":
            return DISCONNECTED_TEAMWORKCONNECTIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKCONNECTIONSTATUS, nil
    }
    return 0, errors.New("Unknown TeamworkConnectionStatus value: " + v)
}
func SerializeTeamworkConnectionStatus(values []TeamworkConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
