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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_CONNECTORSTATUS, nil
        case "INACTIVE":
            return INACTIVE_CONNECTORSTATUS, nil
    }
    return 0, errors.New("Unknown ConnectorStatus value: " + v)
}
func SerializeConnectorStatus(values []ConnectorStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
