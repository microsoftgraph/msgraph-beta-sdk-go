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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
        case "SUCCESS":
            return SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
        case "INPROGRESS":
            return INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
        case "FAILURE":
            return FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
        case "SUCCESSWITHTIMEOUT":
            return SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
        case "NOTATTEMPTED":
            return NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
        case "DISABLED":
            return DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE, nil
    }
    return 0, errors.New("Unknown WindowsAutopilotDeploymentState value: " + v)
}
func SerializeWindowsAutopilotDeploymentState(values []WindowsAutopilotDeploymentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
