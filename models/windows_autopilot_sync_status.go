// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type WindowsAutopilotSyncStatus int

const (
    // Unknown sync status
    UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS WindowsAutopilotSyncStatus = iota
    // Sync is in progress
    INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS
    // Sync completed.
    COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS
    // Sync failed.
    FAILED_WINDOWSAUTOPILOTSYNCSTATUS
)

func (i WindowsAutopilotSyncStatus) String() string {
    return []string{"unknown", "inProgress", "completed", "failed"}[i]
}
func ParseWindowsAutopilotSyncStatus(v string) (any, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS
        case "inProgress":
            result = INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS
        case "completed":
            result = COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS
        case "failed":
            result = FAILED_WINDOWSAUTOPILOTSYNCSTATUS
        default:
            return nil, nil
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
func (i WindowsAutopilotSyncStatus) isMultiValue() bool {
    return false
}
