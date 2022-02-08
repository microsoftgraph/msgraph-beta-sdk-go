package graph
import (
    "strings"
    "errors"
)
// 
type WindowsAutopilotProfileAssignmentStatus int

const (
    UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS WindowsAutopilotProfileAssignmentStatus = iota
    ASSIGNEDINSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    ASSIGNEDOUTOFSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    ASSIGNEDUNKOWNSYNCSTATE_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    NOTASSIGNED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    PENDING_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    FAILED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
)

func (i WindowsAutopilotProfileAssignmentStatus) String() string {
    return []string{"UNKNOWN", "ASSIGNEDINSYNC", "ASSIGNEDOUTOFSYNC", "ASSIGNEDUNKOWNSYNCSTATE", "NOTASSIGNED", "PENDING", "FAILED"}[i]
}
func ParseWindowsAutopilotProfileAssignmentStatus(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "ASSIGNEDINSYNC":
            result = ASSIGNEDINSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "ASSIGNEDOUTOFSYNC":
            result = ASSIGNEDOUTOFSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "ASSIGNEDUNKOWNSYNCSTATE":
            result = ASSIGNEDUNKOWNSYNCSTATE_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "NOTASSIGNED":
            result = NOTASSIGNED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "PENDING":
            result = PENDING_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "FAILED":
            result = FAILED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        default:
            return 0, errors.New("Unknown WindowsAutopilotProfileAssignmentStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotProfileAssignmentStatus(values []WindowsAutopilotProfileAssignmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
