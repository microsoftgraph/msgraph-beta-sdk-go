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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
        case "ASSIGNEDINSYNC":
            return ASSIGNEDINSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
        case "ASSIGNEDOUTOFSYNC":
            return ASSIGNEDOUTOFSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
        case "ASSIGNEDUNKOWNSYNCSTATE":
            return ASSIGNEDUNKOWNSYNCSTATE_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
        case "NOTASSIGNED":
            return NOTASSIGNED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
        case "PENDING":
            return PENDING_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
        case "FAILED":
            return FAILED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS, nil
    }
    return 0, errors.New("Unknown WindowsAutopilotProfileAssignmentStatus value: " + v)
}
func SerializeWindowsAutopilotProfileAssignmentStatus(values []WindowsAutopilotProfileAssignmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
