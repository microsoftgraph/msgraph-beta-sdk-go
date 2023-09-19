package models
import (
    "errors"
)
// Determines when notification previews are visible on an iOS device. Previews can include things like text (from Messages and Mail) and invitation details (from Calendar). When configured, it will override the user's defined preview settings.
type IosNotificationPreviewVisibility int

const (
    // Notification preview settings will not be overwritten.
    NOTCONFIGURED_IOSNOTIFICATIONPREVIEWVISIBILITY IosNotificationPreviewVisibility = iota
    // Always show notification previews.
    ALWAYSSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
    // Only show notification previews when the device is unlocked.
    HIDEWHENLOCKED_IOSNOTIFICATIONPREVIEWVISIBILITY
    // Never show notification previews.
    NEVERSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
)

func (i IosNotificationPreviewVisibility) String() string {
    return []string{"notConfigured", "alwaysShow", "hideWhenLocked", "neverShow"}[i]
}
func ParseIosNotificationPreviewVisibility(v string) (any, error) {
    result := NOTCONFIGURED_IOSNOTIFICATIONPREVIEWVISIBILITY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_IOSNOTIFICATIONPREVIEWVISIBILITY
        case "alwaysShow":
            result = ALWAYSSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
        case "hideWhenLocked":
            result = HIDEWHENLOCKED_IOSNOTIFICATIONPREVIEWVISIBILITY
        case "neverShow":
            result = NEVERSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
        default:
            return 0, errors.New("Unknown IosNotificationPreviewVisibility value: " + v)
    }
    return &result, nil
}
func SerializeIosNotificationPreviewVisibility(values []IosNotificationPreviewVisibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IosNotificationPreviewVisibility) isMultiValue() bool {
    return false
}
