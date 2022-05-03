package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type WindowsAutopilotDeviceRemediationState int

const (
    UNKNOWN_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE WindowsAutopilotDeviceRemediationState = iota
    NOREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    AUTOMATICREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    MANUALREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
)

func (i WindowsAutopilotDeviceRemediationState) String() string {
    return []string{"UNKNOWN", "NOREMEDIATIONREQUIRED", "AUTOMATICREMEDIATIONREQUIRED", "MANUALREMEDIATIONREQUIRED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWindowsAutopilotDeviceRemediationState(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "NOREMEDIATIONREQUIRED":
            result = NOREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "AUTOMATICREMEDIATIONREQUIRED":
            result = AUTOMATICREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "MANUALREMEDIATIONREQUIRED":
            result = MANUALREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        default:
            return 0, errors.New("Unknown WindowsAutopilotDeviceRemediationState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotDeviceRemediationState(values []WindowsAutopilotDeviceRemediationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
