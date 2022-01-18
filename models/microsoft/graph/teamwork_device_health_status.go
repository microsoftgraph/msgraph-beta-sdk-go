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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS, nil
        case "OFFLINE":
            return OFFLINE_TEAMWORKDEVICEHEALTHSTATUS, nil
        case "CRITICAL":
            return CRITICAL_TEAMWORKDEVICEHEALTHSTATUS, nil
        case "NONURGENT":
            return NONURGENT_TEAMWORKDEVICEHEALTHSTATUS, nil
        case "HEALTHY":
            return HEALTHY_TEAMWORKDEVICEHEALTHSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKDEVICEHEALTHSTATUS, nil
    }
    return 0, errors.New("Unknown TeamworkDeviceHealthStatus value: " + v)
}
func SerializeTeamworkDeviceHealthStatus(values []TeamworkDeviceHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
