package windowsupdates
import (
    "errors"
)
type WindowsReleaseHealthStatus int

const (
    RESOLVED_WINDOWSRELEASEHEALTHSTATUS WindowsReleaseHealthStatus = iota
    MITIGATEDEXTERNAL_WINDOWSRELEASEHEALTHSTATUS
    MITIGATED_WINDOWSRELEASEHEALTHSTATUS
    RESOLVEDEXTERNAL_WINDOWSRELEASEHEALTHSTATUS
    CONFIRMED_WINDOWSRELEASEHEALTHSTATUS
    REPORTED_WINDOWSRELEASEHEALTHSTATUS
    INVESTIGATING_WINDOWSRELEASEHEALTHSTATUS
    UNKNOWNFUTUREVALUE_WINDOWSRELEASEHEALTHSTATUS
)

func (i WindowsReleaseHealthStatus) String() string {
    return []string{"resolved", "mitigatedExternal", "mitigated", "resolvedExternal", "confirmed", "reported", "investigating", "unknownFutureValue"}[i]
}
func ParseWindowsReleaseHealthStatus(v string) (any, error) {
    result := RESOLVED_WINDOWSRELEASEHEALTHSTATUS
    switch v {
        case "resolved":
            result = RESOLVED_WINDOWSRELEASEHEALTHSTATUS
        case "mitigatedExternal":
            result = MITIGATEDEXTERNAL_WINDOWSRELEASEHEALTHSTATUS
        case "mitigated":
            result = MITIGATED_WINDOWSRELEASEHEALTHSTATUS
        case "resolvedExternal":
            result = RESOLVEDEXTERNAL_WINDOWSRELEASEHEALTHSTATUS
        case "confirmed":
            result = CONFIRMED_WINDOWSRELEASEHEALTHSTATUS
        case "reported":
            result = REPORTED_WINDOWSRELEASEHEALTHSTATUS
        case "investigating":
            result = INVESTIGATING_WINDOWSRELEASEHEALTHSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSRELEASEHEALTHSTATUS
        default:
            return 0, errors.New("Unknown WindowsReleaseHealthStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsReleaseHealthStatus(values []WindowsReleaseHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsReleaseHealthStatus) isMultiValue() bool {
    return false
}
