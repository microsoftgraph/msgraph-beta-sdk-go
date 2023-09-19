package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := DeviceManagementConfigurationPlatforms(1); p <= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "android", "iOS", "macOS", "windows10X", "windows10", "linux", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementConfigurationPlatforms(v string) (any, error) {
    var result DeviceManagementConfigurationPlatforms
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "android":
                result |= ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "iOS":
                result |= IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "macOS":
                result |= MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "windows10X":
                result |= WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "windows10":
                result |= WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "linux":
                result |= LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
            default:
                return 0, errors.New("Unknown DeviceManagementConfigurationPlatforms value: " + v)
        }
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
func (i DeviceManagementConfigurationPlatforms) isMultiValue() bool {
    return true
}
