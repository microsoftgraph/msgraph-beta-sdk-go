package models
// Determines how often to check if a peer connection is still alive
type VpnDeadPeerDetectionRate int

const (
    // Medium
    MEDIUM_VPNDEADPEERDETECTIONRATE VpnDeadPeerDetectionRate = iota
    // None
    NONE_VPNDEADPEERDETECTIONRATE
    // Low
    LOW_VPNDEADPEERDETECTIONRATE
    // High
    HIGH_VPNDEADPEERDETECTIONRATE
)

func (i VpnDeadPeerDetectionRate) String() string {
    return []string{"medium", "none", "low", "high"}[i]
}
func ParseVpnDeadPeerDetectionRate(v string) (any, error) {
    result := MEDIUM_VPNDEADPEERDETECTIONRATE
    switch v {
        case "medium":
            result = MEDIUM_VPNDEADPEERDETECTIONRATE
        case "none":
            result = NONE_VPNDEADPEERDETECTIONRATE
        case "low":
            result = LOW_VPNDEADPEERDETECTIONRATE
        case "high":
            result = HIGH_VPNDEADPEERDETECTIONRATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVpnDeadPeerDetectionRate(values []VpnDeadPeerDetectionRate) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VpnDeadPeerDetectionRate) isMultiValue() bool {
    return false
}
