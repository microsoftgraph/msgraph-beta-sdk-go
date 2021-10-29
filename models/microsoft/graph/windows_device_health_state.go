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
    switch strings.ToUpper(v) {
        case "CLEAN":
            return CLEAN_WINDOWSDEVICEHEALTHSTATE, nil
        case "FULLSCANPENDING":
            return FULLSCANPENDING_WINDOWSDEVICEHEALTHSTATE, nil
        case "REBOOTPENDING":
            return REBOOTPENDING_WINDOWSDEVICEHEALTHSTATE, nil
        case "MANUALSTEPSPENDING":
            return MANUALSTEPSPENDING_WINDOWSDEVICEHEALTHSTATE, nil
        case "OFFLINESCANPENDING":
            return OFFLINESCANPENDING_WINDOWSDEVICEHEALTHSTATE, nil
        case "CRITICAL":
            return CRITICAL_WINDOWSDEVICEHEALTHSTATE, nil
    }
    return 0, errors.New("Unknown WindowsDeviceHealthState value: " + v)
}
func SerializeWindowsDeviceHealthState(values []WindowsDeviceHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
