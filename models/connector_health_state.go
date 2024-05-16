package models
// Connector health state for connector status
type ConnectorHealthState int

const (
    // Indicates a healthy connector status and no action required.
    HEALTHY_CONNECTORHEALTHSTATE ConnectorHealthState = iota
    // Indicates that a connector needs attention.
    WARNING_CONNECTORHEALTHSTATE
    // Indicates that a connector needs immediate attention to retain functionality.
    UNHEALTHY_CONNECTORHEALTHSTATE
    // unknown
    UNKNOWN_CONNECTORHEALTHSTATE
)

func (i ConnectorHealthState) String() string {
    return []string{"healthy", "warning", "unhealthy", "unknown"}[i]
}
func ParseConnectorHealthState(v string) (any, error) {
    result := HEALTHY_CONNECTORHEALTHSTATE
    switch v {
        case "healthy":
            result = HEALTHY_CONNECTORHEALTHSTATE
        case "warning":
            result = WARNING_CONNECTORHEALTHSTATE
        case "unhealthy":
            result = UNHEALTHY_CONNECTORHEALTHSTATE
        case "unknown":
            result = UNKNOWN_CONNECTORHEALTHSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConnectorHealthState(values []ConnectorHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConnectorHealthState) isMultiValue() bool {
    return false
}
