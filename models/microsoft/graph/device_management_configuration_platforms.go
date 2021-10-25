package graph
import (
    "strings"
    "errors"
)
type DeviceManagementConfigurationPlatforms int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS DeviceManagementConfigurationPlatforms = iota
    MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
)

func (i DeviceManagementConfigurationPlatforms) String() string {
    return []string{"NONE", "MACOS", "WINDOWS10X", "WINDOWS10"}[i]
}
func ParseDeviceManagementConfigurationPlatforms(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS, nil
        case "MACOS":
            return MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS, nil
        case "WINDOWS10X":
            return WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS, nil
        case "WINDOWS10":
            return WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationPlatforms value: " + v)
}
func SerializeDeviceManagementConfigurationPlatforms(values []DeviceManagementConfigurationPlatforms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
