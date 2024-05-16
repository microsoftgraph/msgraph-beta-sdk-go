package models
import (
    "math"
    "strings"
)
// Supported platform types.
type DeviceManagementConfigurationPlatforms int

const (
    // Indicates the settings contained in this configuration don't have a platform set.
    NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 1
    // Indicates that the settings contained in associated configuration applies to the Android operating system. 
    ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 2
    // Indicates that the settings contained in associated configuration applies to the iOS operating system.
    IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 4
    // Indicates that the settings contained in associated configuration applies to the MacOS operating system.
    MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 8
    // Indicates that the settings contained in associated configuration applies to the Windows 10X operating system.
    WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 16
    // Indicates that the settings contained in associated configuration applies to the Windows 10 operating system.
    WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS = 32
    // Indicates that the settings contained in associated configuration applies to the Linux operating system.
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
                return nil, nil
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
