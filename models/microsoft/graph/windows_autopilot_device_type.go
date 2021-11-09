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
    switch strings.ToUpper(v) {
        case "WINDOWSPC":
            return WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE, nil
        case "SURFACEHUB2":
            return SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE, nil
        case "HOLOLENS":
            return HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE, nil
        case "SURFACEHUB2S":
            return SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE, nil
        case "VIRTUALMACHINE":
            return VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE, nil
    }
    return 0, errors.New("Unknown WindowsAutopilotDeviceType value: " + v)
}
func SerializeWindowsAutopilotDeviceType(values []WindowsAutopilotDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
