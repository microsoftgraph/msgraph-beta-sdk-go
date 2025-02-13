package models
import (
    "math"
    "strings"
)
type AppDevelopmentPlatforms int

const (
    DEVELOPERPORTAL_APPDEVELOPMENTPLATFORMS = 1
    UNKNOWNFUTUREVALUE_APPDEVELOPMENTPLATFORMS = 2
)

func (i AppDevelopmentPlatforms) String() string {
    var values []string
    options := []string{"developerPortal", "unknownFutureValue"}
    for p := 0; p < 2; p++ {
        mantis := AppDevelopmentPlatforms(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAppDevelopmentPlatforms(v string) (any, error) {
    var result AppDevelopmentPlatforms
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "developerPortal":
                result |= DEVELOPERPORTAL_APPDEVELOPMENTPLATFORMS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_APPDEVELOPMENTPLATFORMS
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeAppDevelopmentPlatforms(values []AppDevelopmentPlatforms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppDevelopmentPlatforms) isMultiValue() bool {
    return true
}
