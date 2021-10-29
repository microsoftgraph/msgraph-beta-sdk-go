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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE, nil
        case "COMPANYPORTAL":
            return COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE, nil
        case "EMAIL":
            return EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementDerivedCredentialNotificationType value: " + v)
}
func SerializeDeviceManagementDerivedCredentialNotificationType(values []DeviceManagementDerivedCredentialNotificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
