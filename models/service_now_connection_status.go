package models
// Status of ServiceNow Connection
type ServiceNowConnectionStatus int

const (
    // Tenant has disabled the connection
    DISABLED_SERVICENOWCONNECTIONSTATUS ServiceNowConnectionStatus = iota
    // Tenant has enabled the connection
    ENABLED_SERVICENOWCONNECTIONSTATUS
    // Future authentication method to be added here.
    UNKNOWNFUTUREVALUE_SERVICENOWCONNECTIONSTATUS
)

func (i ServiceNowConnectionStatus) String() string {
    return []string{"disabled", "enabled", "unknownFutureValue"}[i]
}
func ParseServiceNowConnectionStatus(v string) (any, error) {
    result := DISABLED_SERVICENOWCONNECTIONSTATUS
    switch v {
        case "disabled":
            result = DISABLED_SERVICENOWCONNECTIONSTATUS
        case "enabled":
            result = ENABLED_SERVICENOWCONNECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SERVICENOWCONNECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServiceNowConnectionStatus(values []ServiceNowConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServiceNowConnectionStatus) isMultiValue() bool {
    return false
}
