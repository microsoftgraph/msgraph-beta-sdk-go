package models
// Possible values for defenderSecurityCenterNotificationsFromApp
type DefenderSecurityCenterNotificationsFromAppType int

const (
    // Not Configured
    NOTCONFIGURED_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE DefenderSecurityCenterNotificationsFromAppType = iota
    // Block non-critical notifications
    BLOCKNONCRITICALNOTIFICATIONS_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE
    // Block all notifications
    BLOCKALLNOTIFICATIONS_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE
)

func (i DefenderSecurityCenterNotificationsFromAppType) String() string {
    return []string{"notConfigured", "blockNoncriticalNotifications", "blockAllNotifications"}[i]
}
func ParseDefenderSecurityCenterNotificationsFromAppType(v string) (any, error) {
    result := NOTCONFIGURED_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE
        case "blockNoncriticalNotifications":
            result = BLOCKNONCRITICALNOTIFICATIONS_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE
        case "blockAllNotifications":
            result = BLOCKALLNOTIFICATIONS_DEFENDERSECURITYCENTERNOTIFICATIONSFROMAPPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDefenderSecurityCenterNotificationsFromAppType(values []DefenderSecurityCenterNotificationsFromAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefenderSecurityCenterNotificationsFromAppType) isMultiValue() bool {
    return false
}
