package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementDerivedCredentialNotificationType int

const (
    NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE DeviceManagementDerivedCredentialNotificationType = iota
    COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
    EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
)

func (i DeviceManagementDerivedCredentialNotificationType) String() string {
    return []string{"NONE", "COMPANYPORTAL", "EMAIL"}[i]
}
func ParseDeviceManagementDerivedCredentialNotificationType(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
        case "COMPANYPORTAL":
            result = COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
        case "EMAIL":
            result = EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementDerivedCredentialNotificationType value: " + v)
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
