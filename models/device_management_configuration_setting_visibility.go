package models
import (
    "errors"
)
// Supported setting types
type DeviceManagementConfigurationSettingVisibility int

const (
    // Default. Not visible.
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY DeviceManagementConfigurationSettingVisibility = iota
    // Visible to setting catalog policy type.
    SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    // Visible to template policy type.
    TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
)

func (i DeviceManagementConfigurationSettingVisibility) String() string {
    return []string{"none", "settingsCatalog", "template", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationSettingVisibility(v string) (any, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "settingsCatalog":
            result = SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "template":
            result = TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
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
