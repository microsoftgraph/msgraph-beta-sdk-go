package models
import (
    "errors"
)
// Deployment states for Autopilot devices
type WindowsAutopilotDeploymentState int

const (
    // The deployment state is unknown.
    UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE WindowsAutopilotDeploymentState = iota
    // The deployment succeeded.
    SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    // The deployment state is in progress.
    INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    // The deployment failed.
    FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    // The deployment timed out but user clicked past failure.
    SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    // The deployment was not run.
    NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    // The deployment is disabled.
    DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    // The deployment succeeded after hitting an initial timeout failure.
    SUCCESSONRETRY_WINDOWSAUTOPILOTDEPLOYMENTSTATE
)

func (i WindowsAutopilotDeploymentState) String() string {
    return []string{"unknown", "success", "inProgress", "failure", "successWithTimeout", "notAttempted", "disabled", "successOnRetry"}[i]
}
func ParseWindowsAutopilotDeploymentState(v string) (any, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "success":
            result = SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "inProgress":
            result = INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "failure":
            result = FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "successWithTimeout":
            result = SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "notAttempted":
            result = NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "disabled":
            result = DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "successOnRetry":
            result = SUCCESSONRETRY_WINDOWSAUTOPILOTDEPLOYMENTSTATE
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
