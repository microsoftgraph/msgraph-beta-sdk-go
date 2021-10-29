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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PLATFORM, nil
        case "IOS":
            return IOS_PLATFORM, nil
        case "ANDROID":
            return ANDROID_PLATFORM, nil
        case "WINDOWS":
            return WINDOWS_PLATFORM, nil
        case "WINDOWSMOBILE":
            return WINDOWSMOBILE_PLATFORM, nil
        case "MACOS":
            return MACOS_PLATFORM, nil
    }
    return 0, errors.New("Unknown Platform value: " + v)
}
func SerializePlatform(values []Platform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
