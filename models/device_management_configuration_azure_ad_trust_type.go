package models
// Required AAD Trust Type
type DeviceManagementConfigurationAzureAdTrustType int

const (
    // No AAD Trust Type specified
    NONE_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE DeviceManagementConfigurationAzureAdTrustType = iota
    // AAD Joined Trust Type
    AZUREADJOINED_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
    // AddWorkAccount
    ADDWORKACCOUNT_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
    // MDM only
    MDMONLY_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
)

func (i DeviceManagementConfigurationAzureAdTrustType) String() string {
    return []string{"none", "azureAdJoined", "addWorkAccount", "mdmOnly"}[i]
}
func ParseDeviceManagementConfigurationAzureAdTrustType(v string) (any, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        case "azureAdJoined":
            result = AZUREADJOINED_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        case "addWorkAccount":
            result = ADDWORKACCOUNT_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        case "mdmOnly":
            result = MDMONLY_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationAzureAdTrustType(values []DeviceManagementConfigurationAzureAdTrustType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementConfigurationAzureAdTrustType) isMultiValue() bool {
    return false
}
