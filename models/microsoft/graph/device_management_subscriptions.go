package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTSUBSCRIPTIONS, nil
        case "INTUNE":
            return INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS, nil
        case "OFFICE365":
            return OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS, nil
        case "INTUNEPREMIUM":
            return INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS, nil
        case "INTUNE_EDU":
            return INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS, nil
        case "INTUNE_SMB":
            return INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS, nil
    }
    return 0, errors.New("Unknown DeviceManagementSubscriptions value: " + v)
}
func SerializeDeviceManagementSubscriptions(values []DeviceManagementSubscriptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
