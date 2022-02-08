package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementComplianceActionType int

const (
    NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE DeviceManagementComplianceActionType = iota
    NOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    BLOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    RETIRE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    WIPE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    REMOVERESOURCEACCESSPROFILES_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    PUSHNOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    REMOTELOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
)

func (i DeviceManagementComplianceActionType) String() string {
    return []string{"NOACTION", "NOTIFICATION", "BLOCK", "RETIRE", "WIPE", "REMOVERESOURCEACCESSPROFILES", "PUSHNOTIFICATION", "REMOTELOCK"}[i]
}
func ParseDeviceManagementComplianceActionType(v string) (interface{}, error) {
    result := NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
    switch strings.ToUpper(v) {
        case "NOACTION":
            result = NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "NOTIFICATION":
            result = NOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "BLOCK":
            result = BLOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "RETIRE":
            result = RETIRE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "WIPE":
            result = WIPE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "REMOVERESOURCEACCESSPROFILES":
            result = REMOVERESOURCEACCESSPROFILES_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "PUSHNOTIFICATION":
            result = PUSHNOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        case "REMOTELOCK":
            result = REMOTELOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementComplianceActionType value: " + v)
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
