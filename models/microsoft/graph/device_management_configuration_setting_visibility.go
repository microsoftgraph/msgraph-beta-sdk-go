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
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "SETTINGSCATALOG":
            result = SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "TEMPLATE":
            result = TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSettingVisibility value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingVisibility(values []DeviceManagementConfigurationSettingVisibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
