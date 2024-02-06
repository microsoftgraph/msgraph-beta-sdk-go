package models
import (
    "errors"
    "math"
    "strings"
)
// 
type ExternalSystemAccessMethods int

const (
    DIRECT_EXTERNALSYSTEMACCESSMETHODS = 1
    ROLECHAINING_EXTERNALSYSTEMACCESSMETHODS = 2
    UNKNOWNFUTUREVALUE_EXTERNALSYSTEMACCESSMETHODS = 4
)

func (i ExternalSystemAccessMethods) String() string {
    var values []string
    options := []string{"direct", "roleChaining", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := ExternalSystemAccessMethods(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseExternalSystemAccessMethods(v string) (any, error) {
    var result ExternalSystemAccessMethods
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "direct":
                result |= DIRECT_EXTERNALSYSTEMACCESSMETHODS
            case "roleChaining":
                result |= ROLECHAINING_EXTERNALSYSTEMACCESSMETHODS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_EXTERNALSYSTEMACCESSMETHODS
            default:
                return 0, errors.New("Unknown ExternalSystemAccessMethods value: " + v)
        }
    }
    return &result, nil
}
func SerializeExternalSystemAccessMethods(values []ExternalSystemAccessMethods) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExternalSystemAccessMethods) isMultiValue() bool {
    return true
}
