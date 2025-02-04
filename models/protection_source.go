package models
import (
    "math"
    "strings"
)
type ProtectionSource int

const (
    NONE_PROTECTIONSOURCE = 1
    MANUAL_PROTECTIONSOURCE = 2
    DYNAMICRULE_PROTECTIONSOURCE = 4
    UNKNOWNFUTUREVALUE_PROTECTIONSOURCE = 8
)

func (i ProtectionSource) String() string {
    var values []string
    options := []string{"none", "manual", "dynamicRule", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := ProtectionSource(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseProtectionSource(v string) (any, error) {
    var result ProtectionSource
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_PROTECTIONSOURCE
            case "manual":
                result |= MANUAL_PROTECTIONSOURCE
            case "dynamicRule":
                result |= DYNAMICRULE_PROTECTIONSOURCE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_PROTECTIONSOURCE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeProtectionSource(values []ProtectionSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProtectionSource) isMultiValue() bool {
    return true
}
