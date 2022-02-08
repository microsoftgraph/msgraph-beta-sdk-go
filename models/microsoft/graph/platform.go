package graph
import (
    "strings"
    "errors"
)
// 
type Platform int

const (
    UNKNOWN_PLATFORM Platform = iota
    IOS_PLATFORM
    ANDROID_PLATFORM
    WINDOWS_PLATFORM
    WINDOWSMOBILE_PLATFORM
    MACOS_PLATFORM
)

func (i Platform) String() string {
    return []string{"UNKNOWN", "IOS", "ANDROID", "WINDOWS", "WINDOWSMOBILE", "MACOS"}[i]
}
func ParsePlatform(v string) (interface{}, error) {
    result := UNKNOWN_PLATFORM
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PLATFORM
        case "IOS":
            result = IOS_PLATFORM
        case "ANDROID":
            result = ANDROID_PLATFORM
        case "WINDOWS":
            result = WINDOWS_PLATFORM
        case "WINDOWSMOBILE":
            result = WINDOWSMOBILE_PLATFORM
        case "MACOS":
            result = MACOS_PLATFORM
        default:
            return 0, errors.New("Unknown Platform value: " + v)
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
