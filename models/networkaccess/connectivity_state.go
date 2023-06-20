package networkaccess
import (
    "errors"
)
// 
type ConnectivityState int

const (
    PENDING_CONNECTIVITYSTATE ConnectivityState = iota
    CONNECTED_CONNECTIVITYSTATE
    INACTIVE_CONNECTIVITYSTATE
    ERROR_CONNECTIVITYSTATE
    UNKNOWNFUTUREVALUE_CONNECTIVITYSTATE
)

func (i ConnectivityState) String() string {
    return []string{"pending", "connected", "inactive", "error", "unknownFutureValue"}[i]
}
func ParseConnectivityState(v string) (any, error) {
    result := PENDING_CONNECTIVITYSTATE
    switch v {
        case "pending":
            result = PENDING_CONNECTIVITYSTATE
        case "connected":
            result = CONNECTED_CONNECTIVITYSTATE
        case "inactive":
            result = INACTIVE_CONNECTIVITYSTATE
        case "error":
            result = ERROR_CONNECTIVITYSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONNECTIVITYSTATE
        default:
            return 0, errors.New("Unknown ConnectivityState value: " + v)
    }
    return &result, nil
}
func SerializeConnectivityState(values []ConnectivityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
