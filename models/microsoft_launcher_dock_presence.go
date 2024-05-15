package models
// Microsoft Launcher Dock Presence selection.
type MicrosoftLauncherDockPresence int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_MICROSOFTLAUNCHERDOCKPRESENCE MicrosoftLauncherDockPresence = iota
    // Indicates the device's dock will be displayed on the device.
    SHOW_MICROSOFTLAUNCHERDOCKPRESENCE
    // Indicates the device's dock will be hidden on the device, but the user can access the dock by dragging the handler on the bottom of the screen.
    HIDE_MICROSOFTLAUNCHERDOCKPRESENCE
    // Indicates the device's dock will be disabled on the device.
    DISABLED_MICROSOFTLAUNCHERDOCKPRESENCE
)

func (i MicrosoftLauncherDockPresence) String() string {
    return []string{"notConfigured", "show", "hide", "disabled"}[i]
}
func ParseMicrosoftLauncherDockPresence(v string) (any, error) {
    result := NOTCONFIGURED_MICROSOFTLAUNCHERDOCKPRESENCE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MICROSOFTLAUNCHERDOCKPRESENCE
        case "show":
            result = SHOW_MICROSOFTLAUNCHERDOCKPRESENCE
        case "hide":
            result = HIDE_MICROSOFTLAUNCHERDOCKPRESENCE
        case "disabled":
            result = DISABLED_MICROSOFTLAUNCHERDOCKPRESENCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMicrosoftLauncherDockPresence(values []MicrosoftLauncherDockPresence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MicrosoftLauncherDockPresence) isMultiValue() bool {
    return false
}
