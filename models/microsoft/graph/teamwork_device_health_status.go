package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkDeviceHealthStatus int

const (
    UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS TeamworkDeviceHealthStatus = iota
    OFFLINE_TEAMWORKDEVICEHEALTHSTATUS
    CRITICAL_TEAMWORKDEVICEHEALTHSTATUS
    NONURGENT_TEAMWORKDEVICEHEALTHSTATUS
    HEALTHY_TEAMWORKDEVICEHEALTHSTATUS
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICEHEALTHSTATUS
)

func (i TeamworkDeviceHealthStatus) String() string {
    return []string{"UNKNOWN", "OFFLINE", "CRITICAL", "NONURGENT", "HEALTHY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkDeviceHealthStatus(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS
        case "OFFLINE":
            result = OFFLINE_TEAMWORKDEVICEHEALTHSTATUS
        case "CRITICAL":
            result = CRITICAL_TEAMWORKDEVICEHEALTHSTATUS
        case "NONURGENT":
            result = NONURGENT_TEAMWORKDEVICEHEALTHSTATUS
        case "HEALTHY":
            result = HEALTHY_TEAMWORKDEVICEHEALTHSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEHEALTHSTATUS
        default:
            return 0, errors.New("Unknown TeamworkDeviceHealthStatus value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkDeviceHealthStatus(values []TeamworkDeviceHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
