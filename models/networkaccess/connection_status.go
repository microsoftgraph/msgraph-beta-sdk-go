package networkaccess
import (
    "errors"
)
type ConnectionStatus int

const (
    OPEN_CONNECTIONSTATUS ConnectionStatus = iota
    ACTIVE_CONNECTIONSTATUS
    CLOSED_CONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_CONNECTIONSTATUS
)

func (i ConnectionStatus) String() string {
    return []string{"open", "active", "closed", "unknownFutureValue"}[i]
}
func ParseConnectionStatus(v string) (any, error) {
    result := OPEN_CONNECTIONSTATUS
    switch v {
        case "open":
            result = OPEN_CONNECTIONSTATUS
        case "active":
            result = ACTIVE_CONNECTIONSTATUS
        case "closed":
            result = CLOSED_CONNECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown ConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeConnectionStatus(values []ConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConnectionStatus) isMultiValue() bool {
    return false
}
