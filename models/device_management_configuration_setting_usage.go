package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := DeviceManagementConfigurationSettingUsage(1); p <= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "configuration", "compliance", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementConfigurationSettingUsage(v string) (any, error) {
    var result DeviceManagementConfigurationSettingUsage
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
            case "configuration":
                result |= CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
            case "compliance":
                result |= COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
            default:
                return 0, errors.New("Unknown DeviceManagementConfigurationSettingUsage value: " + v)
        }
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
func (i DeviceManagementConfigurationSettingUsage) isMultiValue() bool {
    return true
}
