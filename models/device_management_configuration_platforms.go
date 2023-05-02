package models
import (
    "errors"
)
// Supported platform types.
type DeviceManagementConfigurationPlatforms int

const (
    // Default. No platform type specified.
    NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS DeviceManagementConfigurationPlatforms = iota
    // Settings for Android platform.
    ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Settings for iOS platform.
    IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Settings for MacOS platform.
    MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Windows 10 X.
    WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Settings for Windows 10 platform.
    WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Settings for Linux platform.
    LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
)

func (i DeviceManagementConfigurationPlatforms) String() string {
    return []string{"none", "android", "iOS", "macOS", "windows10X", "windows10", "linux", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationPlatforms(v string) (any, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "android":
            result = ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "iOS":
            result = IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "macOS":
            result = MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "windows10X":
            result = WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "windows10":
            result = WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "linux":
            result = LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "unknownFutureValue":
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
