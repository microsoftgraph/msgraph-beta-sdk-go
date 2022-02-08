package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationPlatforms int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS DeviceManagementConfigurationPlatforms = iota
    ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
)

func (i DeviceManagementConfigurationPlatforms) String() string {
    return []string{"NONE", "ANDROID", "IOS", "MACOS", "WINDOWS10X", "WINDOWS10", "LINUX", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeviceManagementConfigurationPlatforms(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "ANDROID":
            result = ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "IOS":
            result = IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "MACOS":
            result = MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "WINDOWS10X":
            result = WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "WINDOWS10":
            result = WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "LINUX":
            result = LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationPlatforms value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationPlatforms(values []DeviceManagementConfigurationPlatforms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
