package models
// Platform Type
type PlatformType int

const (
    // None
    NONE_PLATFORMTYPE PlatformType = iota
    // Android
    ANDROID_PLATFORMTYPE
    // iOS
    IOS_PLATFORMTYPE
    // MacOS
    MACOS_PLATFORMTYPE
    // Windows 10X Platform Type
    WINDOWS10X_PLATFORMTYPE
    // Windows 10
    WINDOWS10_PLATFORMTYPE
    // Linux
    LINUX_PLATFORMTYPE
    // UnknownFutureValue
    UNKNOWNFUTUREVALUE_PLATFORMTYPE
)

func (i PlatformType) String() string {
    return []string{"none", "android", "iOS", "macOS", "windows10X", "windows10", "linux", "unknownFutureValue"}[i]
}
func ParsePlatformType(v string) (any, error) {
    result := NONE_PLATFORMTYPE
    switch v {
        case "none":
            result = NONE_PLATFORMTYPE
        case "android":
            result = ANDROID_PLATFORMTYPE
        case "iOS":
            result = IOS_PLATFORMTYPE
        case "macOS":
            result = MACOS_PLATFORMTYPE
        case "windows10X":
            result = WINDOWS10X_PLATFORMTYPE
        case "windows10":
            result = WINDOWS10_PLATFORMTYPE
        case "linux":
            result = LINUX_PLATFORMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLATFORMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePlatformType(values []PlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlatformType) isMultiValue() bool {
    return false
}
