package models
import (
    "errors"
)
// Supported setting types
type DeviceManagementConfigurationSettingUsage int

const (
    // Default. No setting type specified.
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE DeviceManagementConfigurationSettingUsage = iota
    // Configuration setting type.
    CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    // Compliance setting type.
    COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
)

func (i DeviceManagementConfigurationSettingUsage) String() string {
    return []string{"none", "configuration", "compliance", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationSettingUsage(v string) (any, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "configuration":
            result = CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "compliance":
            result = COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSettingUsage value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingUsage(values []DeviceManagementConfigurationSettingUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
