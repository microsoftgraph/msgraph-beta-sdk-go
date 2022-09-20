package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type MicrosoftTunnelLogCollectionStatus int

const (
    // Log collection is in progress
    PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS MicrosoftTunnelLogCollectionStatus = iota
    // Log collection is completed
    COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    // Log collection has failed
    FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
)

func (i MicrosoftTunnelLogCollectionStatus) String() string {
    return []string{"pending", "completed", "failed"}[i]
}
func ParseMicrosoftTunnelLogCollectionStatus(v string) (interface{}, error) {
    result := PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    switch v {
        case "pending":
            result = PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "completed":
            result = COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "failed":
            result = FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        default:
            return 0, errors.New("Unknown MicrosoftTunnelLogCollectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftTunnelLogCollectionStatus(values []MicrosoftTunnelLogCollectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
