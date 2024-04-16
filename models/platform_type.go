package models
import (
    "errors"
)
// Platform Type
type PlatformType int

const (
    // None Platform Type
    NONE_PLATFORMTYPE PlatformType = iota
    // Android Platform Type
    ANDROID_PLATFORMTYPE
    // Android Enterprise Platform Type
    ANDROIDENTERPRISE_PLATFORMTYPE
    // iOS Platform Type
    IOS_PLATFORMTYPE
    // MacOS Platform Type
    MACOS_PLATFORMTYPE
    // Windows 10X Platform Type
    WINDOWS10X_PLATFORMTYPE
    // Windows 10 Platform Type
    WINDOWS10_PLATFORMTYPE
)

func (i PlatformType) String() string {
    return []string{"none", "android", "androidEnterprise", "iOS", "macOS", "windows10X", "windows10"}[i]
}
func ParsePlatformType(v string) (any, error) {
    result := NONE_PLATFORMTYPE
    switch v {
        case "none":
            result = NONE_PLATFORMTYPE
        case "android":
            result = ANDROID_PLATFORMTYPE
        case "androidEnterprise":
            result = ANDROIDENTERPRISE_PLATFORMTYPE
        case "iOS":
            result = IOS_PLATFORMTYPE
        case "macOS":
            result = MACOS_PLATFORMTYPE
        case "windows10X":
            result = WINDOWS10X_PLATFORMTYPE
        case "windows10":
            result = WINDOWS10_PLATFORMTYPE
        default:
            return 0, errors.New("Unknown PlatformType value: " + v)
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
