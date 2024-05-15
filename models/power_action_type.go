package models
// Power action types
type PowerActionType int

const (
    // Not configured
    NOTCONFIGURED_POWERACTIONTYPE PowerActionType = iota
    // No action
    NOACTION_POWERACTIONTYPE
    // Put device in sleep state
    SLEEP_POWERACTIONTYPE
    // Put device in hibernate state
    HIBERNATE_POWERACTIONTYPE
    // Shutdown device
    SHUTDOWN_POWERACTIONTYPE
)

func (i PowerActionType) String() string {
    return []string{"notConfigured", "noAction", "sleep", "hibernate", "shutdown"}[i]
}
func ParsePowerActionType(v string) (any, error) {
    result := NOTCONFIGURED_POWERACTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_POWERACTIONTYPE
        case "noAction":
            result = NOACTION_POWERACTIONTYPE
        case "sleep":
            result = SLEEP_POWERACTIONTYPE
        case "hibernate":
            result = HIBERNATE_POWERACTIONTYPE
        case "shutdown":
            result = SHUTDOWN_POWERACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePowerActionType(values []PowerActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PowerActionType) isMultiValue() bool {
    return false
}
