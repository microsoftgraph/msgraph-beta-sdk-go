package models
import (
    "errors"
    "math"
    "strings"
)
// Device health monitoring scope
type WindowsHealthMonitoringScope int

const (
    // Undefined
    UNDEFINED_WINDOWSHEALTHMONITORINGSCOPE = 1
    // Basic events for windows device health monitoring
    HEALTHMONITORING_WINDOWSHEALTHMONITORINGSCOPE = 2
    // Boot performance events
    BOOTPERFORMANCE_WINDOWSHEALTHMONITORINGSCOPE = 4
    // Windows updates events
    WINDOWSUPDATES_WINDOWSHEALTHMONITORINGSCOPE = 8
    // PrivilegeManagement
    PRIVILEGEMANAGEMENT_WINDOWSHEALTHMONITORINGSCOPE = 16
)

func (i WindowsHealthMonitoringScope) String() string {
    var values []string
    options := []string{"undefined", "healthMonitoring", "bootPerformance", "windowsUpdates", "privilegeManagement"}
    for p := 0; p < 5; p++ {
        mantis := WindowsHealthMonitoringScope(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsHealthMonitoringScope(v string) (any, error) {
    var result WindowsHealthMonitoringScope
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "undefined":
                result |= UNDEFINED_WINDOWSHEALTHMONITORINGSCOPE
            case "healthMonitoring":
                result |= HEALTHMONITORING_WINDOWSHEALTHMONITORINGSCOPE
            case "bootPerformance":
                result |= BOOTPERFORMANCE_WINDOWSHEALTHMONITORINGSCOPE
            case "windowsUpdates":
                result |= WINDOWSUPDATES_WINDOWSHEALTHMONITORINGSCOPE
            case "privilegeManagement":
                result |= PRIVILEGEMANAGEMENT_WINDOWSHEALTHMONITORINGSCOPE
            default:
                return 0, errors.New("Unknown WindowsHealthMonitoringScope value: " + v)
        }
    }
    return &result, nil
}
func SerializeWindowsHealthMonitoringScope(values []WindowsHealthMonitoringScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsHealthMonitoringScope) isMultiValue() bool {
    return true
}
