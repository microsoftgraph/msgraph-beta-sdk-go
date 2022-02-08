package graph
import (
    "strings"
    "errors"
)
// 
type MicrosoftTunnelServerHealthStatus int

const (
    UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS MicrosoftTunnelServerHealthStatus = iota
    HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
    UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
    WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS
    OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS
    UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS
    UPGRADEFAILED_MICROSOFTTUNNELSERVERHEALTHSTATUS
)

func (i MicrosoftTunnelServerHealthStatus) String() string {
    return []string{"UNKNOWN", "HEALTHY", "UNHEALTHY", "WARNING", "OFFLINE", "UPGRADEINPROGRESS", "UPGRADEFAILED"}[i]
}
func ParseMicrosoftTunnelServerHealthStatus(v string) (interface{}, error) {
    result := UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "HEALTHY":
            result = HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "UNHEALTHY":
            result = UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "WARNING":
            result = WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "OFFLINE":
            result = OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "UPGRADEINPROGRESS":
            result = UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "UPGRADEFAILED":
            result = UPGRADEFAILED_MICROSOFTTUNNELSERVERHEALTHSTATUS
        default:
            return 0, errors.New("Unknown MicrosoftTunnelServerHealthStatus value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftTunnelServerHealthStatus(values []MicrosoftTunnelServerHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
