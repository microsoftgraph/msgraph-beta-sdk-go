package models
// Scheduled Action Type Enum
type DeviceManagementComplianceActionType int

const (
    // No Action
    NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE DeviceManagementComplianceActionType = iota
    // Send Notification
    NOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    // Block the device in AAD
    BLOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    // Retire the device
    RETIRE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    // Wipe the device
    WIPE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    // Remove Resource Access Profiles from the device
    REMOVERESOURCEACCESSPROFILES_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    // Send push notification to device
    PUSHNOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    // Remotely lock the device
    REMOTELOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
)

func (i DeviceManagementComplianceActionType) String() string {
    return []string{"noAction", "notification", "block", "retire", "wipe", "removeResourceAccessProfiles", "pushNotification", "remoteLock"}[i]
}
func ParseDeviceManagementComplianceActionType(v string) (any, error) {
    result := NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    switch v {
        case "noAction":
            result = NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "notification":
            result = NOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "block":
            result = BLOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "retire":
            result = RETIRE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "wipe":
            result = WIPE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "removeResourceAccessProfiles":
            result = REMOVERESOURCEACCESSPROFILES_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "pushNotification":
            result = PUSHNOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "remoteLock":
            result = REMOTELOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceManagementComplianceActionType(values []DeviceManagementComplianceActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementComplianceActionType) isMultiValue() bool {
    return false
}
