package models
// Supported platform types for policies.
type Platform int

const (
    // Default.Indicates the managed device is not known and is associated with 'Unknown' device platform.
    UNKNOWN_PLATFORM Platform = iota
    // Indicates the managed device is Apple device that runs on iOS operation system.
    IOS_PLATFORM
    // Indicates the managed device is a Android device that runs on Android operation system. 
    ANDROID_PLATFORM
    // Indicates the managed device is a Windows device that runs on Windows operation system.
    WINDOWS_PLATFORM
    // Indicates the managed device is a Windows-based mobile device that runs on Windows Mobile operation system.
    WINDOWSMOBILE_PLATFORM
    // Indicates the managed device is Apple device that runs on MacOS operation system.
    MACOS_PLATFORM
    // Indicates the managed device is Apple device that runs on VisionOS operation system.
    VISIONOS_PLATFORM
    // Indicates the managed device is Apple device that runs on tvOS operation system.
    TVOS_PLATFORM
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_PLATFORM
)

func (i Platform) String() string {
    return []string{"unknown", "ios", "android", "windows", "windowsMobile", "macOS", "visionOS", "tvOS", "unknownFutureValue"}[i]
}
func ParsePlatform(v string) (any, error) {
    result := UNKNOWN_PLATFORM
    switch v {
        case "unknown":
            result = UNKNOWN_PLATFORM
        case "ios":
            result = IOS_PLATFORM
        case "android":
            result = ANDROID_PLATFORM
        case "windows":
            result = WINDOWS_PLATFORM
        case "windowsMobile":
            result = WINDOWSMOBILE_PLATFORM
        case "macOS":
            result = MACOS_PLATFORM
        case "visionOS":
            result = VISIONOS_PLATFORM
        case "tvOS":
            result = TVOS_PLATFORM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLATFORM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePlatform(values []Platform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Platform) isMultiValue() bool {
    return false
}
