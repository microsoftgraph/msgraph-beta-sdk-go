package graph
import (
    "strings"
    "errors"
)
// 
type GlobalDeviceHealthScriptState int

const (
    NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE GlobalDeviceHealthScriptState = iota
    PENDING_GLOBALDEVICEHEALTHSCRIPTSTATE
    ENABLED_GLOBALDEVICEHEALTHSCRIPTSTATE
)

func (i GlobalDeviceHealthScriptState) String() string {
    return []string{"NOTCONFIGURED", "PENDING", "ENABLED"}[i]
}
func ParseGlobalDeviceHealthScriptState(v string) (interface{}, error) {
    result := NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE
        case "PENDING":
            result = PENDING_GLOBALDEVICEHEALTHSCRIPTSTATE
        case "ENABLED":
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
