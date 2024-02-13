package models
import (
    "errors"
    "math"
    "strings"
)
// Supported platform types.
type DeviceManagementConfigurationPlatforms int

const (
    // Default. No platform type specified.
    NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 1
    // Settings for Android platform.
    ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 2
    // Settings for iOS platform.
    IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 4
    // Settings for MacOS platform.
    MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 8
    // Windows 10 X.
    WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 16
    // Settings for Windows 10 platform.
    WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 32
    // Settings for Linux platform.
    LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 64
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 128
)

func (i DeviceManagementConfigurationPlatforms) String() string {
    var values []string
    options := []string{"none", "android", "iOS", "macOS", "windows10X", "windows10", "linux", "unknownFutureValue"}
    for p := 0; p < 8; p++ {
        mantis := DeviceManagementConfigurationPlatforms(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
