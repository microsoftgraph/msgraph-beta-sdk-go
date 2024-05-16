package networkaccess
type RemoteNetworkStatus int

const (
    TUNNELDISCONNECTED_REMOTENETWORKSTATUS RemoteNetworkStatus = iota
    TUNNELCONNECTED_REMOTENETWORKSTATUS
    BGPDISCONNECTED_REMOTENETWORKSTATUS
    BGPCONNECTED_REMOTENETWORKSTATUS
    REMOTENETWORKALIVE_REMOTENETWORKSTATUS
    UNKNOWNFUTUREVALUE_REMOTENETWORKSTATUS
)

func (i RemoteNetworkStatus) String() string {
    return []string{"tunnelDisconnected", "tunnelConnected", "bgpDisconnected", "bgpConnected", "remoteNetworkAlive", "unknownFutureValue"}[i]
}
func ParseRemoteNetworkStatus(v string) (any, error) {
    result := TUNNELDISCONNECTED_REMOTENETWORKSTATUS
    switch v {
        case "tunnelDisconnected":
            result = TUNNELDISCONNECTED_REMOTENETWORKSTATUS
        case "tunnelConnected":
            result = TUNNELCONNECTED_REMOTENETWORKSTATUS
        case "bgpDisconnected":
            result = BGPDISCONNECTED_REMOTENETWORKSTATUS
        case "bgpConnected":
            result = BGPCONNECTED_REMOTENETWORKSTATUS
        case "remoteNetworkAlive":
            result = REMOTENETWORKALIVE_REMOTENETWORKSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REMOTENETWORKSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRemoteNetworkStatus(values []RemoteNetworkStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RemoteNetworkStatus) isMultiValue() bool {
    return false
}
