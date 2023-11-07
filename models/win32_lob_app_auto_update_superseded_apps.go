package models
import (
    "errors"
)
// Contains value for auto-update superseded apps.
type Win32LobAppAutoUpdateSupersededApps int

const (
    // Indicates that the auto-update superseded apps is not configured and the app will not auto-update the superseded apps.
    NOTCONFIGURED_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS Win32LobAppAutoUpdateSupersededApps = iota
    // Indicates that the auto-update superseded apps is enabled and the app will auto-update the superseded apps if the superseded apps are installed on the device.
    ENABLED_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS
)

func (i Win32LobAppAutoUpdateSupersededApps) String() string {
    return []string{"notConfigured", "enabled", "unknownFutureValue"}[i]
}
func ParseWin32LobAppAutoUpdateSupersededApps(v string) (any, error) {
    result := NOTCONFIGURED_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS
        case "enabled":
            result = ENABLED_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WIN32LOBAPPAUTOUPDATESUPERSEDEDAPPS
        default:
            return 0, errors.New("Unknown Win32LobAppAutoUpdateSupersededApps value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppAutoUpdateSupersededApps(values []Win32LobAppAutoUpdateSupersededApps) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Win32LobAppAutoUpdateSupersededApps) isMultiValue() bool {
    return false
}
