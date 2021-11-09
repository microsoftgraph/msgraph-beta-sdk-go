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
    switch strings.ToUpper(v) {
        case "NOACTION":
            return NOACTION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "NOTIFICATION":
            return NOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "BLOCK":
            return BLOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "RETIRE":
            return RETIRE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "WIPE":
            return WIPE_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "REMOVERESOURCEACCESSPROFILES":
            return REMOVERESOURCEACCESSPROFILES_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "PUSHNOTIFICATION":
            return PUSHNOTIFICATION_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
        case "REMOTELOCK":
            return REMOTELOCK_DEVICEMANAGEMENTCOMPLIANCEACTIONTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementComplianceActionType value: " + v)
}
func SerializeDeviceManagementComplianceActionType(values []DeviceManagementComplianceActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
