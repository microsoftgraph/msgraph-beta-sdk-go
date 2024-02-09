package models
import (
    "errors"
)
type ConnectorStatus int

const (
    ACTIVE_CONNECTORSTATUS ConnectorStatus = iota
    INACTIVE_CONNECTORSTATUS
)

func (i ConnectorStatus) String() string {
    return []string{"active", "inactive"}[i]
}
func ParseConnectorStatus(v string) (any, error) {
    result := ACTIVE_CONNECTORSTATUS
    switch v {
        case "active":
            result = ACTIVE_CONNECTORSTATUS
        case "inactive":
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
func (i ConnectorStatus) isMultiValue() bool {
    return false
}
