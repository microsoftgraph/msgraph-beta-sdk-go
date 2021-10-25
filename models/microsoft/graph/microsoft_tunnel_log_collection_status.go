package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS, nil
        case "FAILED":
            return FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS, nil
    }
    return 0, errors.New("Unknown MicrosoftTunnelLogCollectionStatus value: " + v)
}
func SerializeMicrosoftTunnelLogCollectionStatus(values []MicrosoftTunnelLogCollectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
