package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementSubscriptions int

const (
    NONE_DEVICEMANAGEMENTSUBSCRIPTIONS DeviceManagementSubscriptions = iota
    INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS
    OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS
    INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS
    INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS
    INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS
)

func (i DeviceManagementSubscriptions) String() string {
    return []string{"NONE", "INTUNE", "OFFICE365", "INTUNEPREMIUM", "INTUNE_EDU", "INTUNE_SMB"}[i]
}
func ParseDeviceManagementSubscriptions(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTSUBSCRIPTIONS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "INTUNE":
            result = INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "OFFICE365":
            result = OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "INTUNEPREMIUM":
            result = INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "INTUNE_EDU":
            result = INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "INTUNE_SMB":
            result = INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS
        default:
            return 0, errors.New("Unknown DeviceManagementSubscriptions value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementSubscriptions(values []DeviceManagementSubscriptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
