package models
import (
    "math"
    "strings"
)
// Supported setting types
type DeviceManagementConfigurationSettingVisibility int

const (
    // Default. Not visible.
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY = 1
    // Visible to setting catalog policy type.
    SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY = 2
    // Visible to template policy type.
    TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY = 4
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY = 8
)

func (i DeviceManagementConfigurationSettingVisibility) String() string {
    var values []string
    options := []string{"none", "settingsCatalog", "template", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := DeviceManagementConfigurationSettingVisibility(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
