package graph
import (
    "strings"
    "errors"
)
// 
type WindowsAutopilotDeploymentState int

const (
    UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE WindowsAutopilotDeploymentState = iota
    SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
)

func (i WindowsAutopilotDeploymentState) String() string {
    return []string{"UNKNOWN", "SUCCESS", "INPROGRESS", "FAILURE", "SUCCESSWITHTIMEOUT", "NOTATTEMPTED", "DISABLED"}[i]
}
func ParseWindowsAutopilotDeploymentState(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "SUCCESS":
            result = SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "INPROGRESS":
            result = INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "FAILURE":
            result = FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "SUCCESSWITHTIMEOUT":
            result = SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "NOTATTEMPTED":
            result = NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "DISABLED":
            result = DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        default:
            return 0, errors.New("Unknown WindowsAutopilotDeploymentState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotDeploymentState(values []WindowsAutopilotDeploymentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
