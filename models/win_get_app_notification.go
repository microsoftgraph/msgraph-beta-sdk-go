// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Contains value for notification status.
type WinGetAppNotification int

const (
    // Show all notifications.
    SHOWALL_WINGETAPPNOTIFICATION WinGetAppNotification = iota
    // Only show restart notification and suppress other notifications.
    SHOWREBOOT_WINGETAPPNOTIFICATION
    // Hide all notifications.
    HIDEALL_WINGETAPPNOTIFICATION
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINGETAPPNOTIFICATION
)

func (i WinGetAppNotification) String() string {
    return []string{"showAll", "showReboot", "hideAll", "unknownFutureValue"}[i]
}
func ParseWinGetAppNotification(v string) (any, error) {
    result := SHOWALL_WINGETAPPNOTIFICATION
    switch v {
        case "showAll":
            result = SHOWALL_WINGETAPPNOTIFICATION
        case "showReboot":
            result = SHOWREBOOT_WINGETAPPNOTIFICATION
        case "hideAll":
            result = HIDEALL_WINGETAPPNOTIFICATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINGETAPPNOTIFICATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWinGetAppNotification(values []WinGetAppNotification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WinGetAppNotification) isMultiValue() bool {
    return false
}
