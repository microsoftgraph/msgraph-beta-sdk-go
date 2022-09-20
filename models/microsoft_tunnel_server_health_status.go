package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type MicrosoftTunnelServerHealthStatus int

const (
    // The state is unknown
    UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS MicrosoftTunnelServerHealthStatus = iota
    // The state is healthy
    HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // The state is unhealthy
    UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // The state is warning
    WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // The state is offline
    OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // The state is upgradeInProgress
    UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // The state is upgradeFailed
    UPGRADEFAILED_MICROSOFTTUNNELSERVERHEALTHSTATUS
)

func (i MicrosoftTunnelServerHealthStatus) String() string {
    return []string{"unknown", "healthy", "unhealthy", "warning", "offline", "upgradeInProgress", "upgradeFailed"}[i]
}
func ParseMicrosoftTunnelServerHealthStatus(v string) (interface{}, error) {
    result := UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "healthy":
            result = HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "unhealthy":
            result = UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "warning":
            result = WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "offline":
            result = OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "upgradeInProgress":
            result = UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "upgradeFailed":
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
