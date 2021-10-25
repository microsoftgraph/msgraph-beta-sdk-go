package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
        case "HEALTHY":
            return HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
        case "UNHEALTHY":
            return UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
        case "WARNING":
            return WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
        case "OFFLINE":
            return OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
        case "UPGRADEINPROGRESS":
            return UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
        case "UPGRADEFAILED":
            return UPGRADEFAILED_MICROSOFTTUNNELSERVERHEALTHSTATUS, nil
    }
    return 0, errors.New("Unknown MicrosoftTunnelServerHealthStatus value: " + v)
}
func SerializeMicrosoftTunnelServerHealthStatus(values []MicrosoftTunnelServerHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
