package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationSettingVisibility int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY DeviceManagementConfigurationSettingVisibility = iota
    SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
)

func (i DeviceManagementConfigurationSettingVisibility) String() string {
    return []string{"NONE", "SETTINGSCATALOG", "TEMPLATE"}[i]
}
func ParseDeviceManagementConfigurationSettingVisibility(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY, nil
        case "SETTINGSCATALOG":
            return SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY, nil
        case "TEMPLATE":
            return TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationSettingVisibility value: " + v)
}
func SerializeDeviceManagementConfigurationSettingVisibility(values []DeviceManagementConfigurationSettingVisibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
