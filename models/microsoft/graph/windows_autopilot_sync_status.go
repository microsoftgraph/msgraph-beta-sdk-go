package graph
import (
    "strings"
    "errors"
)
// 
type WindowsAutopilotSyncStatus int

const (
    UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS WindowsAutopilotSyncStatus = iota
    INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS
    COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS
    FAILED_WINDOWSAUTOPILOTSYNCSTATUS
)

func (i WindowsAutopilotSyncStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "COMPLETED", "FAILED"}[i]
}
func ParseWindowsAutopilotSyncStatus(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS
        case "INPROGRESS":
            result = INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS
        case "COMPLETED":
            result = COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS
        case "FAILED":
            result = FAILED_WINDOWSAUTOPILOTSYNCSTATUS
        default:
            return 0, errors.New("Unknown WindowsAutopilotSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotSyncStatus(values []WindowsAutopilotSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
