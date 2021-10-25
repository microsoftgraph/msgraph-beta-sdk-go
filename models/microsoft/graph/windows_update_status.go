package graph
import (
    "strings"
    "errors"
)
type WindowsUpdateStatus int

const (
    UPTODATE_WINDOWSUPDATESTATUS WindowsUpdateStatus = iota
    PENDINGINSTALLATION_WINDOWSUPDATESTATUS
    PENDINGREBOOT_WINDOWSUPDATESTATUS
    FAILED_WINDOWSUPDATESTATUS
)

func (i WindowsUpdateStatus) String() string {
    return []string{"UPTODATE", "PENDINGINSTALLATION", "PENDINGREBOOT", "FAILED"}[i]
}
func ParseWindowsUpdateStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UPTODATE":
            return UPTODATE_WINDOWSUPDATESTATUS, nil
        case "PENDINGINSTALLATION":
            return PENDINGINSTALLATION_WINDOWSUPDATESTATUS, nil
        case "PENDINGREBOOT":
            return PENDINGREBOOT_WINDOWSUPDATESTATUS, nil
        case "FAILED":
            return FAILED_WINDOWSUPDATESTATUS, nil
    }
    return 0, errors.New("Unknown WindowsUpdateStatus value: " + v)
}
func SerializeWindowsUpdateStatus(values []WindowsUpdateStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
