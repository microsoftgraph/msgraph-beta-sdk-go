package models
import (
    "errors"
)
// Contains properties for Windows device type. Multiple values can be selected. Default value is `none`.
type WindowsDeviceType int

const (
    // No device types supported. Default value.
    NONE_WINDOWSDEVICETYPE WindowsDeviceType = iota
    // Indicates support for Desktop Windows device type.
    DESKTOP_WINDOWSDEVICETYPE
    // Indicates support for Mobile Windows device type.
    MOBILE_WINDOWSDEVICETYPE
    // Indicates support for Holographic Windows device type.
    HOLOGRAPHIC_WINDOWSDEVICETYPE
    // Indicates support for Team Windows device type.
    TEAM_WINDOWSDEVICETYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINDOWSDEVICETYPE
)

func (i WindowsDeviceType) String() string {
    return []string{"none", "desktop", "mobile", "holographic", "team", "unknownFutureValue"}[i]
}
func ParseWindowsDeviceType(v string) (any, error) {
    result := NONE_WINDOWSDEVICETYPE
    switch v {
        case "none":
            result = NONE_WINDOWSDEVICETYPE
        case "desktop":
            result = DESKTOP_WINDOWSDEVICETYPE
        case "mobile":
            result = MOBILE_WINDOWSDEVICETYPE
        case "holographic":
            result = HOLOGRAPHIC_WINDOWSDEVICETYPE
        case "team":
            result = TEAM_WINDOWSDEVICETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSDEVICETYPE
        default:
            return 0, errors.New("Unknown WindowsDeviceType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDeviceType(values []WindowsDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
