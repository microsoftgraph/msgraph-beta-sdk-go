package graph
import (
    "strings"
    "errors"
)
// 
type ConnectorStatus int

const (
    ACTIVE_CONNECTORSTATUS ConnectorStatus = iota
    INACTIVE_CONNECTORSTATUS
)

func (i ConnectorStatus) String() string {
    return []string{"ACTIVE", "INACTIVE"}[i]
}
func ParseConnectorStatus(v string) (interface{}, error) {
    result := ACTIVE_CONNECTORSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_CONNECTORSTATUS
        case "INACTIVE":
            result = INACTIVE_CONNECTORSTATUS
        default:
            return 0, errors.New("Unknown ConnectorStatus value: " + v)
    }
    return &result, nil
}
func SerializeConnectorStatus(values []ConnectorStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
