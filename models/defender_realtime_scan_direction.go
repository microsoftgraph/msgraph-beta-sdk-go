package models
import (
    "errors"
)
// Possible values for monitoring file activity.
type DefenderRealtimeScanDirection int

const (
    // 0 (default) – Monitor all files(bi-directional)
    MONITORALLFILES_DEFENDERREALTIMESCANDIRECTION DefenderRealtimeScanDirection = iota
    // Monitor incoming files only.
    MONITORINCOMINGFILESONLY_DEFENDERREALTIMESCANDIRECTION
    // Monitor outgoing files only.
    MONITOROUTGOINGFILESONLY_DEFENDERREALTIMESCANDIRECTION
)

func (i DefenderRealtimeScanDirection) String() string {
    return []string{"monitorAllFiles", "monitorIncomingFilesOnly", "monitorOutgoingFilesOnly"}[i]
}
func ParseDefenderRealtimeScanDirection(v string) (any, error) {
    result := MONITORALLFILES_DEFENDERREALTIMESCANDIRECTION
    switch v {
        case "monitorAllFiles":
            result = MONITORALLFILES_DEFENDERREALTIMESCANDIRECTION
        case "monitorIncomingFilesOnly":
            result = MONITORINCOMINGFILESONLY_DEFENDERREALTIMESCANDIRECTION
        case "monitorOutgoingFilesOnly":
            result = MONITOROUTGOINGFILESONLY_DEFENDERREALTIMESCANDIRECTION
        default:
            return 0, errors.New("Unknown DefenderRealtimeScanDirection value: " + v)
    }
    return &result, nil
}
func SerializeDefenderRealtimeScanDirection(values []DefenderRealtimeScanDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefenderRealtimeScanDirection) isMultiValue() bool {
    return false
}
