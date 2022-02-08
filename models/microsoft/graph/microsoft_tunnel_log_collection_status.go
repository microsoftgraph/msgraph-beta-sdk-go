package graph
import (
    "strings"
    "errors"
)
// 
type MicrosoftTunnelLogCollectionStatus int

const (
    PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS MicrosoftTunnelLogCollectionStatus = iota
    COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
)

func (i MicrosoftTunnelLogCollectionStatus) String() string {
    return []string{"PENDING", "COMPLETED", "FAILED"}[i]
}
func ParseMicrosoftTunnelLogCollectionStatus(v string) (interface{}, error) {
    result := PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "COMPLETED":
            result = COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "FAILED":
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
