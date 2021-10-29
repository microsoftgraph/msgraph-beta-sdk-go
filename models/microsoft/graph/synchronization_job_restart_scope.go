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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
        case "CONNECTORDATASTORE":
            return CONNECTORDATASTORE_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
        case "ESCROWS":
            return ESCROWS_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
        case "WATERMARK":
            return WATERMARK_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
        case "QUARANTINESTATE":
            return QUARANTINESTATE_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
        case "FULL":
            return FULL_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
        case "FORCEDELETES":
            return FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE, nil
    }
    return 0, errors.New("Unknown SynchronizationJobRestartScope value: " + v)
}
func SerializeSynchronizationJobRestartScope(values []SynchronizationJobRestartScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
