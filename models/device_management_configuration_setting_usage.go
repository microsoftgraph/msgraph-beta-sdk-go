package models
import (
    "errors"
    "math"
    "strings"
)
// Supported setting types
type DeviceManagementConfigurationSettingUsage int

const (
    // Default. No setting type specified.
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE = 1
    // Configuration setting type.
    CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE = 2
    // Compliance setting type.
    COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE = 4
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE = 8
)

func (i DeviceManagementConfigurationSettingUsage) String() string {
    var values []string
    options := []string{"none", "configuration", "compliance", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := DeviceManagementConfigurationSettingUsage(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
