package models
import (
    "errors"
)
// Allow(Not Configured) or prevent(Block) the syncing of Microsoft Edge Browser settings. Option to prevent syncing across devices, but allow user override.
type BrowserSyncSetting int

const (
    // Default – Allow syncing of browser settings across devices.
    NOTCONFIGURED_BROWSERSYNCSETTING BrowserSyncSetting = iota
    // Prevent syncing of browser settings across user devices, allow user override of setting.
    BLOCKEDWITHUSEROVERRIDE_BROWSERSYNCSETTING
    // Absolutely prevent syncing of browser settings across user devices.
    BLOCKED_BROWSERSYNCSETTING
)

func (i BrowserSyncSetting) String() string {
    return []string{"notConfigured", "blockedWithUserOverride", "blocked"}[i]
}
func ParseBrowserSyncSetting(v string) (any, error) {
    result := NOTCONFIGURED_BROWSERSYNCSETTING
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_BROWSERSYNCSETTING
        case "blockedWithUserOverride":
            result = BLOCKEDWITHUSEROVERRIDE_BROWSERSYNCSETTING
        case "blocked":
            result = BLOCKED_BROWSERSYNCSETTING
        default:
            return 0, errors.New("Unknown BrowserSyncSetting value: " + v)
    }
    return &result, nil
}
func SerializeBrowserSyncSetting(values []BrowserSyncSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BrowserSyncSetting) isMultiValue() bool {
    return false
}
