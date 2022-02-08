package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationSettingUsage int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE DeviceManagementConfigurationSettingUsage = iota
    CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
)

func (i DeviceManagementConfigurationSettingUsage) String() string {
    return []string{"NONE", "CONFIGURATION", "COMPLIANCE"}[i]
}
func ParseDeviceManagementConfigurationSettingUsage(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "CONFIGURATION":
            result = CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "COMPLIANCE":
            result = COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
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
