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
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_GLOBALDEVICEHEALTHSCRIPTSTATE, nil
        case "PENDING":
            return PENDING_GLOBALDEVICEHEALTHSCRIPTSTATE, nil
        case "ENABLED":
            return ENABLED_GLOBALDEVICEHEALTHSCRIPTSTATE, nil
    }
    return 0, errors.New("Unknown GlobalDeviceHealthScriptState value: " + v)
}
func SerializeGlobalDeviceHealthScriptState(values []GlobalDeviceHealthScriptState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
