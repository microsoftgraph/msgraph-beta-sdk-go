package graph
import (
    "strings"
    "errors"
)
// 
type SynchronizationJobRestartScope int

const (
    NONE_SYNCHRONIZATIONJOBRESTARTSCOPE SynchronizationJobRestartScope = iota
    CONNECTORDATASTORE_SYNCHRONIZATIONJOBRESTARTSCOPE
    ESCROWS_SYNCHRONIZATIONJOBRESTARTSCOPE
    WATERMARK_SYNCHRONIZATIONJOBRESTARTSCOPE
    QUARANTINESTATE_SYNCHRONIZATIONJOBRESTARTSCOPE
    FULL_SYNCHRONIZATIONJOBRESTARTSCOPE
    FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE
)

func (i SynchronizationJobRestartScope) String() string {
    return []string{"NONE", "CONNECTORDATASTORE", "ESCROWS", "WATERMARK", "QUARANTINESTATE", "FULL", "FORCEDELETES"}[i]
}
func ParseSynchronizationJobRestartScope(v string) (interface{}, error) {
    result := NONE_SYNCHRONIZATIONJOBRESTARTSCOPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "CONNECTORDATASTORE":
            result = CONNECTORDATASTORE_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "ESCROWS":
            result = ESCROWS_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "WATERMARK":
            result = WATERMARK_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "QUARANTINESTATE":
            result = QUARANTINESTATE_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "FULL":
            result = FULL_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "FORCEDELETES":
            result = FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE
        default:
            return 0, errors.New("Unknown SynchronizationJobRestartScope value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationJobRestartScope(values []SynchronizationJobRestartScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
