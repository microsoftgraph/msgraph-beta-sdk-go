package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := DeviceManagementConfigurationSettingVisibility(1); p <= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "settingsCatalog", "template", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementConfigurationSettingVisibility(v string) (any, error) {
    var result DeviceManagementConfigurationSettingVisibility
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
            case "settingsCatalog":
                result |= SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
            case "template":
                result |= TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
            default:
                return 0, errors.New("Unknown DeviceManagementConfigurationSettingVisibility value: " + v)
        }
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
func (i DeviceManagementConfigurationSettingVisibility) isMultiValue() bool {
    return true
}
