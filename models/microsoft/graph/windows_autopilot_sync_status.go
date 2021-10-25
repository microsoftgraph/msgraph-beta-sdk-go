package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS, nil
        case "COMPLETED":
            return COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS, nil
        case "FAILED":
            return FAILED_WINDOWSAUTOPILOTSYNCSTATUS, nil
    }
    return 0, errors.New("Unknown WindowsAutopilotSyncStatus value: " + v)
}
func SerializeWindowsAutopilotSyncStatus(values []WindowsAutopilotSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
