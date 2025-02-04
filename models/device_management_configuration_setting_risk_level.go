package models
import (
    "math"
    "strings"
)
// Setting RiskLevel
type DeviceManagementConfigurationSettingRiskLevel int

const (
    // Default. Low Risk Setting
    LOW_DEVICEMANAGEMENTCONFIGURATIONSETTINGRISKLEVEL = 1
    // Medium Risk Setting
    MEDIUM_DEVICEMANAGEMENTCONFIGURATIONSETTINGRISKLEVEL = 2
    // High Risk Setting
    HIGH_DEVICEMANAGEMENTCONFIGURATIONSETTINGRISKLEVEL = 4
)

func (i DeviceManagementConfigurationSettingRiskLevel) String() string {
    var values []string
    options := []string{"low", "medium", "high"}
    for p := 0; p < 3; p++ {
        mantis := DeviceManagementConfigurationSettingRiskLevel(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementConfigurationSettingRiskLevel(v string) (any, error) {
    var result DeviceManagementConfigurationSettingRiskLevel
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "low":
                result |= LOW_DEVICEMANAGEMENTCONFIGURATIONSETTINGRISKLEVEL
            case "medium":
                result |= MEDIUM_DEVICEMANAGEMENTCONFIGURATIONSETTINGRISKLEVEL
            case "high":
                result |= HIGH_DEVICEMANAGEMENTCONFIGURATIONSETTINGRISKLEVEL
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingRiskLevel(values []DeviceManagementConfigurationSettingRiskLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementConfigurationSettingRiskLevel) isMultiValue() bool {
    return true
}
