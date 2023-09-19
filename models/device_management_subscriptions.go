package models
import (
    "errors"
    "strings"
)
// Tenant mobile device management subscriptions.
type DeviceManagementSubscriptions int

const (
    // None
    NONE_DEVICEMANAGEMENTSUBSCRIPTIONS DeviceManagementSubscriptions = iota
    // Microsoft Intune Subscription
    INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Office365 Subscription
    OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Microsoft Intune Premium Subscription
    INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Microsoft Intune for Education Subscription
    INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Microsoft Intune for Small Businesses Subscription
    INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS
)

func (i DeviceManagementSubscriptions) String() string {
    var values []string
    for p := DeviceManagementSubscriptions(1); p <= INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "intune", "office365", "intunePremium", "intune_EDU", "intune_SMB"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementSubscriptions(v string) (any, error) {
    var result DeviceManagementSubscriptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTSUBSCRIPTIONS
            case "intune":
                result |= INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS
            case "office365":
                result |= OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS
            case "intunePremium":
                result |= INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS
            case "intune_EDU":
                result |= INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS
            case "intune_SMB":
                result |= INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS
            default:
                return 0, errors.New("Unknown DeviceManagementSubscriptions value: " + v)
        }
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
func (i DeviceManagementSubscriptions) isMultiValue() bool {
    return true
}
