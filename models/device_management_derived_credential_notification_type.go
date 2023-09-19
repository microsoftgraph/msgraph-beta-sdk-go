package models
import (
    "errors"
    "strings"
)
// Supported values for the notification type to use.
type DeviceManagementDerivedCredentialNotificationType int

const (
    // None
    NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE DeviceManagementDerivedCredentialNotificationType = iota
    // Company Portal
    COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
    // Email
    EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
)

func (i DeviceManagementDerivedCredentialNotificationType) String() string {
    var values []string
    for p := DeviceManagementDerivedCredentialNotificationType(1); p <= EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "companyPortal", "email"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementDerivedCredentialNotificationType(v string) (any, error) {
    var result DeviceManagementDerivedCredentialNotificationType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
            case "companyPortal":
                result |= COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
            case "email":
                result |= EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
            default:
                return 0, errors.New("Unknown DeviceManagementDerivedCredentialNotificationType value: " + v)
        }
    }
    return &result, nil
}
func SerializeDeviceManagementDerivedCredentialNotificationType(values []DeviceManagementDerivedCredentialNotificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementDerivedCredentialNotificationType) isMultiValue() bool {
    return true
}
