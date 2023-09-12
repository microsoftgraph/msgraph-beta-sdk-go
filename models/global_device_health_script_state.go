package models
import (
    "errors"
)
// Indicates whether global device health scripts are enabled and are in which state
type GlobalDeviceHealthScriptState int

const (
    // Global device health scripts are not configured
    NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE GlobalDeviceHealthScriptState = iota
    // Global device health scripts are configured but not fully enabled
    PENDING_GLOBALDEVICEHEALTHSCRIPTSTATE
    // Global device health scripts are enabled and ready to use
    ENABLED_GLOBALDEVICEHEALTHSCRIPTSTATE
)

func (i GlobalDeviceHealthScriptState) String() string {
    return []string{"notConfigured", "pending", "enabled"}[i]
}
func ParseGlobalDeviceHealthScriptState(v string) (any, error) {
    result := NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE
        case "pending":
            result = PENDING_GLOBALDEVICEHEALTHSCRIPTSTATE
        case "enabled":
            result = ENABLED_GLOBALDEVICEHEALTHSCRIPTSTATE
        default:
            return 0, errors.New("Unknown GlobalDeviceHealthScriptState value: " + v)
    }
    return &result, nil
}
func SerializeGlobalDeviceHealthScriptState(values []GlobalDeviceHealthScriptState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GlobalDeviceHealthScriptState) isMultiValue() bool {
    return false
}
