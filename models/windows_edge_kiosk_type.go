package models
import (
    "errors"
)
// Edge kiosk type
type WindowsEdgeKioskType int

const (
    // public-browsing
    PUBLICBROWSING_WINDOWSEDGEKIOSKTYPE WindowsEdgeKioskType = iota
    // fullscreen
    FULLSCREEN_WINDOWSEDGEKIOSKTYPE
)

func (i WindowsEdgeKioskType) String() string {
    return []string{"publicBrowsing", "fullScreen"}[i]
}
func ParseWindowsEdgeKioskType(v string) (any, error) {
    result := PUBLICBROWSING_WINDOWSEDGEKIOSKTYPE
    switch v {
        case "publicBrowsing":
            result = PUBLICBROWSING_WINDOWSEDGEKIOSKTYPE
        case "fullScreen":
            result = FULLSCREEN_WINDOWSEDGEKIOSKTYPE
        default:
            return 0, errors.New("Unknown WindowsEdgeKioskType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsEdgeKioskType(values []WindowsEdgeKioskType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
