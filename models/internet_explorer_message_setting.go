package models
// What message will be displayed by Edge before switching to Internet Explorer.
type InternetExplorerMessageSetting int

const (
    // Not configured.
    NOTCONFIGURED_INTERNETEXPLORERMESSAGESETTING InternetExplorerMessageSetting = iota
    // Disabled.
    DISABLED_INTERNETEXPLORERMESSAGESETTING
    // Enabled.
    ENABLED_INTERNETEXPLORERMESSAGESETTING
    // KeepGoing.
    KEEPGOING_INTERNETEXPLORERMESSAGESETTING
)

func (i InternetExplorerMessageSetting) String() string {
    return []string{"notConfigured", "disabled", "enabled", "keepGoing"}[i]
}
func ParseInternetExplorerMessageSetting(v string) (any, error) {
    result := NOTCONFIGURED_INTERNETEXPLORERMESSAGESETTING
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_INTERNETEXPLORERMESSAGESETTING
        case "disabled":
            result = DISABLED_INTERNETEXPLORERMESSAGESETTING
        case "enabled":
            result = ENABLED_INTERNETEXPLORERMESSAGESETTING
        case "keepGoing":
            result = KEEPGOING_INTERNETEXPLORERMESSAGESETTING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeInternetExplorerMessageSetting(values []InternetExplorerMessageSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InternetExplorerMessageSetting) isMultiValue() bool {
    return false
}
