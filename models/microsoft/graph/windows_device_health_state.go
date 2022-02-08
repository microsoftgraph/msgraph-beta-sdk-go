package graph
import (
    "strings"
    "errors"
)
// 
type WindowsDeviceHealthState int

const (
    CLEAN_WINDOWSDEVICEHEALTHSTATE WindowsDeviceHealthState = iota
    FULLSCANPENDING_WINDOWSDEVICEHEALTHSTATE
    REBOOTPENDING_WINDOWSDEVICEHEALTHSTATE
    MANUALSTEPSPENDING_WINDOWSDEVICEHEALTHSTATE
    OFFLINESCANPENDING_WINDOWSDEVICEHEALTHSTATE
    CRITICAL_WINDOWSDEVICEHEALTHSTATE
)

func (i WindowsDeviceHealthState) String() string {
    return []string{"CLEAN", "FULLSCANPENDING", "REBOOTPENDING", "MANUALSTEPSPENDING", "OFFLINESCANPENDING", "CRITICAL"}[i]
}
func ParseWindowsDeviceHealthState(v string) (interface{}, error) {
    result := CLEAN_WINDOWSDEVICEHEALTHSTATE
    switch strings.ToUpper(v) {
        case "CLEAN":
            result = CLEAN_WINDOWSDEVICEHEALTHSTATE
        case "FULLSCANPENDING":
            result = FULLSCANPENDING_WINDOWSDEVICEHEALTHSTATE
        case "REBOOTPENDING":
            result = REBOOTPENDING_WINDOWSDEVICEHEALTHSTATE
        case "MANUALSTEPSPENDING":
            result = MANUALSTEPSPENDING_WINDOWSDEVICEHEALTHSTATE
        case "OFFLINESCANPENDING":
            result = OFFLINESCANPENDING_WINDOWSDEVICEHEALTHSTATE
        case "CRITICAL":
            result = CRITICAL_WINDOWSDEVICEHEALTHSTATE
        default:
            return 0, errors.New("Unknown WindowsDeviceHealthState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDeviceHealthState(values []WindowsDeviceHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
