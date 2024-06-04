package models
type WindowsAutopilotDeviceType int

const (
    // Default. Indicates that the device type  is a Windows PC.
    WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE WindowsAutopilotDeviceType = iota
    // Indicates that the device type is a HoloLens.
    HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE
    // Surface Hub 2
    SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE
    // SurfaceHub2S
    SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE
    // VirtualMachine
    VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE
)

func (i WindowsAutopilotDeviceType) String() string {
    return []string{"windowsPc", "holoLens", "surfaceHub2", "surfaceHub2S", "virtualMachine", "unknownFutureValue"}[i]
}
func ParseWindowsAutopilotDeviceType(v string) (any, error) {
    result := WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE
    switch v {
        case "windowsPc":
            result = WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE
        case "holoLens":
            result = HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE
        case "surfaceHub2":
            result = SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE
        case "surfaceHub2S":
            result = SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE
        case "virtualMachine":
            result = VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE
        default:
            return nil, nil
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
func (i WindowsAutopilotDeviceType) isMultiValue() bool {
    return false
}
