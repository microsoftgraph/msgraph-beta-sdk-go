package graph
import (
    "strings"
    "errors"
)
type DeviceManagementConfigurationSettingUsage int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE DeviceManagementConfigurationSettingUsage = iota
    CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
)

func (i DeviceManagementConfigurationSettingUsage) String() string {
    return []string{"NONE", "CONFIGURATION"}[i]
}
func ParseDeviceManagementConfigurationSettingUsage(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE, nil
        case "CONFIGURATION":
            return CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationSettingUsage value: " + v)
}
func SerializeDeviceManagementConfigurationSettingUsage(values []DeviceManagementConfigurationSettingUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
