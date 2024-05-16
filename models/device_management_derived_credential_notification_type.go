package models
import (
    "math"
    "strings"
)
// Supported values for the notification type to use.
type DeviceManagementDerivedCredentialNotificationType int

const (
    // None
    NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE = 1
    // Company Portal
    COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE = 2
    // Email
    EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE = 4
)

func (i DeviceManagementDerivedCredentialNotificationType) String() string {
    var values []string
    options := []string{"none", "companyPortal", "email"}
    for p := 0; p < 3; p++ {
        mantis := DeviceManagementDerivedCredentialNotificationType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
