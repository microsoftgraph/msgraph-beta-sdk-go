package models
// Enum type that represent the status of log collection
type MicrosoftTunnelLogCollectionStatus int

const (
    // Indicates that the log collection is in progress
    PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS MicrosoftTunnelLogCollectionStatus = iota
    // Indicates that the log collection is completed
    COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    // Indicates that the log collection has failed
    FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    // Placeholder value for future expansion enums
    UNKNOWNFUTUREVALUE_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
)

func (i MicrosoftTunnelLogCollectionStatus) String() string {
    return []string{"pending", "completed", "failed", "unknownFutureValue"}[i]
}
func ParseMicrosoftTunnelLogCollectionStatus(v string) (any, error) {
    result := PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    switch v {
        case "pending":
            result = PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "completed":
            result = COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "failed":
            result = FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        default:
            return nil, nil
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
func (i MicrosoftTunnelLogCollectionStatus) isMultiValue() bool {
    return false
}
