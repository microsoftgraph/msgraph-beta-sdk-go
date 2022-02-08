package graph
import (
    "strings"
    "errors"
)
// 
type WindowsAutopilotDeviceType int

const (
    WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE WindowsAutopilotDeviceType = iota
    SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE
    HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE
    SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE
    VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE
)

func (i WindowsAutopilotDeviceType) String() string {
    return []string{"WINDOWSPC", "SURFACEHUB2", "HOLOLENS", "SURFACEHUB2S", "VIRTUALMACHINE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWindowsAutopilotDeviceType(v string) (interface{}, error) {
    result := WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE
    switch strings.ToUpper(v) {
        case "WINDOWSPC":
            result = WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE
        case "SURFACEHUB2":
            result = SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE
        case "HOLOLENS":
            result = HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE
        case "SURFACEHUB2S":
            result = SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE
        case "VIRTUALMACHINE":
            result = VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE
        default:
            return 0, errors.New("Unknown WindowsAutopilotDeviceType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotDeviceType(values []WindowsAutopilotDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
